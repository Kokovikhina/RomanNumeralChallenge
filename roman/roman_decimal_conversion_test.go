package roman

import (
    "fmt"
    "testing"
)

func TestRomanToDecimal (t *testing.T) {
    cases := []struct {
        in RomanDecimal; want int
    }{
        {RomanDecimal{Roman: "XXXIX"}, 39},
        {RomanDecimal{Roman: "CCXLVI"}, 246},
        {RomanDecimal{Roman: "DCCLXXXIX"},789}, 
        {RomanDecimal{Roman: "MMCDXXI"}, 2421},
        {RomanDecimal{Roman: ""}, 0},
        {RomanDecimal{Roman: " "}, 0},
        {RomanDecimal{Roman: " V "}, 5},
        {RomanDecimal{Roman: "U"}, 0},
        {RomanDecimal{Roman: "VU"}, 0},
        {RomanDecimal{Roman: " MM CDXXI "}, 2421},
        {RomanDecimal{Roman: "   dcc   L  xxxIX "}, 789},
    }
    
    for _, c := range cases {
        testname := fmt.Sprintf("%q, %d", c.in, c.want)
        t.Run(testname, func(t *testing.T){
            c.in.Convert()
            got := c.in.Decimal
            if got != c.want {
                t.Errorf("ToDecimal(%q) == %d, want %d", c.in, got, c.want)
            }
        })
    }

}
