package errors

import "fmt"

func PanicIfError(err error){
	if err != nil{
		panic(err)
	}
}


func PrintIfError(err error,message string){
	if err != nil{
		fmt.Println(message,err)
	}
}

func PanicIfErrors(errs ...error){
	for _,err := range errs{
		if err != nil{
			panic(err)
		}
	}
}