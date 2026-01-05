package main

import (
	"github.com/thuta/ggpoker/p2p"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	d := deck.New()
	// 	fmt.Println(d)
	// 	fmt.Println("-------")
	// }

	// server config
	cfg := p2p.ServerConfig{
		Version:    "GGPOKER V0.1-beta",
		ListenAddr: ":3000",
	}
	server := p2p.NewServer(cfg)
	go server.Start()

	// remote server config
	remoteCfg := p2p.ServerConfig{
		Version:    "GGPOKER V0.1-beta",
		ListenAddr: ":4000",
	}
	remoteServer := p2p.NewServer(remoteCfg)
	go remoteServer.Start()
	if err := remoteServer.Connect("localhost:3000"); err != nil {
		panic(err)
	}

	select {}
}
