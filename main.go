package main

import (
	"os"
	"strings"

	"fmt"
)


func main()  {
	entre:= os.Args[1]
	
	txt,err := os.ReadFile(entre)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	} else{
		
		res:= strings.Fields(string(txt))
		
		fmt.Print(res)
	}
		}
	//os.WriteFile("test.txt",[]byte(txt),0644)//
	
		
	
	
