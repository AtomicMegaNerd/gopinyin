package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type runner struct {
	debug   bool
	infile  string
	outfile string
}

func NewRunner(debug bool, infile string, outfile string) *runner {
	return &runner{debug, infile, outfile}
}

func (r *runner) ConvertPinyinText() error {
	if text, err := getLinesFromFile(r.infile); err == nil {
		converter := NewPinyinConverter(r.debug)
		converted_text := converter.DoConvert(text)
		writeLinesToFile(converted_text, r.outfile)
	} else {
		return err
	}
	return nil
}

func getLinesFromFile(infile string) ([]string, error) {
	file, err := os.Open(infile)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Chown()
	file.Close()
	log.Printf("Opened filed %s successfully to read data for conversion", infile)
	return text, nil
}

func writeLinesToFile(convertedText []string, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	for _, line := range convertedText {
		writer.WriteString(fmt.Sprintf("%s\n", line))
	}
	writer.Flush()
	file.Close()

	log.Printf("Saved converted Pinyin data to file %s successfully", outfile)
	return nil
}
