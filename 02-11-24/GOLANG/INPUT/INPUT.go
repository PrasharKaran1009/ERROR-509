package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadInt() int {
	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	/*
	num,_ := strconv.ParseInt(n,10,64);
	return int(num)
	*/
	num, _ := strconv.Atoi(n)
	return num
}

func ReadString () string {
	n,_ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	return n
}