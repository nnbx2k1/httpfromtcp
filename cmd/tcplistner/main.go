package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)


func getLinesChannel(f io.ReadCloser) <- chan string  {
   
	out := make(chan string,1)

    reader := bufio.NewReader(f)

	go func() {
	 defer f.Close()
	 defer close(out)

	 for {
		data , err := reader.ReadString('\n')
		if err != nil {
          if err == io.EOF {
			out <- data
			break
		  }
        fmt.Print("error in reading the file \n")
		break

		}

		out <- data
	 }
	}()
	return out
}



func main(){
	listner,err := net.Listen("tcp","localhost:3000")
	if err != nil {
		fmt.Print("error in listening")
	}
    
	for{
		conn, err := listner.Accept()
		if err != nil{
			fmt.Print("error in accepting the connection")
		}
		for line := range getLinesChannel(conn) {
			fmt.Printf("message recieved: %s \n",line)
		}
	}
}