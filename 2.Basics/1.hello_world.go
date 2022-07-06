// A simple hello world go program
package main

import "fmt"

func main() {
 fmt.Println("Hello, world!")
}

// ## Explanation
// package main <------ Package Declaration

// import "fmt" <-------- Import Statement

// func main() {
//  fmt.Println("Hello, world!")   <------- Function
// }

// ## Understanding the Package Declaration
// The first statement is the package declaration. Packages are used to group related features, and every code 
// file has to declare the package to which its contents belong. The package declaration uses the package
// keyword, followed by the name of the package, as shown in Figure 3-2. The statement in this file specifies a 
// package named main. 

// package (keyword) main(package name)


// ## Understanding the Import Statement
// The next statement is the import statement, which is used to declare dependencies on other packages. The 
// import keyword is followed by the name of the package, which is enclosed in double quotes, as shown in above
// The import statement in the program specifies a package named fmt, which is the built-in Go 
// package for reading and writing formatted strings 

// import (keyword) "fmt" (Package name)

//  ## Understanding the Function
// The remaining statements in the 1.hello_world.go file define a function named main. 
// I describe functions in detail, but the main function is special. 
// When you define a function named main in a package named main, you create an entry point, which is where execution begins in a command-line application. 
// This is the main struvture for a main fucntion

// func (keyword) main (name) () (Parameters) { 
	// fmt.Println("Hello, World!") <----  Code Block
// }

// The basic structure of Go functions is similar to other languages. 
// The func keyword denotes a function  and is followed by the function’s name, which is main in this example.
// The function above defines no parameters, which is denoted by the empty parentheses and 
// produces no result. I describe more complex functions in later examples, but this simple function is enough 
// to get started with.
// The function’s code block contains the statements that will be executed when the function is invoked. 
// Since the main function is the entry point, the function will be invoked automatically when the compiled 
// output from the project is executed.


// ## Understanding the Code Statement
// The main function contains a single code statement. When you declare a dependency on a package with an 
// import statement, the result is a package reference that provides access to the package features. By default, 
// the package reference is assigned the name of the package so that the features provided by the fmt package, 
// for example, are accessed through a fmt package reference, as shown in understanding the import statement above.

// ### Accessing package features
// fmt (package reference) . (period) Println (feature name) ("Hello , World!") (Function Argument)

// This statement invokes a function named Println provided by the fmt package. This function writes 
// a string to the standard out, which means it will be displayed on the console when the project is built and 
// executed in the next section.

// To access the function, the package name is used, followed by a period and then the function: 
// fmt.Println. This function is passed one argument, which is the string that will be written.


// ## USING SEMICOLONS IN GO CODE
// Go has an unusual approach to semicolons: they are required to terminate code statements, but they 
// are not required in source code files. Instead, the Go build tools figure out where the semicolons need to 
// go as they process files, acting as though they had been added by the developer. 
// The result is that semicolons can be used in Go source code files but are not required and are 
// conventionally omitted.
// Some oddities arise if you don’t follow the expected Go code style. For example, you will receive 
// compiler errors if you attempt to put the opening brace for a function or for loop on the next line, 
// like this:

// package main
// import "fmt"
// func main()
// {
//  fmt.Println("Hello, Go")
// }

// The errors report an unexpected semicolon and a missing function body. This is because the Go tools 
// have automatically inserted a semicolon like this:

// package main
// import "fmt"
// func main();
// {
//  fmt.Println("Hello, Go")
// }

// The error messages make more sense when you understand why they arise, although it can be hard to 
// adjust to the expected code format if this is your preferred brace placement.
// I have tried to follow the no-semicolon convention throughout this book, but I have been writing code in 
// languages that require semicolons for decades, so you may find the occasional example where I have 
// added semicolons purely by habit. The go fmt command, which I describe in the “Formatting Go Code”
// section, will remove semicolons and adjust other formatting issues


