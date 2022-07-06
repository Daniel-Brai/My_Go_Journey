//  Using a simple project to explain data types
// as shown below

// package main
// import (
// 	"fmt"
// 	"math/rand"
// )
// func main() {
// 	fmt.Println(rand.Int())
// }

// Use the command prompt to run the command shown in Listing 4-3 in the datatypes folder.
// Running the Example Project
// go run .
// The code in the 1_datatypes.go file will be compiled and executed, producing the following output:
// 5577006791947779410
// The output from the code will always be the same value, even though it is produced by the random 
// number package, as I explain later

// A related feature provided by the fmt package is the ability to compose strings by combining static 
// content with data values, as shown in Listing 4-4.
// Listing 4-4. Composing a String in the 1_datatypes.go File in the datatypes Folder
package main

import (
 "fmt"
 "math/rand"
)

func main() {
	fmt.Println("Value:", rand.Int())
}

// The series of comma-separated values passed to the Println function are combined into a single string, 
// which is then written to the standard output. 
// To compile and execute the code, use the command prompt to  run the command shown in Listing 4-5 in this  folder.

// Listing 4-5. Running the Example Project
// go run .

// The code in the 1_datatypes.go file will be compiled and executed, producing the following output:
// Value: 5577006791947779410

// There are more useful ways to compose strings—which I describe in Part 2 —but this is simple and a 
// useful way for me to provide output in the examples

// Understanding the Basic Data Types
// Go provides a set of basic data types, which are described in Table 4-3. In the sections that follow, I describe 
// these types and explain how they are used. These types are the foundation of Go development, and many of 
// the characteristics of these types will be familiar from other languages. 
// Table 4-3. The Go Basic Data Types
// Name 										Description
// int 											This type represents a whole number, which can be positive or negative. The 
// 												int type size is platform-dependent and will be either 32 or 64 bits. There are 
// 												also integer types that have a specific size, such as int8, int16, int32, and 
// 												int64, but the int type should be used unless you need a specific size.

// uint 											This type represents a positive whole number. The uint type size is platform dependent and will be either 32 or 64 bits. There are also unsigned integer 
// 												types that have a specific size, such as uint8, uint16, uint32, and uint64, but 
// 												the uint type should be used unless you need a specific size.

// byte 											This type is an alias for uint8 and is typically used to represent a byte of data.

// float32, float64 								These types represent numbers with a fraction. These types allocate 32 or 64 
// 												bits to store the value.

// complex64, complex128 							These types represent numbers that have real and imaginary components. 
// 												These types allocate 64 or 128 bits to store the value.

// bool 											This type represents a Boolean truth with the values true and false.

// string 											This type represents a sequence of characters.

// rune 											This type represents a single Unicode code point. Unicode is complicated, 
// 												but—loosely—this is the representation of a single character. The rune type is 
// 												an alias for int32.


// ## Understanding Literal Values
// Go values can be expressed literally, where the value is defined directly in the source code file. Common 
// uses for literal values include operands in expressions and arguments to functions, as shown in Listing 4-6. 

// Notice that I have commented out the math/rand package from the import statement in Listing 4-6. 
// It is an error in Go to import a package that is not used.
// Listing 4-6. Using Literal Values in the 1_datatypes.go File in the datatypes Folder

// package main

// import (
//  "fmt"
//  //"math/rand"
// )

// func main() {
// 	fmt.Println("Hello, Go")
// 	fmt.Println(20 + 20)
// 	fmt.Println(20 + 30)
// }

// The first statement in the main function uses a string literal, which is denoted by double quotes, as an 
// argument to the fmt.Println function. The other statements use literal int values in expressions whose 
// results are used as the argument to the fmt.Println function. Compile and execute the code, and you will 
// see the following output:
// Hello, Go
// 40
// 50

// You don’t have to specify a type when using a literal value because the compiler will infer the type based 
// on the way the value is expressed. For quick reference, Table 4-4 gives examples of literal values for the 
// basic types.

// Table 4-4. Literal Value Examples
// Type 							Examples
// int 							20, -20. Values can also be expressed in hex (0x14), octal (0o24), and binary notation 
// 								(0b0010100).

// uint 							There are no uint literals. All literal whole numbers are treated as int values.

// byte 							There are no byte literals. Bytes are typically expressed as integer literals (such as 101) or run 
// 								literals ('e') since the byte type is an alias for the uint8 type.

// float64 						20.2, -20.2, 1.2e10, 1.2e-10. Values can also be expressed in hex notation (0x2p10), although 
// 								the exponent is expressed in decimal digits.

// bool 							true, false.

// string 							"Hello". Character sequences escaped with a backslash are interpreted if the value is enclosed 
// 								in double quotes ("Hello\n"). Escape sequences are not interpreted if the value is enclosed in 
// 								backquotes (`Hello\n`).

// rune 							'A', '\n', '\u00A5', '¥'. Characters, glyphs, and escape sequences are enclosed in single 
// 								quotes (the ' character)