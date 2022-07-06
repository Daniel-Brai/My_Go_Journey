// USING VARIABLES
// Variables are defined using the var keyword, and, unlike constants, the value assigned to a variable can be 
// changed, as shown in Listing 4-12. 
// Listing 4-12. Using Constants in the 4_variables.go File in the 4_variables Folder
package main

import "fmt"

func main() {
	var price float32 = 275.00
	var tax float32 = 27.50
	fmt.Println(price + tax)
	price = 300
	fmt.Println(price + tax)
}

// Variables are declared using the var keyword, a name, a type, and a value assignment, as illustrated in 

//  Defining a Variable
// Listing 4-12 defines price and tax variables, both of which are assigned float32 values. A new value 
// is assigned to the price variable using the equal sign, which is the Go assignment operator, as illustrated in 
// Figure 4-7. (Notice that I can assign the value 300 to a floating-point variable. This is because the literal value 
// 300 is an untyped constant that can be represented as a float32 value.)

// Figure 4-7. Assigning a new value to a variable
// price = 300 where price = name, 300 = value
// The code in Listing 4-12 writes two strings to the standard out using the fmt.Println function, 
// producing the following output when the code is compiled and executed:
// 302.5
// 327.5


// ## Omitting the Variable’s Data Type
// The Go compiler can infer the type of variables based on the initial value, which allows the type to be 
// omitted, as shown in Listing 4-13.
// Listing 4-13. Omitting a Variable’s Type in the 4_variables.go File in the 4_variableseatures Folder

// package main

// import "fmt"

// func main() {
// 	var price = 275.00
// 	var price2 = price
// 	fmt.Println(price)
// 	fmt.Println(price2)
// }

// The variable is defined using the var keyword, a name, and a value assignment, but the type is omitted, 
// as illustrated by Figure 4-8. The value of the variable can be set using a literal value or the name of a constant 
// or another variable. In the listing, the value of the price variable is set using a literal value, and the value of 
// price2 is set to the current value of price.

// Figure 4-8. Defining a variable without specifying a type
// var price = 275.00 where var = keyword, price = Name and 275.00 = value

// The compiler will infer the type from the value assigned to the variable. The compiler will inspect the 
// literal value assigned to price and infer its type as float64, as described in Table 4-4. The type of price2 will 
// also be inferred as float64 because its value is set using the price value. The code in Listing 4-13 produces 
// the following output when compiled and executed:

// 275
// 275

// Omitting a type doesn’t have the same effect for variables as it does for constants, and the Go compiler 
// will not allow different types to be mixed, as Listing 4-14 shows.
// Listing 4-14. Mixing Data Types in the 4_variables.go File in the 4_variables Folder

// package main

// import "fmt"

// func main() {
// 	var price = 275.00
// 	var tax float32 = 27.50
// 	fmt.Println(price + tax)
// }

// The compiler will always infer the type of literal floating-point values as float64, which doesn’t match 
// the float32 type of the tax variable. Go’s strict type enforcement means that the compiler produces the 
// following error when the code is compiled:
// .\4_variables.go:10:23: invalid operation: price + tax (mismatched types float64 and float32)

// To use the price and tax variables in the same expression, they must have the same type or be 
// convertible to the same type. I explain the different ways types can be converted in Chapter 5.


// Omitting the Variable’s Value Assignment
// Variables can be defined without an initial value, as shown in Listing 4-15.
// Listing 4-15. Defining a Variable Without an Initial Value in the 4_variables.go File in the 4_variables Folder
// package main

// import "fmt"

// func main() {
// 	var price float32
// 	fmt.Println(price)
// 	price = 275.00
// 	fmt.Println(price)
// }


// Variables are defined using the var keyword followed by a name and a type, as illustrated by Figure 4-9. 
// The type cannot be omitted when there is no initial value.
// var price float32

// Variables defined this way are assigned the zero value for the specified type, as described in Table 4-5.
// Table 4-5. The Zero Values for the Basic Data Types
// Type			Zero Value
// int 			0
// uint 			0
// byte 			0
// float64 		0
// bool 			false
// string “ ” 		(the empty string)
// rune 			0

// The zero value for numeric types is zero, which you can see by compiling and executing the code. The 
// first value displayed in the output is the zero value, followed by the value assigned explicitly in a subsequent 
// statement:
// 0
// 275


// Defining Multiple Variables with a Single Statement
// A single statement can be used to define several variables, as shown in Listing 4-16.
// Listing 4-16. Defining Variables in the 4_variables.go File in the 4_variables Folder

// package main

// import "fmt"

// func main() {
// 	var price, tax = 275.00, 27.50
// 	fmt.Println(price + tax)
// }

// This is the same approach used to define constants, and the initial value assigned to each variable is 
// used to infer its type. A type must be specified if no initial values are assigned, as shown in Listing 4-17, and 
// all variables will be created using the specified type and assigned their zero value.
// Listing 4-17. Defining Variables Without Initial Values in the 4_variables.go File in the 4_variables Folder
package main

