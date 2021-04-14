package server

import (
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli"
)

var (
	config string
)

func APP() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "PyNote"
	app.Usage = "高效率云笔记"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config",
			Usage:       "set config path",
			Value:       "",
			Required:    false,
			Destination: &config,
		},
	}
	if config == "" {
		config = "config.json"
	}
	app.Action = func(c *cli.Context) error {
		s := NewServer(config)
		s.Run()
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
