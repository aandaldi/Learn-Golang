# Array, Slice, and Map

## Array
- an array is a collection of elements with the same data type.
 array have limit, we can set when declare the array
- default value from an array depending on the data-type of varible
- declare an array like:
        
    -       var <varname> [<limit>]<datatype>
            <varname>[n] = <val1>
            <varname>[n] = <val2>
    -       var <varname> = [<limit>]<datatype>{<val1>, <val2}
    -       var <varname> = [<limit>]<datatype>{
                        <val1>,
                        <val2>,
            }

- you can also declare an array without limit element, just replace the limit wih `...` like:
    
    -     var <arrayname> = [...]<datatype>{<val1>, <val2>, <val3>}

- for create array multi medimensi, create like:
        
        var <arrayname> = [<limit1>][<limit2>]<datatype>{{<val1>,<val2>}, {<val1>, <val2>}}
    - first dimension is total array of global dimention array,
    - second dimention is element total of sub dimention
- we can use loop for get element one by one from an array like:
    ~~~
    var <arrayname> = [...]<datatype>{<val1>, <val2>, <val3>}
    for <varCounter> := 0; <varCounter> < len(<arrayname>); <increament>{
        fmt.Prinln(<arrayname[varCounter]>)
    }
    ~~~
* declare array with keyword `for`-`range`
    ~~~
    var <elements> = [...]<datatype>{<val1>, <val2>, <val3>}
    for <varCounter>, <element> := range <elements>{
        fmt.Prinln(<varCounter>, <element>)
    }
    ~~~
    
    if we will not to used <var_counter>, we can replace with keyword `_` for placed the value in trash

* declare element array with keyword `make` like:
    ~~~
    var <arrayname> = make([]<datatype>, <limit>)
    <arrayname>[<n>] = <valn>
    <arrayname>[<n>] = <valn>
    <arrayname>[<n>] = <valn>
    ~~~