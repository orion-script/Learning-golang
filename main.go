/////////////// Lesson -2

// package main

// import "fmt"

// func main() {

// 	fmt.Println("Hello, ninjas!")

// }

/////////////// Lesson -3

// package main

// import "fmt"

// func main() {

// 	// string variables
// 	var nameOne string = "mario"
// 	var nameTwo = "luigi"
// 	var nameThree string

// 	fmt.Print(nameOne, nameTwo, nameThree)

// 	nameOne = "peach"
// 	nameThree = "bowser"

// 	fmt.Print(nameOne, nameTwo, nameThree)

// 	// the following is allowed inside functions only
// 	nameFour := "yoshi"
// 	fmt.Print(nameFour)

// 	// int variables
// 	var ageOne int = 20
// 	var ageTwo = 30
// 	ageThree := 40

// 	fmt.Print(ageOne, ageTwo, ageThree)

// 	// bits & memory
// 	// var numOne int8 = 25
// 	// var numTwo int8 = 128 // too large a number for 8-bit
// 	// var numTwo uint = -25 unsigned ints cannot be negative

// 	var scoreOne float32 = 25.98
// 	var scoreTwo float64 = 1965385877.5
// 	var scoreThree = 1.5 // inferred as float64

// 	fmt.Print(scoreOne, scoreTwo, scoreThree)

// 	// for more info see https://golang.org/ref/spec#Numeric_types

// }

/////////////// Lesson -4

// package main

// import "fmt"

// func main() {
// 	age := 35
// 	name := "shaun"

// 	// Print
// 	fmt.Print("hello, ")
// 	fmt.Print("world! \n")
// 	fmt.Print("new line \n")

// 	// Println
// 	fmt.Println("hello ninjas!")
// 	fmt.Println("goodbye ninjas!")
// 	fmt.Println("my age is", age, "and my name is", name)

// 	// Printf (formatted string), %_ = format specifier
// 	fmt.Printf("my age is %v and my name is %v \n", age, name) // %v = value in default format
// 	fmt.Printf("my age is %q and my name is %q \n", age, name) // %q = quotes
// 	fmt.Printf("age is of type %T \n", age)                    // %T is the type
// 	fmt.Printf("you scored %f points! \n", 225.55)             // %f = float format
// 	fmt.Printf("you scored %0.1f points! \n", 225.55)          // %0.2f = float with 2 decimal points

// 	// Sprintf (save formatted strings)
// 	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
// 	fmt.Println("the saved string is:", str)

// 	// see more format specifiers here:
// 	// https://golang.org/pkg/fmt/

// }

/////////////// Lesson -5

// package main

// import "fmt"

// func main() {
// 	// var ages [3]int = [3]int{20, 25, 30}
// 	var ages = [3]int{20, 25, 30}

// 	names := [4]string{"yoshi", "mario", "peach", "bowser"}
// 	names[1] = "luigi"

// 	fmt.Println(ages, len(ages))
// 	fmt.Println(names, len(names))

// 	// slices (use arrays under the hood)
// 	var scores = []int{100, 50, 60}
// 	scores[2] = 25
// 	scores = append(scores, 85)

// 	fmt.Println(scores, len(scores))

// 	// slice ranges
// 	rangeOne := names[1:4] // doesn't include pos 4 element
// 	rangeTwo := names[2:]  //includes the last element
// 	rangeThree := names[:3]

// 	fmt.Println(rangeOne, rangeTwo, rangeThree)
// 	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

// 	rangeOne = append(rangeOne, "koopa")
// 	fmt.Println(rangeOne)

// }

/////////////// Lesson -6

// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// func main() {
// 	greeting := "hello there friends!"

// 	fmt.Println(strings.Contains(greeting, "hello"))
// 	fmt.Println(strings.Contains(greeting, "bye"))
// 	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
// 	fmt.Println(strings.ToUpper(greeting))
// 	fmt.Println(strings.Index(greeting, "ll"))
// 	fmt.Println(strings.Split(greeting, " "))

// 	// the original value is unchanged
// 	fmt.Println("original string value =", greeting)

// 	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

// 	sort.Ints(ages)
// 	fmt.Println(ages)

// 	index := sort.SearchInts(ages, 30)
// 	fmt.Println(index)

// 	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

// 	sort.Strings(names)
// 	fmt.Println(names)

// 	fmt.Println(sort.SearchStrings(names, "bowser"))
// }

/////////////// Lesson -7

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	x := 0
// 	for x < 5 {
// 		fmt.Println("value of x is:", x)
// 		x++ // infinite loop without this
// 	}

// 	for i := 0; i < 5; i++ {
// 		fmt.Println("value of i is:", i)
// 	}

// 	names := []string{"mario", "luigi", "yoshi", "peach"}

// 	for i := 0; i < len(names); i++ {
// 		fmt.Println(names[i])
// 	}

// 	for index, val := range names {
// 		fmt.Printf("the value at pos %v is %v \n", index, val)
// 		val = "new string"
// 	}

// 	for _, val := range names {
// 		fmt.Print(val, ",")
// 		val = "new string"
// 	}

// 	// changing val in a loop does not mutate the original slice
// 	fmt.Println(names)

// }

/////////////// Lesson -8

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	age := 45

// 	fmt.Println(age <= 50)
// 	fmt.Println(age >= 50)
// 	fmt.Println(age == 45)
// 	fmt.Println(age != 50)

// 	if age < 30 {
// 		fmt.Println("age is less than 30")
// 	} else if age < 40 {
// 		fmt.Println("age is less than 40")
// 	} else {
// 		fmt.Println("age is not less than 40")
// 	}

// 	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

// 	for index, val := range names {
// 		if index == 1 {
// 			fmt.Println("continuing at pos", index)
// 			continue
// 		}
// 		if index > 2 {
// 			fmt.Println("breaking at pos", index)
// 			break
// 		}
// 		fmt.Printf("the value at pos %v is %v \n", index, val)
// 	}
// }


/////////////// Lesson -9