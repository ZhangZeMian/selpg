package main

import (
	"bufio"
	"io"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	f, err := os.OpenFile("receiverText.txt", os.O_WRONLY, 0666)

	if err != nil {
		os.Stderr.WriteString("fail to open receiver")
	}
	writer := bufio.NewWriter(f)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			os.Stderr.Write([]byte("read bytes from Reader error\n"))
			os.Exit(5)
		}
		writer.Write(line)
		writer.Flush()
	}
	/*
		writer := bufio.NewWriter(os.Stdout)
		os.Stdout.Write([]byte("rAAAA\n"))
		for {
			os.Stdout.Write([]byte("1\n"))
			line, err := reader.ReadBytes('\n')
			if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
				if err == io.EOF {
					break
				}
				os.Stderr.Write([]byte("read bytes from Reader error\n"))
				os.Exit(5)
			}
			writer.Write(line)
			os.Stdout.Write([]byte("2\n"))
			writer.Flush()
		}
	*/
}
