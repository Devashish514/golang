package main

import "fmt"

// x:=0 // Invalid
// var packageLevelVar int = 10 // allowed

func main() {
	fmt.Println("chapter 2: Primitive Types..")
	// Integers..
	var num1 int = 100000
	var num2 int = 10_000_000 // To make longer integers readable, we can put underscores in the middle

	fmt.Println(num1)
	fmt.Println(num2)

	var num3 int
	fmt.Println("Zero Value Integer:", num3) // zero value for integer is 0

	// Special Integer Type (byte)
	// a byte is an alias for uint8;

	var numByte byte = 255
	fmt.Println("Byte:", numByte)

	//Integer Operator..
	var val1 int = 10
	var val2 int = 3

	fmt.Println("Multiply Operator", val1*val2)
	fmt.Println("Divide", val1/val2) // Return Integer, if want float answer , then convert the integers to float;
	fmt.Println("Divide", float64(val1)/float64((val2)))

	// Divide by 0 , causes panic...
	// fmt.Println(val1 / 0) //error..

	var x int = 10
	x *= 2
	fmt.Println("x", x)
	x += 2
	fmt.Println("x", x)
	x -= 2
	fmt.Println("x", x)
	x /= 2
	fmt.Println("x", x)

	// Comparison operator...
	if x >= val1 {
		fmt.Println("X is greater than val1")
	}

	// Floating Point Integer...
	// Zero value for float is also 0
	var type1 float32
	var type2 float64

	fmt.Println("type1", type1)
	fmt.Println("type2", type2)

	type1 = -3.1415
	type2 = -3.1415

	fmt.Println(type1 == float32(type2))

	// Strings and Runes...

	// rune type is an alias for int32
	var myRune rune = '😊'

	// Print each Unicode character as a rune
	fmt.Printf("Rune: %c\n", myRune)

	// %c verb to display it as a Unicode character, and we print its Unicode code point using the %U verb.

	// Print the Unicode code point of emoji character
	fmt.Printf("Code point: %U\n", myRune)

	var s string
	fmt.Println("Zero Value of String:", s)

	s = "hello world"

	// String are comparable
	var ss string = "bye"
	fmt.Println("String Comparison", s != ss)

	// Concatenating Strings....

	var s1 = "Good"
	var s2 = "Evening!"

	fmt.Println(s1[0]) //prints byte value for G = 71
	// If wants to print G
	fmt.Println("string slice", s1[:1])

	fmt.Println(s1 + "\n" + s2)

	// In Go, Strings are immutable, we can reassign value but cannot change it

	// s1[0] = "g" // error
	// fmt.Println(s1)

	// In Go, a non zero number or string is not boolean true, in fact no other type can be interpreted to a bool , explicitly or implicitly/....
	// we can use comparison operators to  get bool..
	if s1 == "Good" {
		fmt.Println("Its Good")
	}

	// Declaring Variables....
	// // Most verbose way..

	// var y int = 10
	// var z = 10
	// var m int // assigned with zero value

	// Declaring Multiple variables..

	// Same type..
	var a, b int = 10, 20
	fmt.Println(a, b)

	// all zero value of same type
	var c, d string
	fmt.Println(c, d)

	// Different Type
	var e, f = 10, "Hello"
	fmt.Println(e, f)

	// Declaration List...
	var (
		one          int
		two                 = 20
		three        string = "Str"
		alpha, beeta        = 12, 14
		g, h         bool
	)

	fmt.Println(one, two, three, alpha, beeta, g, h)

	// Similar to var , we can declare multiple variable with :=

	m1, m2 := "John", 12
	fmt.Println(m1, m2)

	x1 := 10
	x1, y1 := 20, "Hello"

	fmt.Println(x1, y1)

	// Const...
	const KEY int = 10

	const (
		ifKey   = "id"
		nameKey = "name"
	)

	fmt.Println(KEY)
	fmt.Println(ifKey, nameKey)

	// KEY = 20 // error

	// Typed and Untyped Constants..
	const typedX int = 65

	const untypedX = 10
	fmt.Println(typedX, untypedX)

}
