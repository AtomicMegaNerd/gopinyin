package main

import (
	"flag"
	"log"
)

func main() {
	infilePtr := flag.String("infile", nil, "The path of the file to convert")
	outfilePtr := flag.String("outfile", nil, "The path of the output from this program")
	debugPtr := flag.Bool("debug", false, "Display debug information?")

	log.Println("Starting Pinyin Conversion Program...")
}
