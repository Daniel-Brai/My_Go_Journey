// Using Constants
// Constants are names for specific values, which allows them to be used repeatedly and consistently. There 
// are two ways to define constants in Go: typed constants and untyped constants. Listing 4-7 shows the use of 
// typed constants. 
// Listing 4-7. Defining Typed Constants in the 2_constants.go File in the cosntants Folder
package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	const price float32 = 275.00
	const tax float32 = 27.50
	fmt.Println(price + tax)
}

// Typed constants are defined using the const keyword, followed by a name, a type, and a value 
// assignment, as illustrated below

// const price float32 = 275.00

//  where const = keyword, 
			// price = Name
			// float32 = Type
			// 275.00 = value Assignment

// This statement creates a float32 constant named price whose value is 275.00. The code in Listing 4-7
// creates two constants and uses them in an expression that is passed to the fmt.Println function. Compile 
// and run the code, and you will receive the following output

// 302.5

//  ## Understanding Untyped Constants
// Go has strict rules about its data types and doesn’t perform automatic type conversions, which can 
// complicate common programming tasks, as Listing 4-8 demonstrates. 
// Listing 4-8. Mixing Data Types in the 2_constants.go File in the basicFeatures Folder

// package main

// import (
//  "fmt"
//  //"math/rand"
// )

// func main() {
// 	const price float32 = 275.00
// 	const tax float32 = 27.50
// 	const quantity int = 2
// 	fmt.Println("Total:", quantity * (price + tax))
// }

// The new constant’s type is int, which is an appropriate choice for a quantity that can represent only a 
// whole number of products, for example. The constant is used in the expression passed to the fmt.Println
// function to calculate a total price. But the compiler reports the following error when the code is compiled:
// .\2_constants.go:12:26: invalid operation: quantity * (price + tax) (mismatched types int and 
// float32)

// Most programming languages would have automatically converted the types to allow the expression 
// to be evaluated, but Go’s stricter approach means that int and float32 types cannot be mixed. The 
// untyped constant feature makes constants easier to work with because the Go compiler will perform limited 
// automatic conversion, as shown in Listing 4-9.
// Listing 4-9. Using an Untyped Constant in the main.go File in the basicFeatures Folder

// package main

// import (
//  "fmt"
//  //"math/rand"
// )

// func main() {
// 	const price float32 = 275.00
// 	const tax float32 = 27.50
// 	const quantity = 2
// 	fmt.Println("Total:", quantity * (price + tax))
// }

// An untyped constant is defined without a data type, as illustrated below

// const quantity = 2

// where const = keyword,
		// quantity = variable name
		// 2 = variable anme assignment

// Omitting the type when defining the quantity constant tells the Go compiler that it should be more 
// flexible about the constant’s type. When the expression passed to the fmt.Println function is evaluated, 
// the Go compiler will convert the quantity value to a float32. Compile and execute the code, and you will 
// receive the following output:

// Total: 605