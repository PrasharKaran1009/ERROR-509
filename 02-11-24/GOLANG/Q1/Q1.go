package q1

import (
	"GOLANG/INPUT"
	"fmt"
	"strings"
)

func logic() string {

	n := input.ReadInt()
	// var answer string = "1"
	// for i:=1;i<n;i++ {
	// 	answer+="0"
	// }

	//TLE ->1500ms
	return "1" + strings.Repeat("0", n-1)
}

func Answer1() {
	n := input.ReadInt()

	for ; n > 0; n-- {
		fmt.Println(logic())
	}
}

//just print 1 and other 0 it will be minimum oneness
