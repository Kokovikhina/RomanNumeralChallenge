package main

import (
    "fmt"
    "os"
    "romannumeralchallenge/roman"
)

func main() {
    //var then variable name then variable type
    var romanNumber roman.RomanDecimal
    if(len(os.Args) < 2){
        fmt.Println("Enter Roman Number: ")    
        //Taking input from user
        fmt.Scanln(&romanNumber.Roman)
    }else{  
        //NewRoman func returns address
        romanNumber = *roman.New(os.Args[1])   
    }
    fmt.Println(romanNumber.Convert())
    fmt.Println(romanNumber.Roman)

}
