# Learn Variable and Data Type
## Variable
* For create variable with var and datatype, any 2 ways: 

    1. `var <nama-variabel> <data-type>`
    2. `var <nama-variabel> <data-type> = <value>`
    <br/>

* For declare variable without datatype, use with `=` with `var` or `:=` without `var`: 
    1. `var <variable-name> = <value>`
    2. `<variable-name> := <value>`    declare with `:=` can only be place on block function (type inference technique)
    <br/>

* For declare multi variable simultaneously, the way is write the variables with comma (`,`) separate. eg:
    1. `var <var1>, <var2>, <var3> <data-type>` 
    1. `<var1>, <var2>, <var3> = <value1>,  <value1>,  <value1>`
    1. `<var1>, <var2>, <var3> := <value1>,  <value1>,  <value1>`
    <br/>

* For type inference, can declare multivariable value
    <br/>

* Variable underscore (`_`) is reserved variable, which can be used to store unused values/variable. eg:
    1. `_ = <value>`
    1. `var <var1>, _ = <val1>, <val2> `
    1. `<var1>, _ := <val1>, <val2> `
    <br/>

* Declare variabel with `new` keyword.
  - `new` keyword used to create a variabel pointer with spesific data type. Default value will match the data type.
  - Declare with `<var> := new(<data-type>)` eg:
    ~~~
      name := new(string)
      fmt.Println(name)   // 0x20818a220 --> Because name save the pointer string 
      fmt.Println(*name)  // ""          --> dereference, to get value of pointer
    ~~~
  <br/>

* Declare with `make` keyword can only be used to create several types of variable, namely:

  - channel,
  - slice,
  - map
  <br/>


## Data Type

### A. Integer data type
- There are several types of non-decimal or non floating point numeric data types in Go. In general, there are 2 types of data in this category that need to be known.
  
  1. `uint`, for positif number.
  1. `int`, for positif or negatif number.

- if use `fmt.Printf()` function for formating data type, its only run when use template `%d` for integer formats

- 


### A. String
* Using `fmt.Printf()` for print variable/argument with string type and using with `%s` symbol or just print string argument. eg :
  ~~~ 
    var firstName = "aan"
    var *lastName string
    lastName = "aldi"
    
    fmt.Printf("Welcome \n")
    fmt.Printf("Halo %s %s \n", firstName, lastName) // Halo aan aldi 


  ~~~
