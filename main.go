package main

import (
	"fmt"

	"github.com/akamensky/argparse"
	scanner "github.com/jsypower/PortScan/Scanner"
)

func main() {
	fmt.Println("Port Scanner v0.1a")

	parser := argparse.NewParser("print", "Prints provided string to stdout")

	if scanner.SynScan("tcp", "scanme.nmap.org", 80) {
		fmt.Println("Port: Open")
	} else {
		fmt.Println("Port: Closed")
	}

}
