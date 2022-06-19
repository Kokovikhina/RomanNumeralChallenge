# RomanNumeralChallenge

## Table of Content:

   - [About roman numeral challenge](#about-roman-numeral-challenge)
   - [Instructions - how to run and test ](#instructions)
   - [Reference links](#reference-links)


## About roman numeral challenge: 

We'd like you to write some code to parse and output roman numerals and share it with us as a GitHub repo.
If we wrapped your code into a console application called `roman`, we'd expect it to work like this:

```
$ roman ix
9
ix

$ roman XIIII
14
XIV

$ roman MCMXCIX
1999
MCMXCIX
```

Additional Notes

* [Roman numerals on Wikipedia](https://en.wikipedia.org/wiki/Roman_numerals)
* Do create an abstract data type to represent a roman number.
* Do write unit tests.
* Do include links to any references you used.
* Do apply Postel's Law
* Don't take much more than an hour.
* Don't use someone else's code to do the parsing (we want to see how _you_ do it!).

We'd prefer your implementation to be written in  Go; if you use a different language, please advise why you chose it.


## Instructions 

### To run program  
  ```
  $ go run .
  Enter Roman Number: x 
  10
  x
  ```

### To run program with roman number as argument  
  ```
  $ go run . x
  10
  x
  ```
### To run tests 
  ```
  $ cd roman
  $ go test -v

    === RUN   TestRomanToDecimal
    === RUN   TestRomanToDecimal/{"XXXIX"_'\x00'},_39
    === RUN   TestRomanToDecimal/{"CCXLVI"_'\x00'},_246
    === RUN   TestRomanToDecimal/{"DCCLXXXIX"_'\x00'},_789
    === RUN   TestRomanToDecimal/{"MMCDXXI"_'\x00'},_2421
    === RUN   TestRomanToDecimal/{""_'\x00'},_0
    === RUN   TestRomanToDecimal/{"_"_'\x00'},_0
    === RUN   TestRomanToDecimal/{"_V_"_'\x00'},_5
    === RUN   TestRomanToDecimal/{"U"_'\x00'},_0
    === RUN   TestRomanToDecimal/{"VU"_'\x00'},_0
    === RUN   TestRomanToDecimal/{"_MM_CDXXI_"_'\x00'},_2421
    === RUN   TestRomanToDecimal/{"___dcc___L__xxxIX_"_'\x00'},_789
    --- PASS: TestRomanToDecimal (0.00s)
        --- PASS: TestRomanToDecimal/{"XXXIX"_'\x00'},_39 (0.00s)
        --- PASS: TestRomanToDecimal/{"CCXLVI"_'\x00'},_246 (0.00s)
        --- PASS: TestRomanToDecimal/{"DCCLXXXIX"_'\x00'},_789 (0.00s)
        --- PASS: TestRomanToDecimal/{"MMCDXXI"_'\x00'},_2421 (0.00s)
        --- PASS: TestRomanToDecimal/{""_'\x00'},_0 (0.00s)
        --- PASS: TestRomanToDecimal/{"_"_'\x00'},_0 (0.00s)
        --- PASS: TestRomanToDecimal/{"_V_"_'\x00'},_5 (0.00s)
        --- PASS: TestRomanToDecimal/{"U"_'\x00'},_0 (0.00s)
        --- PASS: TestRomanToDecimal/{"VU"_'\x00'},_0 (0.00s)
        --- PASS: TestRomanToDecimal/{"_MM_CDXXI_"_'\x00'},_2421 (0.00s)
        --- PASS: TestRomanToDecimal/{"___dcc___L__xxxIX_"_'\x00'},_789 (0.00s)
    PASS
    ok      romannumeralchallenge/roman     0.291s
  ```


## Reference links

* https://go.dev/doc/code

* https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68

* https://www.geeksforgeeks.org/fmt-scanln-function-in-golang-with-examples/

* https://go101.org/article/string.html

* https://en.wikipedia.org/wiki/Roman_numerals

* https://adrianwit.medium.com/abstract-class-reinvented-with-go-4a7326525034

* https://stackoverflow.com/questions/18152419/chaining-functions-in-go
   
* https://stackoverflow.com/questions/10105935/how-to-convert-an-int-value-to-string-in-go
   
*  https://go.dev/doc/effective_go#embedding 
    
* https://gobyexample.com/testing 
    
* https://go.dev/blog/maps 
    
* https://www.golangprograms.com/example-of-abstraction-using-interfaces-in-golang.html

