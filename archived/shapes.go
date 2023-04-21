package main

// import "fmt"

// type square struct {
// 	sideLength float64
// }

// type triangle struct {
// 	height float64
// 	base   float64
// }

// type shape interface {
// 	getArea() float64
// }

// func (s square) getArea() float64 {
// 	return s.sideLength * s.sideLength
// }

// func (t triangle) getArea() float64 {
// 	return 0.5 * t.base * t.height
// }

// func printArea(s shape) {
// 	fmt.Println(s.getArea())
// }

// func shapeMain() {
// 	s := square{2}
// 	t := triangle{2, 3}

// 	printArea(s)
// 	printArea(t)
// }

// type cat struct{}

// type dog struct{}

// // All Pets have ability to MakeSound
// type pet interface {
// 	makeSound() string
// }

// func (c cat) makeSound() string {
// 	return "meow"
// }

// func (d dog) makeSound() string {
// 	return "woof"
// }

// func print(p pet) {
// 	fmt.Println(p.makeSound())
// }

// func checkAnimals() {
// 	c := cat{}
// 	d := dog{}

// 	pets := []pet{c, d}
// 	for _, p := range pets {
// 		print(p)
// 		fmt.Println(p.makeSound())
// 	}

// }
