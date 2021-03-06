package main

import (
	"flag"
	"github.com/bygo/o/cmd"
	_ "github.com/bygo/o/cmd/make"
	_ "github.com/bygo/o/cmd/new"
)

func main() {
	flag.Usage = cmd.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		cmd.Usage()
	}

	if c, ok := cmd.Commands[args[0]]; ok {
		err := c.Flag.Parse(args[1:])
		if err != nil {
			panic(err)
		}
		c.Run(c, args[1:])
	}
}
