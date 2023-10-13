// Create 2 files named a.txt b.txt in your current project directory.
// Create your main.go file.
// Inside main.go file, Write a function that takes the fileName as an input.
// Inside the function, you have to call os.Remove() which would remove a file from the current directory.
// Now call this function from your main function & pass the above created file names to this function.

package main

import (
	"fmt"
	"os"
)

func remFile(fileName string) error { //remove file func
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Error removing file:", err)
	}
	return err
}

func main() {
	A := "a.txt"
	B := "b.txt"

	// f, err := os.Create(A)
	// if err != nil {
	// 	fmt.Println("Error creating file A:", err)
	// 	return
	// }

	// f2, err := os.Create(B)
	// if err != nil {
	// 	fmt.Println("Error creating file B:", err)
	// 	return
	// }
	err := remFile(A) //rem file A
	if err != nil {
		fmt.Println("Error removing file A:", err)
	}

	err = remFile(B)
	if err != nil {
		fmt.Println("Error removing file B:", err)
	}

	// fmt.Println(f, f2)
}
