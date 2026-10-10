package folderfunc

import (
	"fmt"
	"strconv"
)

func Bin(phrase []string) []string {
	var  res []string 
	for i:=0; i<len(phrase); i++{
		if phrase[i]== "(bin)" && i>0{
			data,err := strconv.ParseInt(res[i-1],2,64)
			if err!= nil{
				fmt.Println("Il y a un problème avec bin")
			}
			res[len(res)-1] = strconv.Itoa(int(data))
			i++
		}else{
			res = append(res, phrase[i])
		}
	}
	return res

}


