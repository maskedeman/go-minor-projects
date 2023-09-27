package main
import(
	"fmt"
)
func main(){
	var lang= map[int] string{
	
		0:"Nepali",
		1:"English",
		2:"Deutsch",
	}
	var i int
	var value string
	for i, value= range lang{
			fmt.Println("index:",i,"value:",value )
	}
	var states=make(map[int]string,3)
	states[1]="Bagmati"
	states[2]="Karnali"
	states[3]="Lumbini"
	delete (states , 2)// delete key from the map
	fmt.Println("states",states)
	//delete(states, "Lumbini")
	//fmt.Println("states after deleting Lumbini",states)
	fmt.Println("State at 1", states[1])

}
