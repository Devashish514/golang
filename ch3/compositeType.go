package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ch3 Composite Types...")

	// Arrays..
	var arr [3]int // zero value array
	fmt.Println(arr)

	var arr2 = [4]int{1, 2, 3, 4}
	fmt.Println(arr2)

	// Sparse array...
	var sparse = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("sparse array", sparse)

	// you can pass  ..., instead of number to initialize array
	var arr3 = [...]int{1, 2, 3, 4}
	fmt.Println(arr3)

	// arrays are comparable..
	fmt.Println("array equality comparison: ", arr2 == arr3)

	// Arrays in Go are one dimensional, but can simulate multi-dimensional..
	// This declare `multiDimenArr` to be a array of length 2 whose type is an array of ints of length 3...

	var multiDimenArr [2][3]int
	fmt.Println("multi Dimensional array", multiDimenArr) // [[0 0 0] [0 0 0]]

	// reading writing arrays
	arr3[0] = 10
	fmt.Println(arr3[0])

	// getting length of Array
	fmt.Println(len(arr3))

	// In Arrays, size is the part of the type , means [3]int is differnet type from [4]int

	// Slices...
	// In slice we don't specify the size when we declare it..

	var sliceX = []int{1, 2, 3, 4}
	fmt.Println(sliceX)

	var nilSlice []int // zero value for slice is `nil`
	fmt.Println("nilSlice", nilSlice)

	fmt.Println("Zero Value Slice", nilSlice == nil)

	// nil is different from null , it represents lack of value for some types...it has no type.

	//empty slice
	var emptySlice = []int{}
	fmt.Println("emptySlice", emptySlice)

	fmt.Println("empty slice", emptySlice == nil)

	// Slices are not comparable.... its compile error to compare slices..

	// Built-in Functions...
	// append

	emptySlice = append(emptySlice, 10)
	emptySlice = append(emptySlice, 20, 30, 40)

	fmt.Println(emptySlice)

	var intSlice = []int{50, 60, 70, 80}

	// append one slice to another using ...
	emptySlice = append(emptySlice, intSlice...)
	fmt.Println(emptySlice)

	// Go is call by value..

	// capacity
	fmt.Println("capacity", cap(emptySlice))

	fmt.Println("capacity of Array", cap(arr))

	//make
	// this will create a slice of length 0, cap 10
	newSlicewithCap := make([]int, 0, 10)

	newSlicewithCap = append(newSlicewithCap, 10)
	fmt.Println(newSlicewithCap)

	// Slicing Slices..

	var bigSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 0}
	slice1 := bigSlice[1:4]
	slice2 := bigSlice[4:]
	slice3 := bigSlice[:6]
	slice4 := bigSlice[:]

	fmt.Println("bigSlice:", bigSlice)
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)

	// Slices share memory sometimes...
	// This means , chnages to an element in a slice affect all slices that share that element...

	bigSlice[0] = 20
	slice2[3] = 10
	slice3[1] = 100
	fmt.Println("bigSlice", bigSlice)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)

	// changes to bigSlice modifies all subslices

	// slicing becomes more confusing when combined with append()

	var aSlice = []int{1, 2, 3, 4}
	var bSlice = aSlice[:2]
	fmt.Println(cap(aSlice), cap(bSlice))

	bSlice = append(bSlice, 30)

	fmt.Println("aSlice", aSlice)
	fmt.Println("bSlice", bSlice)

	var cSlice = make([]int, 0, 5)
	cSlice = append(cSlice, 1, 2, 3, 4)

	dSlice := cSlice[:2]
	eSlice := cSlice[2:]

	fmt.Println(cap(cSlice), cap(dSlice), cap(eSlice))

	dSlice = append(dSlice, 30, 40, 50)
	cSlice = append(cSlice, 60)
	eSlice = append(eSlice, 70)

	fmt.Println("cSlice", cSlice)
	fmt.Println("dSlice", dSlice)
	fmt.Println("eSlice", eSlice)

	// To avoid this complicated behavior , either never use append on sliced slice or use full slice expression...
	// full slice expression includes third part, which indicates the last position in the parent's slice's capacity..

	fSlice := cSlice[:2:2]
	gSlice := cSlice[2:4:4]

	fSlice = append(fSlice, 30, 40, 50)
	cSlice = append(cSlice, 60)
	gSlice = append(gSlice, 70)

	fmt.Println("cSlice", cSlice)
	fmt.Println("fSlice", fSlice)
	fmt.Println("gSlice", gSlice)

	// Converting Arrays to SLices..

	anArr := [4]int{1, 2, 3, 4}
	anArrSlice := anArr[:]
	fmt.Println(anArrSlice)

	// copy function
	// used to create slices, independent of original
	// takes two parameters , destination and source copy(d,s)

	originalSlice := []int{1, 2, 3, 4}
	copySlice := make([]int, 4)

	num := copy(copySlice, originalSlice)
	fmt.Println("num", num)

	copySlice[0] = 10
	fmt.Println("copiedSlice", copySlice)
	fmt.Println("originalSlice", originalSlice)

	// copy function returns the number of values copied, if required we can assign it a variable , otherwise we can directly use it.

	// if you dont need to copy entire slice
	copySlice2 := make([]int, 2)
	copy(copySlice2, originalSlice) // copy only 2 elements , as slice has length of 2
	fmt.Println(copySlice2)

	// middle of source slice...
	copySlice3 := make([]int, 2)
	copy(copySlice3, originalSlice[2:])
	fmt.Println(copySlice3)

	// we can copy array by making slice of arrays..

	// Maps
	// syntax map[keyType]valueType
	// zero value for a map is nil
	// Maps are not comparable..

	var nilMap map[int]string
	fmt.Println("nil_Map", nilMap)

	// on read to nilMap always return zero value for the map's value type, writing to a nil map cause panic

	// nilMap[1] = "le" // throw error - panic: assignment to entry in nil map

	totalWins := map[string]int{} // empty Map
	// we can read and write to a empty map , its different from nil map....

	teams := map[string]int{
		"RR":  1,
		"CSK": 5,
	}

	totalWins["KKR"] = 1

	fmt.Println(teams, totalWins)
	fmt.Println(totalWins["KKR"])

	// we can use `make` to create maps with a default size...
	// maps created with make still have length 0 and can grow past initially specified size..

	ages := make(map[int]string, 10)
	fmt.Println("Map Length ", len(ages)) // 0

	// Keys for a map can be any comparable type...

	fmt.Println(totalWins["RCB"]) // output 0 , even if no such key in the totalWins Map it return zero value for the value type..

	// comma ok idiom...
	// when we need to find if key exist or not in map , we use comma ok idiom...

	m := map[string]int{
		"hello": 0,
		"world": 5,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok) // 0 , true

	v, ok = m["good bye"]
	fmt.Println(v, ok) // 0, false...

	// here v has the value , and ok is boolean, represent whether the value exist or not,....

	// Deleting from maps

	delete(m, "world")

	fmt.Println(m)

	// usings Maps as Sets..
	// Go doesnt have Sets , we can simulate it using Maps

	intSet := map[int]bool{}

	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(intSet[5])
	fmt.Println(intSet[100])
	// you cannot have duplicate keys in a map..

	// Structs...

	type Person struct {
		name string
		age  int
		pet  bool
	}

	var fred Person = Person{name: "fred", age: 10, pet: false}
	fmt.Println(fred)

	// zero value struct has every field set to the fields zero value..
	bob := Person{}
	fmt.Println(bob)

	julia := Person{
		"julia",
		40,
		true,
	}

	fmt.Println(julia)

	//anonymous structs...

	pet := struct {
		name string
		kind string
	}{
		name: "fido",
		kind: "dog",
	}

	fmt.Println("anonymous struct", pet)

}
