package main

import "pynote/server"

func main() {
	config := "./config.json"
	s := server.NewServer(config)
	s.Run()
}
