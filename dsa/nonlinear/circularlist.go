package main
import(
	"fmt"
	"container/ring"
)
func main() {
	var integers []int
	integers = []int{2,5,65,185,1465,15,7,98,3,13}

	var circular_list *ring.Ring
	circular_list = ring.New(len(integers))

	var i int
	for i = 0; i < circular_list.Len(); i++ {
		circular_list.Value = integers[i]
		circular_list = circular_list.Next()
	}

	circular_list.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()

	for i = 0; i < circular_list.Len(); i++ {
		fmt.Print(circular_list.Value, ",")
		circular_list = circular_list.Prev()
	}
	fmt.Println()

	circular_list = circular_list.Move(2)
	circular_list.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()
}