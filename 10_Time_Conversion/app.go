package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	arr := strings.Split(s,":")
	s = s[:len(s)-2]
	if strings.Contains(arr[2], "AM"){
		if arr[0]!="12"{
			return s
		} else{
			return "00"+s[2:]
		}
	} else{
		if arr[0]=="12"{
			return s
		} else {
			hour, err := strconv.Atoi(arr[0])
			if err!=nil{
				log.Fatal()
			}
			hour+=12
			return strconv.Itoa(hour)+s[2:]
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}