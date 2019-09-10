package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the circularArrayRotation function below.
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {

	indexMap := make(map[int32]int32)
	for index, value := range a {
		indexMap[int32(index)] = value
	}

	indexMap = rotateArray(indexMap, a, k)

	var outputSlice []int32
	for _, value := range queries {
		outputSlice = append(outputSlice, indexMap[value])
	}

	return outputSlice
}

func rotateArray(indexMap map[int32]int32, a []int32, k int32) map[int32]int32 {

	k = k % int32(len(a))
	rotatedArray := make(map[int32]int32)

	for key, value := range indexMap {
		if (key + k) <= int32(len(a)-1) {
			rotatedArray[(key + k)] = value
		} else if (key + k) > int32(len(a)-1) {
			kee := (key + k) - int32(len(a)-1) - 1
			rotatedArray[kee] = value
		}
	}
	return rotatedArray
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nkq := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nkq[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nkq[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	qTemp, err := strconv.ParseInt(nkq[2], 10, 64)
	checkError(err)
	q := int32(qTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	var queries []int32

	for i := 0; i < int(q); i++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := circularArrayRotation(a, k, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
