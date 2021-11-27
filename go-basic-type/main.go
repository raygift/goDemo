package main

import (
	"fmt"
)

func main() {
	s := int32(0x1234)
	fmt.Printf("s %b\n", s)
	b := byte(s)
	fmt.Printf("0x%x\n", b)
	fmt.Printf("0x%b\n", b)
}