// ## Compiling and Running Source Code
// The go build command compiles Go source code and produces an executable. Run the command shown in 
// below in the tools folder to compile the code. 

// Using the Compiler
// go build 1.hello_world.go

// The compiler processes the statements in the main.go file and generates an executable file, which is 
// named main.exe on Windows and main on other platforms. (The compiler will start creating files with more 
// useful names once I introduce modules in the “Defining a Module” section.)
// Run the command shown in above in the tools folder to run the executable.

// Running the Compiled Executable
// run ./1.hello_world

// The project’s entry point—the function named main in the package also named main—is executed and 
// produces the following output:
// Hello, World!


//  ## CONFIGURING THE GO COMPILER
// The behavior of the Go compiler can be configured using additional arguments, although the default 
// settings are sufficient for most projects. The two most useful are -a, which forces a complete rebuild 
// even for files that have not changed, and -o, which specifies the name of the compiled output file. 
// Use the go help build command to see the full list of options available. By default, the compiler 
// generates an executable file, but there are different outputs available—see https://golang.org/
// cmd/go/#hdr-Build_modes for details.

// Cleaning Up
// To remove the output from the compilation process, run the command shown in Listing 3-4 in the tools
// folder. 

// go clean 1.hello_world.go

// The compiled executable created in the previous section is removed, leaving only the source code 
// file behind.


//  ## Using the Go Run Command
// Most regular development is run using the go run command. Run the command shown in Listing 3-5 in the 
// tools folder. 

// go run 1.hello_world.go

// The file is compiled and executed in a single step, without creating the executable file in the tools
// folder. An executable file is created but in a temporary folder, from which it is then run. (It is this series of 
// temporary locations that caused the Windows firewall to seek permission every time the go run command 
// was used. Each time the command was run, an executable was created in a new temporary 
// folder, which appeared to be a completely new file to the firewall.)
// The command in Listing 3-5 produces the following output:
// Hello, Word!


//  ## Defining a Module
// The previous section demonstrated that you can get started just by creating a code file, but a more common 
// approach is to create a Go module, which is the conventional first step when starting a new project. Creating 
// a Go module allows a project to easily consume third-party packages and can simplify the build process. Run 
// the command shown above a tools folder. 

// Creating a Module can be done using the command below
// go mod init tools

// This command adds a file named go.mod to a tools folder. The reason that most projects start with 
// the go mod init command is that it simplifies the build process. Instead of specifying a particular code file, 
// the project can be built and executed using a period, indicating the project in the current directory. 
// Run the command shown in Listing above in the tools folder to compile and execute the code it contains 
// without specifying the name of a code file.


//  ## Compiling and Executing a Project
//  go run .

// The go.mod file has other uses—as later chapters demonstrate—but I start all the examples in the rest of 
//  the book with the go mod init command to simplify the build process.

// ## Debugging Go Code
// The standard debugger for Go applications is called Delve. 
// It is a third-party tool, but it is well-supported  and recommended by the Go development team. Delve supports Windows, macOS, Linux, and FreeBSD.  
// To install the Delve package, open a new command prompt and run the command shown below. 
// go install github.com/go-delve/delve/cmd/dlv@latest

// The go install command downloads and installs a package and is used to install tools such as 
// debuggers. A similar command—go get—performs a similar task for packages that provide code features 
// that are to be included in an application, as demonstrated in Chapter 12. 
// To make sure the debugger is installed, run the command shown in Listing 3-9.

// Running the Debugger
// dlv version

// If you receive an error that the dlv command cannot be found, then try specifying the path directly. By 
// default, the dlv command will be installed in the ~/go/bin folder (although this can be overridden by setting 
// the GOPATH environment variable), as shown in above


// ## Running the Debugger with a Path
// ~/go/bin/dlv

// If the package has been installed correctly, you will see the output similar to the following, although you 
// may see a different version number and build ID:
// Delve Debugger
// Version: 1.7.1
// Build: $Id: 3bde2354aafb5a4043fd59838842c4cd4a8b6f0b $


