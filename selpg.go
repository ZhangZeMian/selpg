package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
)

type selpg_args struct {
	startPage        int
	endPage          int
	inFile           string
	pageLen          int
	pageType         bool //ture for -f, false for -l
	printDestination string
}

var programName []byte

func main() {
	args := new(selpg_args)
	receive_args(args)
	check_args(args)
	process_input(args)
}

func receive_args(args *selpg_args) {
	flag.IntVar(&(args.startPage), "s", -1, "start page")
	flag.IntVar(&(args.endPage), "e", -1, "end page")
	flag.IntVar(&(args.pageLen), "l", 72, "page len")
	flag.StringVar(&(args.printDestination), "d", "", "print destionation")
	flag.BoolVar(&(args.pageType), "f", false, "type of print")
	flag.Parse()
	othersArg := flag.Args()
	if len(othersArg) > 0 {
		args.inFile = othersArg[0]
	} else {
		args.inFile = ""
	}
}

func check_args(args *selpg_args) {
	if args.startPage == -1 || args.endPage == -1 {
		os.Stderr.Write([]byte("you should input -s -e at least\n"))
		os.Exit(0)
	}
	if args.startPage < 1 || args.startPage > (math.MaxInt32-1) {
		os.Stderr.Write([]byte("invalid start page\n"))
		os.Exit(1)
	}
	if args.endPage < 1 || args.endPage > (math.MaxInt32-1) || args.endPage < args.startPage {
		os.Stderr.Write([]byte("invalid end page\n"))
		os.Exit(2)
	}
	if args.pageLen < 1 || args.pageLen > (math.MaxInt32-1) {
		os.Stderr.Write([]byte("invalid page length\n"))
		os.Exit(3)
	}
}

func process_input(args *selpg_args) {
	var (
		reader  *bufio.Reader
		lineCtr int
		pageCtr int
	)

	//init reader
	if args.inFile == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		fileIn, err := os.Open(args.inFile)
		defer fileIn.Close()
		if err != nil {
			os.Stderr.Write([]byte("open file error\n"))
			os.Exit(4)
		}
		reader = bufio.NewReader(fileIn)
	}

	if args.printDestination == "" { //output to os.stdout
		writer := bufio.NewWriter(os.Stdout)
		if args.pageType == true { //-f type
			process_input_f(reader, writer, args, &pageCtr) //-f
		} else { //-l type
			process_input_l(reader, writer, args, &pageCtr, &lineCtr)
		}
	} else { //output to another command by pipe
		cmd_grep := exec.Command("./" + args.printDestination)
		stdin_grep, grep_error := cmd_grep.StdinPipe()
		if grep_error != nil {
			fmt.Println("Error happened about standard input pipe ", grep_error)
			os.Exit(30)
		}
		writer := stdin_grep
		if grep_error := cmd_grep.Start(); grep_error != nil {
			fmt.Println("Error happened in execution ", grep_error)
			os.Exit(30)
		}
		if args.pageType == true { //-d type
			process_input_f_d(reader, writer, args, &pageCtr)
		} else { //-l type
			process_input_l_d(reader, writer, args, &pageCtr, &lineCtr)
		}
		stdin_grep.Close()
		//make sure all the infor in the buffer could be read
		if err := cmd_grep.Wait(); err != nil {
			fmt.Println("Error happened in Wait process")
			os.Exit(30)
		}
	}
	if pageCtr < args.startPage {
		os.Stderr.Write([]byte("start page is greater than total page\n"))
		os.Exit(9)
	}
	if pageCtr < args.endPage {
		os.Stderr.Write([]byte("end page is greater than total page\n"))
		os.Exit(10)
	}
}

func process_input_f(reader *bufio.Reader, writer *bufio.Writer, args *selpg_args, pageCtr *int) {
	*pageCtr = 1
	for {
		char, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Stderr.Write([]byte("read byte from Reader fail\n"))
			os.Exit(7)
		}
		if *pageCtr >= args.startPage && *pageCtr <= args.endPage {
			errW := writer.WriteByte(char)
			if errW != nil {
				os.Stderr.Write([]byte("Write byte to out fail\n"))
				os.Exit(8)
			}
			writer.Flush()
		}
		if char == '\f' {
			(*pageCtr)++
		}
	}
}

func process_input_l(reader *bufio.Reader, writer *bufio.Writer, args *selpg_args, pageCtr *int, lineCtr *int) {
	*lineCtr = 0
	*pageCtr = 1
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			os.Stderr.Write([]byte("read bytes from Reader error\n"))
			os.Exit(5)
		}
		*lineCtr++
		if *pageCtr >= args.startPage && *pageCtr <= args.endPage {
			_, errW := writer.Write(line)
			if errW != nil {
				os.Stderr.Write([]byte("Write to file fail\n"))
				os.Exit(6)
			}
			writer.Flush()
		}
		if *lineCtr >= args.pageLen {
			*lineCtr = 0
			*pageCtr++
		}
	}
}

func process_input_f_d(reader *bufio.Reader, writer io.WriteCloser, args *selpg_args, pageCtr *int) {
	*pageCtr = 1
	for {
		char, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Stderr.Write([]byte("read byte from Reader fail\n"))
			os.Exit(7)
		}
		if *pageCtr >= args.startPage && *pageCtr <= args.endPage {

			writer.Write([]byte{char}) //
		}
		if char == '\f' {
			*pageCtr++
		}
	}
}

func process_input_l_d(reader *bufio.Reader, writer io.WriteCloser, args *selpg_args, pageCtr *int, lineCtr *int) {
	*lineCtr = 0
	*pageCtr = 1
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			os.Stderr.Write([]byte("read bytes from Reader error\n"))
			os.Exit(5)
		}
		*lineCtr++
		if *pageCtr >= args.startPage && *pageCtr <= args.endPage {
			_, errW := writer.Write(line)
			if errW != nil {
				os.Stderr.Write([]byte("Write to file fail\n"))
				os.Exit(6)
			}
		}
		if *lineCtr >= args.pageLen {
			*lineCtr = 0
			*pageCtr++
			writer.Write([]byte("\f"))
		}
	}
}
