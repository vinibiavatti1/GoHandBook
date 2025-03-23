// Map
// ...

package datatypes

import "fmt"

// Declaring Maps
func DeclaringMaps() {

	// Declaring an Empty Map
	mp1 := map[string]int{}
	fmt.Println(mp1) // Output: map[]

	// Declaring a Map with Predefined Values
	mp2 := map[string]int{
		"A": 1,
		"B": 2,
	}
	fmt.Println(mp2) // Output: map[A:1 B:2]
}

// Manipulating Maps
// Demonstrates common operations on maps.
func ManipulatingMaps() {

	// Declaring a Map
	mp := map[string]int{
		"A": 1,
		"B": 2,
	}

	// Acessing Elements by Key
	el := mp["A"]
	fmt.Println(el) // Output: 1

	// Getting Length
	ln := len(mp)
	fmt.Println(ln) // Output: 2

	// Adding Elements
	mp["C"] = 3
	fmt.Println(mp) // Output: map[A:1, B:2, C:3]

	// Updating Elements
	mp["A"] = 9
	fmt.Println(mp) // Output: map[A:9, B:2, C:3]

	// Deleting Elements
	// We can use the built-in function "delete()" to delete keys from a map
	delete(mp, "A")
	fmt.Println(mp) // Output: map[B:2, C:3]
}

// Iterating Over Maps
// We can use the for loop to iterate over maps.
func IteratingOverMaps() {

	// Declaring a Map
	mp := map[string]int{
		"A": 1,
		"B": 2,
	}

	// Iterating over a Slice
	for k, v := range mp {
		fmt.Println(k, v) // Output: A 1, B 2
	}

	// Iterating over a Slice (Only Keys)
	for k := range mp {
		fmt.Println(k) // Output: A, B
	}

	// Iterating over a Slice (Only Values)
	for _, v := range mp {
		fmt.Println(v) // Output: 1, 2
	}
}
