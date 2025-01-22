package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	input, _ := in.ReadString('\n')
	input = input[:len(input)-1]
	num, _ := strconv.Atoi(input)

	for i := 0; i < num; i++ {
		str, _ := in.ReadString('\n')
		str = str[:len(str)-1]
		f := true
		if len(str) < 2 {
			f = false
		} else if (str[len(str)-1] != 'D') || (str[0] != 'M') {
			f = false
		} else {
			prev := str[0]
			for _, el := range str[1:len(str)] {
				if (prev == 'M') && (el == 'M') {
					f = false
					break
				} else if (prev == 'R') && (el != 'C') {
					f = false
					break
				} else if (prev == 'C') && (el != 'M') {
					f = false
					break
				} else if (prev == 'D') && (el != 'M') {
					f = false
					break
				}
				prev = uint8(el)
			}
		}
		if f == false {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
