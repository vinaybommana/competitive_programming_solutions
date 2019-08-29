package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {

	int32asIntValues := make([]int, len(a))
	for i, val := range a {
		int32asIntValues[i] = int(val)
	}

	numberSlice := int32asIntValues
	// sort the slice
	sort.Ints(numberSlice)

	// calculating weight of numbers
	numbersWeight := make(map[int]int)
	for _, value := range int32asIntValues {
		_, exist := numbersWeight[value]
		if exist {
			numbersWeight[value]++
		} else {
			numbersWeight[value] = 1
		}
	}

	var uniqueElements []int
	var highestNumberPick int

	for key := range numbersWeight {
		uniqueElements = append(uniqueElements, key)
	}
	sort.Ints(uniqueElements)

	for index := 0; index < len(uniqueElements)-1; index++ {
		if Abs(uniqueElements[index+1]-uniqueElements[index]) <= 1 {
			sum := numbersWeight[uniqueElements[index]] + numbersWeight[uniqueElements[index+1]]
			if sum > highestNumberPick {
				highestNumberPick = sum
			}
		}
	}

	for _, value := range numbersWeight {
		if value > highestNumberPick {
			highestNumberPick = value
		}
	}

	return int32(highestNumberPick)
}

// Abs : return absolute value of given integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

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
