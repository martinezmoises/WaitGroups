package main

import (
	"flag"
	"fmt"
	"strings"
)

//domnstrate flags

func main() {
	msg := flag.String("msg", "Howdy, stranger!", "the message")
	num := flag.Int("num", 1, "How many times to print the message")
	caps := flag.Bool("U", false, "Specify whether to uppercase the string(true/false)")
	reverse := flag.Bool("r", false, "reverse the string (true or false) ")

	flag.Parse()
	// check if it is should be uppercase before printing it.
	if *caps {
		*msg = strings.ToUpper(*msg)
	}

	//check if we shoul reverse the string
	if *reverse {
		//Reversing the string
		var result string
		for _, v := range *msg {
			result = string(v) + result
		}

		//Write the reversed string to the msg variable
		*msg = result
	}

	//print the s string
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}

}
