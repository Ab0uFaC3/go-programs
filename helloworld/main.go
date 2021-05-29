/*
	Package is like a Project or a Workspace
	It is a collection of common source code files. While working on discrete application, one package is used
	The only requirement for every file inside of a package is that the very first line of each file must declare the package that it belongs to.

	-* 2 types of packages *- : To diff, name of package decides ("main" is only executable package)
		    1. 	Executable - when compiled, actual runnable file or an executable file is obtained
		    2. 	Reusable - Like code dependencies or libraries. Instead, we put in a lot of reusable logic or helper functions or stuff
		           	       that will just help us reuse some code on future projects in the future.
*/
package main

/*
	The import statement is used to give our package access to some code that is written,
	written inside of another package.

	fmt - format, implements formatted I/O information specifically on the terminal for debugging
*/
import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
