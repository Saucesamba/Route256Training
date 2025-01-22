package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//create bufers for i/o operations
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	input, _ := in.ReadString('\n')
	input = input[:len(input)-1]
	num, _ := strconv.Atoi(input)

	for i := 0; i < num; i++ {

		sal, _ := in.ReadString('\n')
		sal = sal[:len(sal)-1]
		min_num := 9
		min_ind := 0
		for ind, el := range sal {
			n := int(el) - 48
			if n <= min_num {
				min_num = n
				min_ind = ind
			} else {
				break
			}
		}
		if len(sal) != 1 {
			fmt.Println(sal[:min_ind] + sal[min_ind+1:])

		} else {
			fmt.Println(0)
		}
	}
}
