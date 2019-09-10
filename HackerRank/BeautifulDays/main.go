package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {

	var BeautifulDays int32
	for index := i; index <= j; index++ {
		if Abs(index-reversedNumber(index))%k == 0 {
			BeautifulDays++
		}
	}
	return BeautifulDays
}

func reversedNumber(number int32) int32 {
	var reversedNumber int32
	for ; number > 0; number /= 10 {
		reversedNumber = (reversedNumber * 10) + (number % 10)
	}
	return reversedNumber
}

// Abs returns absolute value of the given number
func Abs(number int32) int32 {
	if number < 0 {
		return -number
	}
	return number
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	ijk := strings.Split(readLine(reader), " ")

	iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

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
