package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

/*
I am using bufio for efficient input handling.
The bufio package provides buffered I/O, which is generally more efficient than using
standard input functions like fmt.Scan or fmt.Scanf. Buffered reading reduces
the number of I/O operations by reading data in chunks, which is especially
beneficial when dealing with larger inputs or multiple lines of input.
*/
var reader = bufio.NewReader(os.Stdin)

/*
readInt takes input as a string, trims it to remove any leading or trailing whitespace,
converts it to an integer, and returns the integer value.
*/

func readInt() int {
	n,_ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	/*num,_ := strconv.ParseInt(n,10,64);
	return int(num)*/
	num,_:=strconv.Atoi(n)
	return num
}

func solutionERR() int {
	var n = readInt()

	num := make([][]int, n)
	for i := range num {
		num[i] = make([]int, n)
	}

	// Input the matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num[i][j] = readInt()
		}
	}

	sum := 0
	for i := 0; i < 2*n-1; i++ {
		mini := 0
		var j, k int

		if i < n {
			k = n - 1 - i
		} else {
			j = i - n + 1
		}

		for j < n && k < n {
			if num[j][k] < mini { 
				mini = num[j][k]
			}
			j++
			k++
		}

		sum += int(math.Abs(float64(mini)))
	}

	return sum
}

func err() {
	var t = readInt()
	for t > 0 {
		t--
		fmt.Println(solutionERR())
	}
}

/*
   04
   03 14
   02 13 24
   01 12 23 34
   00 11 22 33 44
   10 21 32 43
   20 31 42
   30 41
   40

take minimum element and add there absolute and return that sum
*/
