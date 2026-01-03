package tunnel

import (
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"

	"golang.org/x/crypto/chacha20poly1305"
)

const (
	MaxPayloadSize   = 16383
	TagSize          = 16
	NonceSize        = 12
	MaxInitialJitter = 128
)

// SecureConn 包装 net.Conn，处理全加密流 (Stealth-Pass)
type SecureConn struct {
	net.Conn
	enc      cipher.AEAD
	dec      cipher.AEAD
	encNonce [NonceSize]byte
	decNonce [NonceSize]byte
	muEnc    sync.Mutex
	muDec    sync.Mutex
	readBuf  []byte
}

func NewSecureConn(conn net.Conn, keyStr string, isServer bool) (*SecureConn, error) {
	h := sha256.New()
	h.Write([]byte(keyStr))
	key := h.Sum(nil)

	enc, _ := chacha20poly1305.New(key)
	dec, _ := chacha20poly1305.New(key)

	sc := &SecureConn{
		Conn:    conn,
		enc:     enc,
		dec:     dec,
		readBuf: make([]byte, 0),
	}

	if !isServer {
		if _, err := sc.Write(nil); err != nil {
			return nil, err
		}
		jitterLen := rand.Intn(MaxInitialJitter)
		if jitterLen > 0 {
			jitter := make([]byte, jitterLen)
			rand.Read(jitter)
			sc.Conn.Write(jitter)
		}
	} else {
		dummy := make([]byte, 1)
		_, err := sc.Read(dummy)
		if err != nil && err != io.EOF {
			return nil, err
		}
	}

	return sc, nil
}

func (s *SecureConn) incNonce(n *[NonceSize]byte) {
	for i := 0; i < NonceSize; i++ {
		n[i]++
		if n[i] != 0 {
			break
		}
	}
}

func (s *SecureConn) Write(p []byte) (n int, err error) {
	s.muEnc.Lock()
	defer s.muEnc.Unlock()

	if len(p) == 0 {
		lenBuf := make([]byte, 2)
		binary.BigEndian.PutUint16(lenBuf, 0)
		lenNonce := s.encNonce
		s.incNonce(&s.encNonce)
		header := s.enc.Seal(nil, lenNonce[:], lenBuf, nil)
		_, err := s.Conn.Write(header)
		return 0, err
	}

	totalSent := 0
	for totalSent < len(p) {
		chunkSize := len(p) - totalSent
		if chunkSize > MaxPayloadSize {
			chunkSize = MaxPayloadSize
		}

		payload := p[totalSent : totalSent+chunkSize]
		lenBuf := make([]byte, 2)
		binary.BigEndian.PutUint16(lenBuf, uint16(chunkSize))

		lenNonce := s.encNonce
		s.incNonce(&s.encNonce)
		header := s.enc.Seal(nil, lenNonce[:], lenBuf, nil)
		if _, err := s.Conn.Write(header); err != nil {
			return totalSent, err
		}

		nonce := s.encNonce
		s.incNonce(&s.encNonce)
		body := s.enc.Seal(nil, nonce[:], payload, nil)
		if _, err := s.Conn.Write(body); err != nil {
			return totalSent, err
		}
		totalSent += chunkSize
	}
	return totalSent, nil
}

func (s *SecureConn) Read(p []byte) (n int, err error) {
	s.muDec.Lock()
	defer s.muDec.Unlock()

	for {
		if len(s.readBuf) > 0 {
			n = copy(p, s.readBuf)
			s.readBuf = s.readBuf[n:]
			return n, nil
		}

		headerBuf := make([]byte, 2+TagSize)
		if _, err := io.ReadFull(s.Conn, headerBuf); err != nil {
			return 0, err
		}

		lenNonce := s.decNonce
		s.incNonce(&s.decNonce)
		lenPlain, err := s.dec.Open(nil, lenNonce[:], headerBuf, nil)
		if err != nil {
			return 0, fmt.Errorf("decrypt header failed: %v", err)
		}

		chunkLen := binary.BigEndian.Uint16(lenPlain)
		if chunkLen == 0 {
			continue
		}

		bodyBuf := make([]byte, int(chunkLen)+TagSize)
		if _, err := io.ReadFull(s.Conn, bodyBuf); err != nil {
			return 0, err
		}

		nonce := s.decNonce
		s.incNonce(&s.decNonce)
		payload, err := s.dec.Open(nil, nonce[:], bodyBuf, nil)
		if err != nil {
			return 0, fmt.Errorf("decrypt body failed: %v", err)
		}

		n = copy(p, payload)
		if n < len(payload) {
			s.readBuf = append(s.readBuf, payload[n:]...)
		}
		return n, nil
	}
}
