package main

import (
	"awesomeProject/pipeline"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	p := createNetworkPipeline("small.in", 512, 4)
	// time.Sleep(time.Hour)
	writeToFile(p, "small.out")
	printFile("small.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	for v := range p {
		fmt.Println(v)
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount

	var sortResults []<-chan int

	for i := 0; i< chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i * chunkSize), 0)
		source := pipeline.ReaderSource(
			bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount
	var sortAddr []string
	for i := 0; i< chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i * chunkSize), 0)
		source := pipeline.ReaderSource(
			bufio.NewReader(file), chunkSize)
		addr := ":" + strconv.Itoa(7000 + i)

		pipeline.NetworkSink(addr, pipeline.InMemSort(source))

		sortAddr = append(sortAddr, addr)

	}
	// return nil

	var sortResults []<-chan int
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}