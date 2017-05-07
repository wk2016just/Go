package main
import "fmt"
func main() {
   var numbers []int
   printSlice(numbers)


   numbers = append(numbers, 0)
   numbers = append(numbers, 1) //自动扩充cap
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)


   numbers1 := make([]int, len(numbers), (cap(numbers))*2)

   /* copy content of numbers to numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)  
   numbers1 = append(numbers, 2,3,4,2,2,2,2,2,2,2,2,2,2,2,2,2)
   printSlice(numbers1)  
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)}