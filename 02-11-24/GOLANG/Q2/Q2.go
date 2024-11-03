package q2 
import (
	"GOLANG/INPUT"
	"fmt"
)

func logic () int {
	s:= input.ReadString()
	t:= input.ReadString()

	count_sequence := 0
	
	for i:=0; i<len(s) && i<len(t) ; i++ {
		if s[i] == t[i] {
			count_sequence++
		} else {
			break
		}
	}

	if count_sequence == 0 { return len(s)+len(t)}
	 
	return (len(s)-count_sequence) + (len(t)-count_sequence)+ (count_sequence+1)
}
func Answer2 (){
	n:=input.ReadInt()

	for ;n>0;n-- {
		fmt.Println(logic())
	}
}