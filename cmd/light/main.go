package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/petderek/light"
)

const howto = `light [port] [red|green|yellow] [on|off]

Used to control a multicolor usb light. only supports three colors.
`

func main() {
	flag.Parse()
	c, err := light.Infer(flag.Arg(1), flag.Arg(2))
	if err != nil || c == light.ZERO {
		usage()
		os.Exit(1)
	}
	err = light.Send(flag.Arg(0), c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to send light cmd: %s", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprint(os.Stderr, howto)
}