import "fmt"

func main() {
	var price, tax float64
	price = 275.00
	tax = 27.50
	fmt.Println(price + tax)
}

// Listing 4-16 and Listing 4-17 both produce the same output when compiled and executed:
// 302.5

// Using the Short Variable Declaration Syntax
// The short variable declaration provides a shorthand for declaring variables, as shown in Listing 4-18. 
// Listing 4-18. Using the Short Variable Declaration Syntax in the 4_variables.go File in the 4_variables Folder

package main

import "fmt"

func main() {
	price := 275.00
	fmt.Println(price)
}


// The shorthand syntax specifies a name for the variable, a colon, an equal sign, and the initial value, as 
// illustrated by Figure 4-10. The var keyword is not used, and a data type cannot be specified.
// Figure 4-10. The short variable declaration syntax
// The code in Listing 4-18 produces the following output when the code is compiled and executed:
// 275

// Multiple variables can be defined with a single statement by creating comma-separated lists of names 
// and values, as shown in Listing 4-19.
// Listing 4-19. Defining Multiple Variables in the 4_variables.go File in the 4_variables Folder

package main

import "fmt"

func main() {
	price, tax, inStock := 275.00, 27.50, true
	fmt.Println("Total:", price + tax)
	fmt.Println("In stock:", inStock)
}

// No types are specified in the shorthand syntax, which means that variables of different types can 
// be created, relying on the compiler to infer types from the values assigned to each variable. 
// The code in  Listing 4-19 produces the following output when compiled and executed:
// Total: 302.5
// In stock: true

// The short variable declaration syntax can be used only within functions, such as the main function in 
// Listing 4-19. Go functions are described in detail later.


// Using the Short Variable Syntax to Redefine Variables
// Go doesn’t usually allow variables to be redefined but makes a limited exception when the short syntax is 
// used. To demonstrate the default behavior, Listing 4-20 uses the var keyword to define a variable that has the 
// same name as one that already exists within the same function.

// Listing 4-20. Redefining a Variable in the 4_variables.go File in the 4_variables Folder
// package main

// import "fmt"

// func main() {
// 	price, tax, inStock := 275.00, 27.50, true
// 	fmt.Println("Total:", price + tax)
// 	fmt.Println("In stock:", inStock)

// 	var price2, tax = 200.00, 25.00
// 	fmt.Println("Total 2:", price2 + tax)
// }

// The first new statement uses the var keyword to define variables named price2 and tax. There is 
// already a variable named tax in the main function, which causes the following error when the code is 
// compiled:
// .\4_variables.go:10:17: tax redeclared in this block

// However, redefining a variable is allowed if the short syntax is used, as shown in Listing 4-21, as 
// long as at least one of the other variables being defined doesn’t already exist and the type of the variable 
// doesn’t change.
// Listing 4-21. Using the Short Syntax in the 4_variables.go File in the 4_variables Folder

// package main

// import "fmt"

// func main() {
// 	price, tax, inStock := 275.00, 27.50, true
// 	fmt.Println("Total:", price + tax)
// 	fmt.Println("In stock:", inStock)
// 	price2, tax := 200.00, 25.00
// 	fmt.Println("Total 2:", price2 + tax)
// }

// Compile and execute the project, and you will see the following output:
// Total: 302.5
// In stock: true
// Total 2: 225

// ## Using the Blank Identifier
// It is illegal in Go to define a variable and not use it, as shown in Listing 4-22. 
// Listing 4-22. Defining Unused Variables in the 4_variables.go File in the 4_variables Folder

// package main

// import "fmt"

// func main() {
// 	price, tax, inStock, discount := 275.00, 27.50, true, true
// 	var salesPerson = "Alice"
// 	fmt.Println("Total:", price + tax)
// 	fmt.Println("In stock:", inStock)
// }

// The listing defines variables named discount and salesperson, neither of which is used in the rest of 
// the code. When the code is compiled, the following error is reported:
// .\4_variables.go:6:26: discount declared but not used
// .\4_variables.go:7:9: salesPerson declared but not used
// One way to resolve this problem is to remove the unused variables, but this isn’t always possible. For 
// these situations, Go provides the blank identifier, which is used to denote a value that won’t be used, as 
// shown in Listing 4-23.
// Listing 4-23. Using the Blank Identifier in the 4_variables.go File in the 4_variables Folder

// package main
// import "fmt"
// func main() {
// 	price, tax, inStock, _ := 275.00, 27.50, true, true
// 	var _ = "Alice"
// 	fmt.Println("Total:", price + tax)
// 	fmt.Println("In stock:", inStock)
// }

// The blank identifier is the underscore (the _ character), and it can be used wherever using a name 
// would create a variable that would not subsequently be used. The code in Listing 4-23 produces the 
// following output when compiled and executed:
// Total: 302.5
// In stock: true
// This is another feature that appears unusual, but it is important when using functions in Go. As I explain 
// in Chapter 8, Go functions can return multiple results, and the blank identifier is useful when you need some 
// of those result values but not others.