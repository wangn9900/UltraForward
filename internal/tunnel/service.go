package tunnel

import (
	"context"
	"io"
	"log"
	"net"
	"sync/atomic"
)

// TrafficCounter 原子计数器，用于统计流量
type TrafficCounter struct {
	Upload   int64
	Download int64
}

// TransitServer (中转入口)
type TransitServer struct {
	RuleID     uint
	ListenAddr string
	TargetAddr string // Exit 节点的地址
	Key        string
	Stats      *TrafficCounter
}

func (s *TransitServer) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	log.Printf("[Transit] Rule %d listening on %s", s.RuleID, s.ListenAddr)

	go func() {
		<-ctx.Done()
		l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go s.handle(conn)
	}
}

func (s *TransitServer) handle(in net.Conn) {
	defer in.Close()
	// 连接到 Exit 节点
	out, err := net.Dial("tcp", s.TargetAddr)
	if err != nil {
		log.Printf("[Transit] Failed to connect to exit %s: %v", s.TargetAddr, err)
		return
	}
	defer out.Close()

	// 开启隧道加密
	sConn, err := NewSecureConn(out, s.Key, false)
	if err != nil {
		log.Printf("[Transit] Handshake failed: %v", err)
		return
	}

	// 流量统计闭包
	go func() {
		n, _ := io.Copy(sConn, in)
		atomic.AddInt64(&s.Stats.Upload, n)
	}()
	n, _ := io.Copy(in, sConn)
	atomic.AddInt64(&s.Stats.Download, n)
}

// ExitServer (出口落地)
type ExitServer struct {
	RuleID     uint
	ListenAddr string
	LocalDest  string // 最终转发目标 (如 127.0.0.1:8443)
	Key        string
}

func (s *ExitServer) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	log.Printf("[Exit] Rule %d listening on %s", s.RuleID, s.ListenAddr)

	go func() {
		<-ctx.Done()
		l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go s.handle(conn)
	}
}

func (s *ExitServer) handle(in net.Conn) {
	defer in.Close()

	// 解密隧道流量
	sConn, err := NewSecureConn(in, s.Key, true)
	if err != nil {
		log.Printf("[Exit] Decrypt handshake failed: %v", err)
		return
	}

	// 连接到最终目标
	out, err := net.Dial("tcp", s.LocalDest)
	if err != nil {
		log.Printf("[Exit] Failed to connect to dest %s: %v", s.LocalDest, err)
		return
	}
	defer out.Close()

	go io.Copy(out, sConn)
	io.Copy(sConn, out)
}
