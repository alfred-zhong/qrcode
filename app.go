package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

const cmd_to_open_file = "open"

type Cmd struct {
	filePath string
	noOpen   bool
	message  string
}

func main() {
	cmd := Cmd{}
	flag.StringVar(&cmd.filePath, "f", "", "file path to save the qrcode image")
	flag.BoolVar(&cmd.noOpen, "n", false, "don't open image when created")
	flag.Parse()

	if flag.NArg() > 0 {
		for i := 0; i < flag.NArg(); i++ {
			if i > 0 {
				cmd.message += " "
			}
			cmd.message += flag.Arg(i)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		bb := make([]byte, 4<<10)
		n, err := reader.Read(bb)
		if err != nil {
			fmt.Fprintln(os.Stderr, "read from os.Stdin error: %v", err)
			return
		}

		if n == 0 {
			fmt.Println("no data read from os.Stdin")
			return
		} else {
			cmd.message = string(bb[:n])
		}
	}

	// resolve filePath if it's empty
	resolveFilePath(&cmd.filePath)

	// fmt.Printf("filepath:" + cmd.filePath + ", message:" + cmd.message + "\n")
	createImage(cmd.filePath, cmd.message)

	// Open after created the image
	if !cmd.noOpen {
		openFile(cmd.filePath)
	}
}

// resolve filePath if it's empty
func resolveFilePath(filePath *string) {
	if *filePath == "" {
		tempFilePath := os.TempDir()

		// if it's in OSX, temp directory is in "/var/folders/".
		// this folder is not suitable.
		// in linux, maybe no problem
		if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
			tempFilePath = "/tmp"
		}

		// append file name with time unix nano string
		tempFilePath += string(os.PathSeparator) +
			strconv.FormatInt(time.Now().UnixNano(), 10) + ".png"

		*filePath = tempFilePath
	}
}

// Open file if can (with the default program, in OSX with "open")
func openFile(filePath string) {
	cmdPath, err := exec.LookPath(cmd_to_open_file)
	if err == nil {
		// fmt.Println(cmdPath)
		cmd := exec.Command(cmdPath, filePath)
		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, "image fail to open")
		}
	}
}
