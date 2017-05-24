package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"os"
)

func createImage(filepath string, message string) {
	if filepath != "" {
		err := qrcode.WriteFile(message, qrcode.Medium, 256, filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		png, err := qrcode.Encode(message, qrcode.Medium, 256)
		if err != nil {
			fmt.Println(err)
			return
		}

		os.Stdout.Write(png)
	}
}
