package main
import(
	"fmt"
)
func main(){
var TwoDArray [5][5] int //fixed allocation
TwoDArray[2][3]=10
TwoDArray[4][3]=12
fmt.Println(TwoDArray)

var rows int //dynamic allocation
var cols int
rows=8
cols=4
var twodslices = make([][]int, rows)
var i int
for i=range twodslices{
	twodslices[i]=make([]int ,cols )
}
var arr = [] int{5,6,7,8,9}
var slic1 = arr[: 3] //create slice from array 0-3 index
fmt.Println("slice1",slic1)
var slic2 = arr[1:5]//create slice from array 1-5 index
fmt.Println("slice2",slic2)
var slic3 = append(slic2, 12) //append method
fmt.Println("slice3",slic3)
}
