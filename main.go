package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type cfg struct {
	filepath string
	noOpen   bool
}

func main() {
	var cfg cfg
	flag.StringVar(&cfg.filepath, "f", "", "file path to save the qrcode image")
	flag.BoolVar(&cfg.noOpen, "n", false, "don't open image when created")
	flag.Parse()

	var content string

	if flag.NArg() > 0 {
		// read data from cli
		content = strings.Join(flag.Args(), " ")
	} else {
		// read data from stdin
		reader := bufio.NewReader(os.Stdin)
		bb := make([]byte, 1024)

		for {
			n, err := reader.Read(bb)

			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Fprintf(os.Stderr, "read from os.Stdin error: %v", err)
					return
				}
			}

			content += string(bb[:n])
		}
	}

	fp, err := createQRCode(cfg.filepath, content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create QR code fail: %v", err)
		os.Exit(1)
	}

	// Open after created the image
	if !cfg.noOpen {
		open(fp)
	}
}
