package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// default will be 0
	var x1 int
	fmt.Println(x1)

	var x2 int = 5
	var y2 int = 7
	var sum2 int = x2 + y2
	fmt.Println(sum2)

	x3 := 5
	y3 := 7
	sum3 := x3 + y3
	fmt.Println(sum3)

	x4 := 5
	if x4 > 4 {
		fmt.Println("More than 4 do sth")
	} else if x4 < 2 {
		fmt.Println("Less than 2 do sth")
	} else {
		fmt.Println("Else do sth")
	}

	// array
	var a1 [5]int
	a1[2] = 7
	fmt.Println(a1)

	a2 := [5]int{1, 3, 5, 2, 6}
	fmt.Println(a2)

	// slice
	s1 := []int{1, 3, 5, 2, 6}
	s1 = append(s1, 999)
	fmt.Println(s1)

	// map
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 5
	fmt.Println(vertices)
	fmt.Println(vertices["square"])
	delete(vertices, "square")
	fmt.Println(vertices)

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// array range
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	// map range
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}

	// call func sum4
	result := sum4(9, 7)
	fmt.Println(result)

	// call func sqrt
	result1, err1 := sqrt(16)
	// result1, err1 := sqrt(-16)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

	// create a struct
	p := person{name: "John", age: 10}
	fmt.Println(p)
	fmt.Println(p.age)

	// pointers
	ip1 := 7
	fmt.Println(ip1)
	// memory address
	fmt.Println(&ip1)

	ip2 := 7
	inc(ip2)
	fmt.Println(ip2)

	ip3 := 7
	incp(&ip3)
	fmt.Println(ip3)
}

func sum4(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func inc(x int) {
	x++
}

func incp(x *int) {
	*x++
}

type person struct {
	name string
	age  int
}
