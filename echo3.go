//Echo 3 print its command-lin Args
package main

import (
	"fmt"
	"os"
	)
	
func main() {
	for i:=1;i<len(os.Args);i++ {
		fmt.Println("Arg ", i, os.Args[i])
	}
}