// ## DEBUGGING WITH THE PRINTLN FUNCTION
// I like debuggers like Delve, but I use them only for problems that I can’t figure out using my go-to 
// debugging technique: the Println function. I use Println because it is quick, simple, and reliable 
// and because most bugs (at least in my code) are caused because a function didn’t receive the value I
// expected or because a particular statement isn’t being executed when I expect. These simple issues are 
// easily diagnosed with a message writing to the console.
// If the output from my Println messages doesn’t help, then I fire up the debugger, set a breakpoint, 
// and step through my code. Even then, once I get a sense of the cause of a problem, I tend to go back to 
// Println statements to confirm my theory.
// Many developers are reluctant to admit they find debuggers awkward or confusing and end up secretly 
// using Println anyway. Debuggers are confusing, and there is no shame in using all the tools at your 
// disposal. The Println function and the debugger are complementary tools, and what’s important is 
// that bugs get fixed, regardless of how that is done.
// Preparing for Debugging
// The main.go file doesn’t contain enough code to debug. Add the statements shown in Listing 3-11 to create a 
// loop that will print out a series of numeric values.
// Adding a Loop in the main.go File in a tools Folder

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, World!")
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }


// I describe the for syntax in Chapter 6, but for this chapter, I just need some code statements to 
// demonstrate how the debugger works. Compile and execute the code with the go run . command; you will 
// receive the following output:
// Hello, Go
// 0
// 1
// 2
// 3
// 4


// Using the Debugger
// To start the debugger, run the command shown below in a tools folder ( the folder cantaing your project).

// dlv debug main.go
// This command starts the text-based debugging client, which can be confusing at first but is extremely 
// powerful once you get used to the way it works. The first step is to create a breakpoint, which is done by 
// specifying a location in the code, as shown in Listing 3-13.

// Listing 3-13. Creating a Breakpoint
// break bp1 main.main:3

// The break command creates a breakpoint. The arguments specify a name for the breakpoint and 
// a location. Locations can be specified in different ways, but the location used in Listing 3-13 specifies a 
// package, a function within that package, and a line within that function, as illustrated by Figure 3-6.

// Specifying a breakpoint location
// The name for the breakpoint is bp1, and the location specifies the third line in the main function in the 
// main package. The debugger displays the following confirmation message:
// Breakpoint 1 set at 0x697716 for main.main() c:/tools/main.go:8
// Next, I am going to create a condition for the breakpoint so that execution will be halted only when a 
// specified expression evaluates to true. Enter the command shown in Listing 3-14 into the debugger and 
// press Return. 

// Listing 3-14. Specifying a Breakpoint Condition in the Debugger
// condition bp1 i == 2
// The arguments for the condition command specify a breakpoint and an expression. This command 
// tells the debugger that the breakpoint named bp1 should halt execution only when the expression i == 2 is 
// true. To start execution, enter the command shown in Listing 3-15 and press Return.

// Listing 3-15. Starting Execution in the Debugger
// continue

// The debugger starts to execute the code, producing the following output:
// Hello, World!
// 0
// 1

// Execution is halted when the condition specified in Listing 3-15 is true, and the debugger displays the 
// code and the point at which execution stops, which I have marked in bold:
// > [bp1] main.main() c:/tools/main.go:8 (hits goroutine(1):1 total:1) (PC: 0x207716)
//  3: import "fmt"
//  4:
//  5: func main() {
//  6: fmt.Println("Hello, World")
//  7: for i := 0; i < 5; i++ {
// => 8: fmt.Println(i)
//  9: }
//  10: }

// The debugger provides a full set of commands for inspecting and altering the state of the application, 
// the most useful of which are shown in Table 3-2. (See https://github.com/go-delve/delve for the full set 
// of commands supported by the debugger.)
// Table 3-2. Useful Debugger State Commands
// Command 														Description
// print <expr> 												This command evaluates an expression and displays the result. It can 
															// be used to display a value (print i) or perform a more complex test 
															// (print i > 0).
							
