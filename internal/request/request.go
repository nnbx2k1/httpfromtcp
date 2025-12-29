package request

import (
	"fmt"
	"io"
	"strings"
	"errors"
)



type RequestLine struct{
	HttpVersion string
	RequestTarget string
	Method string
}

type Request struct {
	RequestLine RequestLine
}


var ERROR_MALFORMED_REQUEST_LINE = fmt.Errorf("malformed request line")
var ERROR_INCOMPLETE_START_LINE = fmt.Errorf("incomplete start line")
var ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("unsupported http version")
var SEPARATOR = "\r\n"


func ParseRequestLine(b string) (*RequestLine, string, error){
   idx := strings.Index(b , SEPARATOR)
   if idx == -1{
	return nil, b, ERROR_INCOMPLETE_START_LINE
   }

   startLine := b[:idx]
   outputString := b[idx+len(SEPARATOR):]

   parts := strings.Split(startLine," ")
   if len(parts) != 3 {
	return nil, outputString , ERROR_MALFORMED_REQUEST_LINE
   }

   httpParts := strings.Split(parts[2],"/")

   if len(httpParts)!=2 || httpParts[0] != "HTTP" {
	 return nil,outputString,ERROR_MALFORMED_REQUEST_LINE
   }

   if httpParts[1] != "1.1"{
	return nil,outputString,ERROR_UNSUPPORTED_HTTP_VERSION
   }



   Rl := &RequestLine{
	Method: parts[0],
	RequestTarget: parts[1],
	HttpVersion: httpParts[1],
   }


   return Rl,outputString, nil
}

func RequestFromReader(reader io.Reader) (*Request  , error){

   data, err := io.ReadAll(reader)
   if err != nil {
      return nil, errors.Join(
		fmt.Errorf("unable to io.ReadAll"),
		err,
	  )
   }

   str := string(data)
   rl ,_ ,err := ParseRequestLine(str)

   if err != nil {
	 return nil, err
   }

   return &Request{
	RequestLine: * rl,
   }, err
}

