package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	scanner "github.com/jsypower/PortScan/Scanner"
)

func GeneratePortList(_sPorts []string) []string {
	_iPorts := make([]string, len(_sPorts))
	for i, port := range _sPorts {
		_iPorts[i] = port
	}
	return _iPorts
}

func GeneratePortRange(_sPorts []string) []string {

	iStartPort, err := strconv.Atoi(_sPorts[0])
	if err != nil {
		log.Fatalf("[!] Error parsing initial port: %v", err)
		os.Exit(1)
	}

	iEndPort, err := strconv.Atoi(_sPorts[1])
	if err != nil {
		log.Fatalf("[!] Error parsing final port: %v", err)
		os.Exit(1)
	}

	_iPorts := make([]string, iEndPort-iStartPort+1)
	for i := range _iPorts {
		_iPorts[i] = fmt.Sprintf("%d", iStartPort+i)
	}

	return _iPorts
}

func main() {
	log.Println("Port Scanner v0.1a")

	S_FLAG_HOST := flag.String("host", "localhost", "Specify the target host")
	S_FLAG_PORT := flag.String("port", "1-1024", "Specify the target port to scan.")
	flag.Parse()
	*S_FLAG_HOST = strings.ReplaceAll(*S_FLAG_HOST, " ", "")
	*S_FLAG_PORT = strings.ReplaceAll(*S_FLAG_PORT, " ", "")

	log.Printf("Target: %s", *S_FLAG_HOST)
	var _iPorts []string
	// Convert Hyphenated Port Range
	if strings.Contains(*S_FLAG_PORT, "-") && !strings.Contains(*S_FLAG_PORT, ",") {
		_sPorts := strings.Split(*S_FLAG_PORT, "-")
		_iPorts = GeneratePortRange(_sPorts)
	}

	// Convert CSV Port Range
	if strings.Contains(*S_FLAG_PORT, ",") && !strings.Contains(*S_FLAG_PORT, "-") {
		_sPorts := strings.Split(*S_FLAG_PORT, ",")
		_iPorts = GeneratePortList(_sPorts)
	}

	log.Printf("Ports:\n%v", _iPorts)

	for i := range _iPorts {
		log.Printf("[i] Testing connection to %s:%s", *S_FLAG_HOST, _iPorts[i])
		if scanner.SynScan("tcp", *S_FLAG_HOST, _iPorts[i]) {
			log.Println("[i] Port is OPEN")
		} else {
			log.Println("[!] Port is CLOSED")
		}
	}
}
