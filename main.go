// Filename: main.go
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
flag.Parse()

for i := 0 i < *num; i++ {
	fmt.Println(*msg)
}

}
