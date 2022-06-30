package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/beom-yun/sbcoin/explorer"
	"github.com/beom-yun/sbcoin/rest"
)

func usage() {
	fmt.Printf("Welcone to SB Coin!\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest' and 'both'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 1000)
	default:
		usage()
	}
}
