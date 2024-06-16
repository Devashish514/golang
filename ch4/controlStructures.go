package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Ch4: Blocks, Shadows and control Structures")
	// Shadowing Variables...
	// shadowing variable means , variable that has the same name as a variable in a containing block...

	x := 10
	if x > 5 {
		fmt.Println(x)

		x := 5

		fmt.Println(x)
	}
	fmt.Println(x)

	// Universe block...
	// fmt.Println(true)
	// true := 10
	// fmt.Println(true)

	// if statement..
	// in go , if statement condition don't require to be inside ()

	n := rand.Intn(10)

	if n == 0 {
		fmt.Println("That's Too low:", n)
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// Go adds a ability to declare variables that are scoped to the condition and to both if and else blocks...

	if randomInt := rand.Intn(10); randomInt == 0 {
		fmt.Println("That's too low:", n)
	} else if n > 5 {
		fmt.Println("That's Too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	// Once the if else end, randomint is undefined....
	// fmt.Println(randomInt) // not defined

	// For Loops , four ways...
	// A complete for loop
	// we cannot use var to initialize a variable inside a for loop, use :=
	// we can shadow a variable inside for loop...
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Condition only for loop...
	// we remove the initialization and increment ...
	// similar to while loop from other languages...

	// i := 1

	// for i < 100 {
	// 	fmt.Println(i)
	// 	i *= 2
	// }

	// Infinite for statement...
	// here all three parts are removed...

	// for {
	// 	fmt.Println("Hello")
	// }

	// To get out of infinite loop , we can use break statement..

	// for {
	// 	if i > 100 {
	// 		break
	// 	}
	// 	fmt.Println("Hello")
	// 	i *= 2
	// }

	// continue : it is used to skip the current iteration...

	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 		continue
	// 	}
	// 	if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 		continue
	// 	}
	// 	if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for range loop...
	// used to iterate over elements of Go's built in types..
	// we can use it with strings, slices, maps and arrays...

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// first is the index and other is the value at that index..

	// Go requires to access all variables declared, even the one declared as part of a for loop.\
	// if you don't need to access the key/index , use an undeerscore as variable name...

	for _, v := range evenVals {
		fmt.Println(v)
	}

	// if you want the key/index but not the value, we can leave off the second variable..

	for i := range evenVals {
		fmt.Println(i)
	}

	uniqueNames := map[string]bool{
		"Fred":  true,
		"Raul":  true,
		"Wilma": true,
	}

	for k := range uniqueNames {
		fmt.Println(k)
	}

	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	// In above example, it prints the k,v of map, but the order is different for each outer loop...
	// this is actually a security feature....

	// Iterating over string..

	samples := []string{"hello", "apple_♪!"}

	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r)) // In the first column we have index, in the second numeric value of the letter, and third  numeric value of letter converted to string..
		}
		fmt.Println()
	}
	// for "apple_♪!" first column skips 7 and 8, second value at 6 is 9834 thats far larger than what can fit in byte, as we know strings were made out of bytes..
	// thats actually a special behavior from iterating over a string with a for-range loop.
	// whenever a for-range loop encounters a multibyte rune in a string, it converts the UTF-8 representation into a single 32-bit number and assign it to the value.
	// if for-range loop encounters a byte that doesn't represent a valid UTF-8 value , hex value 0xfffd is returned..

	// for-range value is a copy..
	// each time a for-range loop iterates over a compound type , it copies the value to the value variable, modification to value will not modify the value in compound type...

	oddVals := []int{3, 5, 7, 9, 11}

	for _, v := range oddVals {
		v *= 2
	}
	fmt.Println(oddVals)

	// Labels...

}
