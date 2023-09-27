package main
import(
	"fmt"
)
func main(){
var arr= [10]int {0,12,54,14,98,65,145,498,56,20}  //array initialization
var i int
for i=0; i<len(arr);i++{
	fmt.Println("Elements in arr:",arr[i])
}
var value int
for i, value= range arr{  //index and value
	fmt.Println("range",value)
}

for _, value=range arr{  //index is ignored
	fmt.Println("No range", value)
}
}