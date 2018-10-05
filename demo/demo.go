package stack

import (
	"fmt"
	"math"
)

type name struct {
	first string
	last  string
}

var (
	he  = name{"wang", "bo"}
	she = name{"wang", "shuo"}
	p   = &name{"wang", "yuan"}
)

func add(x int, y int) (int, int) {
	return y, x
}
func main() {

	var map1 map[string]int
	map1["one"] = 1

	map1["two"] = 2
	//var map2  =make(map[string]int)
	fmt.Println(map1["two"])

	//fmt.Println(myname.first,myname.last)
	a := 10
	b := 11.1
	c := 19
	//fmt.Println("sss")
	fmt.Println(math.Sqrt(b))
	//fmt.Println(a)
	add(a, c)
	//fmt.Println(add(a,c))
	p := &a
	*p = 21
	//fmt.Println(a)
}
