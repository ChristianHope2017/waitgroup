package main

import (
	"flag"
	"fmt"
	"strings"
)

func one (wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hola")
}
func two(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("ni hao")
}
func three (wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
}
func main(){
	var wg sync.WaitGroup
	wg.Add(3)
	go one(&wg)
	go two(&wg)
	go three(&wg)
	//Let's Delay
	wg.Wait()
}
