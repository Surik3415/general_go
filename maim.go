package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	for {
		fmt.Println("Enter some hex num or stop if you want to exit:")

		var input string
		fmt.Scanln(&input)

		// fmt.Println(input)
		input = strings.ToLower(input)

		if input == "stop" {
			break
		}

		i := new(big.Int)
		// _, ok := i.SetString(processHex(input), 16)
		// if !ok {
		// 	fmt.Println("it is not hex format of number")
		// } else {
		// 	fmt.Println(i)
		// }

		if _, ok := i.SetString(processHex(input), 16); !ok {
			fmt.Println("it is not hex format of number")
			continue
		}

		fmt.Println(i)
	}

}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}
