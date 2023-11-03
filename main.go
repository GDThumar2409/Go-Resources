package main

import (
	"example/hello/mockdb"
	"fmt"
	"log"
	"math"
	"strconv"
)

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println("My Age is ", p.age)
}

type count int

func (c count) String() string {
	return fmt.Sprint("Count is ", strconv.Itoa(int(c)))
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func foo(i ...int) int {
	t := 0
	for _, v := range i {
		t += v
	}
	return t
}

func main() {
	fmt.Println("Hello World")

	//var c count = 1
	/*
		for i := 0; i < 10; i++ {
			fmt.Println("i is ", i)
		}
	*/

	/*
		x := []int{1, 2, 3, 4}

		for i, v := range x {
			fmt.Println(i, v)
		}
	*/

	/*
		m := map[string]int{
			"Gunjan": 24,
			"Nidhi":  22,
		}

		for k, v := range m {
			fmt.Println(k, v)
		}
	*/

	/*
		x := []string{"gunjan", "nidhi"}
		a := x
		b := make([]string, len(x))
		copy(b, x)
		x[1] = "gn"
		fmt.Println(x)
		fmt.Println(a)
		fmt.Println(b)

		b = append(b, a...)
		fmt.Println(b)
	*/

	//fmt.Println(c)

	fmt.Println(foo(1, 2, 3))

	p := person{
		name: "gunjan",
		age:  24,
	}

	p.speak()

	c := circle{
		radius: 2,
	}

	fmt.Println(info(c))

	s := square{
		length: 2,
		width:  4,
	}

	fmt.Println(info(s))

	db := mockdb.MockDatastore{
		Users: make(map[int]mockdb.User),
	}

	srvc := mockdb.Service{
		Ds: db,
	}

	u1 := mockdb.User{
		ID:    1,
		First: "James",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)

}
