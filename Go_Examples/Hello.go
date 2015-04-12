/*
1. We want to calculate a sum of squares of some integers, excepting negatives
2. The first line of the input will be an integer N (1 <= N <= 100)
3. Each of the following N test cases consists of one line containing an integer X (0 < X <= 100), followed by X integers (Yn, -100 <= Yn <= 100) space-separated on the next line
4. For each test case, calculate the sum of squares of the integers excepting negatives, and print the calculated sum to the output. No blank line between test cases
5. (Take input from standard input, and output to standard output)
*/
package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

func lvl1(s *bufio.Scanner, n int, output string) string{
	if n>0 {
		output = lvl1(s,n-1, output)
		output = lvl2(s,output)
	}
	return output
}

func lvl2(s *bufio.Scanner, output string) string{
	var sum int
	s.Scan()
	len, err :=strconv.Atoi(s.Text())
	if err != nil{
		fmt.Println(err)
	}
	sum = lvl3(s, 0, len)
	t := strconv.Itoa(sum)
	output = output + t + "\n"
	return output
}

// sum the squares of an array of integer by recursion.
func lvl3(s *bufio.Scanner, sum int, len int) int{
	if len > 0 {
		s.Scan()
		val, err :=strconv.Atoi(s.Text())
		if err != nil{
			fmt.Println(err)
		}
		if val < 0{
			sum = lvl3(s, sum, len-1)
		}else{
			sum = lvl3(s, sum + val*val, len-1)
		}
	}
	return sum
}

func main() {
	var output string
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	count, err :=strconv.Atoi(s.Text())
	if err != nil{
		fmt.Println(err)
	}
	output = lvl1(s, count, output)	
	fmt.Println(output)
}