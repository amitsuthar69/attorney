package main

import (
	"log"

	"attorney/internal/proxy"
)

func main() {
	// Proxy with default config
	p := proxy.NewProxy(proxy.Config{
		ListenAddr: ":8888",
	})

	if err := p.Start(); err != nil {
		log.Fatalf("Proxy startup failed: %v", err)
	}
}
