package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func valid_sign(ch string) bool {
	f := true
	if ch[0] == '-' {
		if ch[1] == '0' {
			f = false
		}
	} else if ch[0] == '0' && len(ch) != 1 {
		f = false
	}
	for _, x := range ch {
		if x > '9' && x < '0' {
			if x != '-' {
				f = false
				break
			}
		}
	}
	return f
}
func validate_array(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] && str[i] == ' ' {
			return false
		}
	}
	return true
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	input, _ := in.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)

	for i := 0; i < num; i++ {
		//flag for print
		f := true
		//get count of elemets in string
		num_of_el, _ := in.ReadString('\n')
		num_of_el = strings.TrimSpace(num_of_el)
		num_of_el_int, _ := strconv.Atoi(num_of_el)
		//get string and split it by " "
		input_string, _ := in.ReadString('\n')
		if validate_array(input_string) == false {
			fmt.Println("no")
			continue
		}
		input_string = strings.TrimSpace(input_string)
		var mas_el []string = strings.Split(input_string, " ")
		//second input string
		s_input_string, _ := in.ReadString('\n')
		if validate_array(s_input_string) == false {
			fmt.Println("no")
			continue
		}

		s_input_string = strings.TrimSpace(s_input_string)

		var s_mas_el []string = strings.Split(s_input_string, " ")
		
		if len(mas_el) != len(s_mas_el) {
			f = false
		}

		//конвертнули первый массив в инты
		mas_el_int := make([]int, len(mas_el))
		for i, str := range mas_el {
			elem, err := strconv.Atoi(str)

			if err != nil {
				f = false
				break
			}
			mas_el_int[i] = elem
		}

		//конвертнули второй массив в инты
		s_mas_el_int := make([]int, len(s_mas_el))
		for i, str := range s_mas_el {
			elem, err := strconv.Atoi(str)

			if err != nil {
				f = false
				break
			}
			s_mas_el_int[i] = elem
		}

		//validate chars in string
		for _, x := range s_mas_el {
			if valid_sign(x) == false {
				f = false
				break
			} else {
				continue
			}
		}

		sort.Ints(mas_el_int) // sorting first array

		//chek equals in arrays
		for j := 0; j < num_of_el_int; j++ {
			if mas_el_int[j] != s_mas_el_int[j] {
				f = false
				break
			}
		}

		if f == false {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}
	}

}
