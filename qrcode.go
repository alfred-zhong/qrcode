package main

import (
	"os"
	"strconv"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

func createQRCode(filepath string, message string) (path string, err error) {
	path = resolveFilePath(filepath)

	if path != "" {
		err := qrcode.WriteFile(message, qrcode.Medium, 256, path)
		if err != nil {
			return "", err
		}
	} else {
		png, err := qrcode.Encode(message, qrcode.Medium, 256)
		if err != nil {
			return "", err
		}

		os.Stdout.Write(png)
	}
	return path, nil
}

// resolve filePath if it's empty
func resolveFilePath(filepath string) string {
	if filepath == "" {
		tempFilepath := os.TempDir()

		// append file name with time unix nano string
		tempFilepath += string(os.PathSeparator) +
			strconv.FormatInt(time.Now().UnixNano(), 10) + ".png"

		return tempFilepath
	}

	return filepath
}
