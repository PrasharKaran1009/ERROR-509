// this one is not taking input correctly in looping so here the normal approach look at last
package q3

import (
	"GOLANG/INPUT"
	"fmt"
)

func logic() int {

	n := input.ReadInt()
	r := input.ReadInt()

	//family := make([]int, n) //this is not needed as we only need to calculate total passengers and unhappy ones
	unhappy := 0
	total_passengers := 0

	for i := 0; i < n; i++ {
		x := input.ReadInt()
		//family[i] = x
		unhappy += x % 2
		total_passengers += x
	}

	if unhappy == 0 {
		return total_passengers
	}

	happy := total_passengers - unhappy
	remaining_seat := 2*r - happy

	if remaining_seat/2 >= unhappy {
		happy += unhappy
	} else {
		happy += remaining_seat - unhappy
	}

	return happy
}

func Answer3() {
	n := input.ReadInt()

	for ; n > 0; n-- {
		fmt.Println(logic())
	}
} 


/*
package q3

import (
	"fmt"
)
 
func logic() int {
	var n, r int
	fmt.Scan(&n, &r)
 
	unhappy := 0
	totalPassengers := 0
 
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		totalPassengers += x
		unhappy += x % 2
	}
 
	if unhappy == 0 {
		return totalPassengers
	}
 
	happy := totalPassengers - unhappy
	remainingSeat := 2*r - happy
 
	if remainingSeat/2 >= unhappy {
		happy += unhappy
	} else {
		happy += remainingSeat - unhappy
	}
 
	return happy
}
 
func Answer3() {
	var t int
	fmt.Scan(&t)
 
	for t > 0 {
		fmt.Println(logic())
		t--
	}
}
*/