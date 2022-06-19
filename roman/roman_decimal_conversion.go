package roman

import (
    "strconv"
    "romannumeralchallenge/utils"
)

type Roman interface {
    New(string)
    Convert() string
}

type RomanDecimal struct {
    Roman utils.String
    Decimal int
}

//Create new object and return address instead of value
func New (number string) *RomanDecimal{
    return &RomanDecimal{
        Roman: utils.String(number),
        Decimal: 0,
    }
}

//Start with upper case because public(exported)
func (r *RomanDecimal) Convert() string {
    //Remove whitespaces at the begining and end and make everthing upper case
    r.Roman.TrimSpace().ToUpper().ToString()

    //Return 0 if no Roman number is given
    if (string(r.Roman) == ""){
        return "0"
    }
    romanDecimalMapping := map[string] int{
        "I": 1,
        "V": 5,
        "X": 10, 
        "L": 50, 
        "C": 100,
        "D": 500,
        "M": 1000,
    }   
    //Hold conversion result in int for any future use
    r.Decimal = 0 
    var right_val int  = romanDecimalMapping[string(r.Roman[len(r.Roman)-1])]
    for i := len(r.Roman)-1; i >= 0; i-- {
        if (string(r.Roman[i]) == " "){
            continue
        }
        var left_val int = romanDecimalMapping[string(r.Roman[i])]
        if(left_val==0){
            r.Decimal = 0
            return "0"
        }
        //Check for substraction
        if left_val < right_val {
            r.Decimal = r.Decimal - left_val
        } else {
            r.Decimal = r.Decimal + left_val
        }   
        right_val = left_val
    
    }   
    //Convert integer to string before returning result
    return strconv.Itoa(r.Decimal)
}
