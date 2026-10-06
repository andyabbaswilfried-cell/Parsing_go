package main

import ("os"

		"fmt"

)


func main()  {
	entre:= os.Args[1]
	txt,err := os.ReadFile(entre)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println(string(txt))
}