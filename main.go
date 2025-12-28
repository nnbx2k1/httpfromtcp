package main

import (
	"fmt"
	"os"
)

func main(){

	f,err := os.Open("messages.txt");

	if err != nil {
	    fmt.Print("error in opening the file")
	}

	for {
		data := make([]byte,8)
		n, err := f.Read(data)
		if err != nil {
        fmt.Print("error in reading the file")
		break
		}

		fmt.Printf("readed info %s \n",string(data[:n]))
	}


}