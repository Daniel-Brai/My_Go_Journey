// UNDERSTANDING IOTA
// The iota keyword can be used to create a series of successive untyped integer constants without 
// needing to assign individual values to them. Here is an iota example: 
// ...
// const (
//  Watersports = iota
//  Soccer
//  Chess
// )
// ...

// This pattern creates a series of constants, each of which is assigned an integer value, starting at zero. 
// You can see examples of iota in Part 3

// Defining Multiple Constants with a Single Statement
// A single statement can be used to define several constants, as shown in Listing 4-10.

// Listing 4-10. Defining Multiple Constants in the main.go File in the basicFeatures Folder
package main

import (
 "fmt"
 //"math/rand"
)

func main() {
	const price, tax float32 = 275, 27.50
	const quantity, inStock = 2, true
	fmt.Println("Total:", quantity * (price + tax))
	fmt.Println("In stock: ", inStock)
}

// The const keyword is followed by a comma-separated list of names, an equal sign, and a commaseparated list of values, as illustrated by Figure 4-5. If a type is specified, all the constants will be created with 
// this type. If the type is omitted, then untyped constants are created, and each constant’s type will be inferred 
// from its value.

// Compiling and executing the code in Listing 4-10 produces the following output:
// Total: 605
// In stock: true


//  ## Revisiting Literal Values
// Untyped constants may seem like an odd feature, but they make working with Go a lot easier, and you will 
// find yourself relying on this feature, often without realizing, because literal values are untyped constants, 
// which means that you can use literal values in expressions and rely on the compiler to deal with mismatched 
// types, as shown in Listing 4-11. 
// Listing 4-11. Using a Literal Value in the £_iota.go File in the basicFeatures Folder

// package main

// import (
//  "fmt"
//  //"math/rand"
// )

// func main() {
// 	const price, tax float32 = 275, 27.50
// 	const quantity, inStock = 2, true
// 	fmt.Println("Total:", 2 * quantity * (price + tax))
// 	fmt.Println("In stock: ", inStock)
// }


// The highlighted expression uses the literal value 2, which is an int value as described in Table 4-4, 
// along with two float32 values. Since the int value can be represented as a float32, the value will be 
// converted automatically. When compiled and executed, this code produces the following output:
// Total: 1210
// In stock: true