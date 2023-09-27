package main
import (
"fmt"

)
func powerSeries(a int) (int, int){
	return a*a, a*a*a
}
func main(){
	var sq int
	var cu int
	sq, cu= powerSeries(5)
	fmt.Println("Square:",sq,"\nCube:",cu)
}