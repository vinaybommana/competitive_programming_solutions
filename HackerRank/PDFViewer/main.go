package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the designerPdfViewer function below.
func designerPdfViewer(h []int32, word string) int32 {
	lcAlphabets := "abcdefghijklmnopqrstuvwxyz"
	lcMap := make(map[string]int32)
	var maxHeight int32
	for index, s := range lcAlphabets {
		lcMap[string(s)] = h[index]
	}

	stringLength := len(word)
	for _, s := range word {
		if maxHeight < lcMap[string(s)] {
			maxHeight = lcMap[string(s)]
		}
	}

	return maxHeight * int32(stringLength)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	hTemp := strings.Split(readLine(reader), " ")

	var h []int32

	for i := 0; i < 26; i++ {
		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
		checkError(err)
		hItem := int32(hItemTemp)
		h = append(h, hItem)
	}

	word := readLine(reader)

	result := designerPdfViewer(h, word)

	fmt.Fprintf(writer, "%d\n", result)

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
