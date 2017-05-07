package main
import  "fmt"
 

func main() {
var numbers []int /* a slice of unspecified size */
/* numbers == []int{0,0,0,0,0}*/
numbers = make([]int,5,5) /* a slice of length 5 and capacity 5*/

numbers2 := make([]int,5,5)
fmt.Println(numbers)
fmt.Println(numbers2)
}