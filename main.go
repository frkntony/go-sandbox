// frkntony 16th, feb 2020
package main

import (
	"fmt"
	"github.com/frkntony/gopg/explodelib"
	"math"
	"net/http"
	"strconv"
)

// section 4
func greetings(name string) string {
	return "Hello, " + name
}

// section 4
func getSum(num1 int, num2 int)  int {
	return num1 + num2
}

// section 11
func adder () func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

// section 12
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// section 12
// value receiver
func (p Person) Greet() string {
	return "Hello, " + p.firstName + ". Your age is: " + strconv.Itoa(p.age)
}

// section 12
// pointer receiver
func (p *Person) hasBirthday() {
	 p.age++
}

// section 13
type Shape interface {
	area() float64
}

// section 13
type Circle struct {
	x, y, radius float64
}

// section 13
type Rectangle struct {
	width, height float64
}

// section 13
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// section 13
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// section 13
func getArea(s Shape) float64  {
	return s.area()
}

// section 14
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>First GO REST API</h1>")
}

// section 14
func testRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>TEST</h1>")
}


/** used to separate sections in stdout */
var separator = "*******************************************************"

func main() {

	///////////////////////////////////////////////////////////
	/* section 1 - hello world */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 1 - Hello World!")
	fmt.Println("Hello World!")
	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 2 - variables */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 2 - variables")

	// var declaration
	var sec2Name = "Tony"
	var age = 23
	var anotherSize float32 = 2.3
	fmt.Printf("My name is %s and I'm %d y/o\n", sec2Name, age)
	fmt.Printf("%T\n", anotherSize) // %T - means type

	// shorthand
	sec2Hobby := "programming go"
	sec2length := 3.5 // default float64
	fmt.Printf("I'm %s for %f hours\n", sec2Hobby, sec2length)

	// const declaration
	const isCool = true
	fmt.Printf("Golang is cool, do you agree? Say true for 'yes': %v\n", isCool)

	fmt.Println(separator)
	///////////////////////////////////////////////////////////
	/* section 3 - packages */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 3 - packages")
	fmt.Println(explodelib.Reverse("srehpog ollay"))
	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 4 - functions */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 4 - functions")
	fmt.Println(greetings("Tony"))
	fmt.Println(getSum(1,5))
	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 5 - arrays */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 5 - arrays")

	// declare
	var fruitArr [2]string
	// assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr[1])

	// fruitArrayTwo assign values at declaration without array size
	fruitArrTwo := []string{"Apple", "Orange", "Grape", "Watermelon"}
	fmt.Println(fruitArrTwo)

	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 6 - conditionals */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 6 - conditionals")

	sec6x := 10
	sec6y := 10

	// if else else if
	if sec6x < sec6y {
		fmt.Printf("%d is less than %d\n", sec6x, sec6y)
	} else if sec6x == sec6y {
		fmt.Println("equals")
	} else {
		fmt.Println("whoops!")
	}

	// switch
	switch sec6y {
	case 10:
		fmt.Printf("It's a %d\n", sec6y)
		break
	default:
		fmt.Println("Nothing seemed to work")
		break
	}

	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 7 - loops */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 7 - loops")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// default loop
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	// FizzBuzz
	for i := 1; i <= 30; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
 		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 8 - maps */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 8 - maps")

	// map[string] - key data type, string - value data type
	emails := make(map[string]string)

	emails["Adam"] = "Adam@gmail.com"
	emails["Carol"] = "Carol@gmail.com"
	emails["Tony"] = "Tony@gmail.com"

	// map for key tony
	fmt.Println(emails["Tony"])

	// delete element
	delete(emails, "Tony")
	fmt.Println(emails)

	// assign at declare
	emailsTwo := map[string]string{"John":"John@seneca.com", "Blair":"Blair@seneca.com"}
	emailsTwo["Simon"] = "Simon@seneca.com"
	fmt.Println(emailsTwo)

	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 9 - range */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 9 - range")

	ids := []int{33,76,54,23,11,2}

	// index + element
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// no index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	// sum ids
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum:", sum)

	// Range with map
	for k, v := range emails {
		fmt.Printf("Key: %s, value: %s\n", k, v)
	}

	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 10 - pointers */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 10 - pointers")

	a := 5
	b := &a // memory address

	fmt.Println(a, b)

	// pointer
	fmt.Println(a + *b)

	// change val with pointer
	*b = 15
	fmt.Println(a)

	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 11 - closures */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 11 - closures")

	closureSum := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(closureSum(i))
	}

	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 12 - structs */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 12 - structs")

	// init a person
	person1 := Person{firstName: "Adam", lastName: "Loveless", city: "Moscow", gender: "m", age: 23}
	fmt.Println(person1)

	// alternative
	person2 := Person{"Carol", "Brox", "Moscow", "m", 23}
	fmt.Println(person2)

	fmt.Println(person1.Greet())
	person1.hasBirthday()
	fmt.Println(person1.Greet())

	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 13 - interfaces */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 13 - interfaces ")

	circle := Circle{0, 0,5}
	rectangle := Rectangle{10, 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))

	fmt.Println(separator)

	///////////////////////////////////////////////////////////
	/* section 14 - web */
	///////////////////////////////////////////////////////////
	fmt.Println("Section 14 - web ")
	http.HandleFunc("/", index)
	http.HandleFunc("/test", testRouter)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}

