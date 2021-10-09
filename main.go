package main

import (
	"flag"
	"log"
)

func main() {
	infilePtr := flag.String("infile", "", "The path of the file to convert")
	outfilePtr := flag.String("outfile", "", "The path of the output from this program")
	debugPtr := flag.Bool("debug", false, "Display debug information?")

	flag.Parse()

	log.Println("Starting Pinyin conversion program...")

	runner := NewRunner(*debugPtr, *infilePtr, *outfilePtr)
	err := runner.ConvertPinyinText()
	if err != nil {
		panic(err)
	}
}