// set <variable> = <value> 									This command changes the value of the specified variable.

// locals 														This command prints the value of all local variables.
													
// whatis <expr> 												This command prints the type of the specified expression such as 
															// whatis i.

// Run the command shown in Listing 3-16 to display the current value of the variable named i.

// Listing 3-16. Printing a Value in the Debugger
// print i

// The debugger displays the response 2, which is the current value of the variable and matches the 
// condition I specified for the breakpoint in Listing 3-16. The debugger provides a full set of commands for 
// controlling execution, the most useful of which are shown in Table 3-3. 

// Table 3-3. Useful Debugger Commands for Controlling Execution
// Command 		Description
// continue 	This command resumes execution of the application.
// next 		This command moves to the next statement.
// step 		This command steps into the current statement.
// stepout 		This command steps out of the current statement.
// restart 		This command restarts the process. Use the continue command to begin execution.
// exit 		This command exits the debugger.


// Enter the continue command to resume execution, which will produce the following output:
// 2
// 3
// 4
// Process 3160 has exited with status 0

// The condition I specified for the breakpoint is no longer met, so the program runs until it terminates. 
// Use the exit command to quit the debugger and return to the command prompt.



//  ## Fixing Common Problems in Go Code
// The go vet command identifies statements likely to be mistakes. Unlike a linter, which will often focus on 
// style issues, the go vet command finds code that compiles but that probably won’t do what the developer 
// intended. 
// I like the go vet command because it spots errors that other tools miss, although the analyzers don’t 
// spot every mistake and will sometimes highlight code that isn’t a problem. In Listing 3-26, I have added a 
// statement to the 1.hello_world.go file that deliberately introduces a mistake into the code.
// Listing 3-26. Adding a Statement in the 1.hello_world.go File in the tools Folder
// package main

// import "fmt"

// func main() {
// 	PrintHello()
// 	for i := 0; i < 5; i++ {
// 		i = i
// 		PrintNumber(i)
// 	}
// }

// func PrintHello() {
//  fmt.Println("Hello, Go")
// }


// func PrintNumber(number int) {
//  fmt.Println(number)
// }

// The new statement assigns the variable i to itself, which is allowed by the Go compiler but is likely to be 
// a mistake. To analyze the code, use the command prompt to run the command shown in Listing 3-27 in the 
// tools folder.

// Listing 3-27. Analyzing Code
// go vet 1.hello_world.go

// The go vet command will inspect the statements in the main.go file and produce the following 
// warning:

// # _/C_/tools
// .\1.hello_world.go:8:9: self-assignment of i to i

// The warnings produced by the go vet command specify the location in the code where a problem has 
// been detected and provide a description of the issue.
// The go vet command applies multiple analyzers to code, and you can see the list of analyzers at 
// https://golang.org/cmd/vet. You can select individual analyzers to enable or disable, but it can be difficult 
// to know which analyzer has generated a specific message. To figure out which analyzer is responsible for a 
// warning, run the command shown in Listing 3-28 in the tools folder.

// Listing 3-28. Identifying an Analyzer
// go vet -json main.go
// The json argument generates output in the JSON format, which groups warnings by analyzer, like this:

// # _/C_/tools {
// 	"_/C_/tools": {
// 		"assign": [
// 			{
// 				"posn": "C:\\tools\\main.go:8:9",
// 				"message": "self-assignment of i to i"
// 			}
// 		]
// 	}
// }

// Using this command reveals that the analyzer named assign is responsible for the warning generated 
// for the main.go file. Once the name is known, the analyzer can be enabled or disabled, as shown in 

// Listing 3-29.
// go vet -assign=false
// go vet -assign

// The first command in Listing 3-29 runs all the analyzers except assign, which is the analyzer that 
// produced the warning for the self-assignment statement. The second command runs only the assign
// analyzer.