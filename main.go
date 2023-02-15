package main

import (
	"flag"
	"fmt"
	"strings"
)

//demonstrate flags

func main(){
// set the flags
msg := flag.String("msg", "howdy, stranger!", "the message to display")
num := flag.Int("num", 1, "how many times to print the message")
caps := flag.Bool("caps", false, "Should the string be all caps")
reverse := flag.Bool("r", false, "reverse the string (true or false)")
flag.Parse()
// check if it should be uppercase before printing
if *caps{
	*msg = strings.ToUpper(*msg)
}
// check if we should reverse the string
if *reverse{
	//reverse string
	var result string
	for _, value := range *msg{
		result = string(value) + result
		
	}
	// Write the reverse string to *msg
	*msg = result
}
// print the string
for i := 0; i < *num; i++ {
	fmt.Println(*msg)
}


}