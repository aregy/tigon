package main

import (
	"os"
	"fmt"
	"io"
)
// window set to a constant size (in bytes)
const window = 20

func View(file string){
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	var frameNum int
	buffer := make([]byte, min(fileInfo.Size(), window))
	for {
		if frameNum > 0 {
			fmt.Printf(">> Continue... view %d remaining bytes\n", fileInfo.Size() - int64(window * frameNum))
			fmt.Scanln()
			buffer = make([]byte, fileInfo.Size() - int64(frameNum * window))
		}
		n, err := f.Read(buffer)
		if err != nil {
			if err == io.EOF {
				return
			}
		}
		frameNum += 1
		fmt.Printf(">> Incoming windowfull #%d (of %d bytes)...\n", frameNum, len(buffer[:n]))
		fmt.Println(string(buffer))
	}
}

func min (n ...int64) int64{
	if len(n) == 0 {
		panic(fmt.Errorf("Empty set has no minimum member"))
	}
	m := n[0]
	for i := 0; i < len(n); i++ {
		if n[i] < m {
			m = n[i]
		}
	}
	return m
}

func main(){
	if len(os.Args) < 2 {
		panic(fmt.Errorf("No file name provided"))
	}
	View(os.Args[1])
}
