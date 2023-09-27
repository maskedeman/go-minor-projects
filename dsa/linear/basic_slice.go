package main
import (
	"fmt"
)

func main(){
	var sli=[] int{0,4,6,8,5,54,89,74,98,15,16,778}
	sli= append(sli,10)
	fmt.Println("Capacity",cap(sli))
	fmt.Println("Length ",len(sli))
	var  value int
	var i int
	for i, value= range sli {
		fmt.Println("Index",i,"Range", value)
	}
}