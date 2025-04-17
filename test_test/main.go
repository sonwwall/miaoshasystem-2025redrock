package main

import "fmt"

type test struct {
}

func (t *test) print() string {
	return fmt.Sprint("hello")
}

func testprint() *test {
	return &test{}
}

func main() {
	a := testprint()
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := testprint().print
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	c := testprint().print()
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
