package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

func (p Person) Info() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func (p *Person) Birthday() {
	p.age++
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func printAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Площадь: %.2f\n", shape.Area())
	}
}

type Stringer interface {
	String() string
}

type Book struct {
	title  string
	author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book Title: %s, Author: %s", b.title, b.author)
}

func main() {

	person := Person{name: "Oleg", age: 30}
	person.Info()

	person.Birthday()
	fmt.Printf("After birthday: ")
	person.Info()

	circle := Circle{radius: 5}
	fmt.Printf("Circle area: %.2f\n", circle.Area())

	rect := Rectangle{width: 4, height: 6}
	fmt.Printf("Rectangle area: %.2f\n", rect.Area())

	shapes := []Shape{circle, rect}
	printAreas(shapes)

	book := Book{title: "Go Programming", author: "John Doe"}
	fmt.Println(book.String())
}
