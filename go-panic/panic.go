package main

import "fmt"

func entry(lang *string, aname *string) {

	defer func() {
		fmt.Println(recover())
	}()

	if lang == nil {
		panic("Error: Language cannot be nil")
	} else if aname == nil {
		panic("Error: Author name cannot be nil")
	} else {
		fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname)
	}
}

// Main function
func main() {

	A_lang := "GO Language"
	A_name := "Aswath"

	entry(&A_lang, nil)
	entry(&A_lang, &A_name)
}
