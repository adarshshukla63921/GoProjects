package main 
import "fmt"

type Rectange struct{
	lenght int
	breadth int
}
func (r Rectange) area() int {
	return r.lenght*r.breadth
}
func main(){
	fmt.Println("Using structs in go.")
	var v Rectange=Rectange{10,20}
	fmt.Println(v.lenght)
	fmt.Println(v.breadth)
	fmt.Println(v)
	fmt.Println(v.area())
}