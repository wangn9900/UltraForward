package remote

import (
	"bytes"
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
}

func NewSSHClient(host string, port int, user, pass string) *SSHClient {
	return &SSHClient{
		Host: host,
		Port: port,
		Config: &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(pass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         15 * time.Second,
		},
	}
}

func (s *SSHClient) Run(cmd string) (string, error) {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	client, err := ssh.Dial("tcp", addr, s.Config)
	if err != nil {
		return "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	session.Stderr = &b
	err = session.Run(cmd)
	return b.String(), err
}

// DeployAgent 自动化部署 UltraForward Agent
func (s *SSHClient) DeployAgent(configJSON string) (string, error) {
	script := fmt.Sprintf(`#!/bin/bash
mkdir -p /etc/ultraforward
echo '%s' > /etc/ultraforward/config.json

# 下载二进制 (假设 GitHub 已放好)
ARCH=$(uname -m)
curl -L -o /usr/local/bin/ultra-agent https://github.com/wangn9900/UltraForward/releases/latest/download/ultra-agent-linux-$ARCH
chmod +x /usr/local/bin/ultra-agent

cat << EOF > /etc/systemd/system/ultra-agent.service
[Unit]
Description=UltraForward Stealth-Pass Agent
After=network.target

[Service]
ExecStart=/usr/local/bin/ultra-agent -config /etc/ultraforward/config.json
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable ultra-agent
systemctl restart ultra-agent
`, configJSON)

	return s.Run(fmt.Sprintf("bash -c \"%s\"", script))
}
