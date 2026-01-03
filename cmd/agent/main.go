package main

import (
	"flag"
	"log"

	"github.com/wangn9900/UltraForward/internal/tunnel"
)

func main() {
	configPath := flag.String("config", "/etc/ultraforward/config.json", "Path to tunnel config")
	flag.Parse()

	log.Println("[UltraForward Agent] Starting Stealth-Pass service...")
	if err := tunnel.RunStandalone(*configPath); err != nil {
		log.Fatalf("Critical Error: %v", err)
	}
}
