# Function, Pointer and Struct

## [Function](https://github.com/aandaldi/Learn-Golang/tree/aan/Learn-step-by-step/function-pointer-struct/function.go)
* Function is certain set of code blocks that is wrapped with spesific name. The proper functions will make code modular and dry (in short, don't repeat yourself), no need to write a lot of code
* to declare some funtions like: 
    ~~~
    func <funcName>(<param-if-need>){
        //code
        //<use param if it is defined>
    }
    ~~~
* call function like this, `<functionName>(<param-if-it-is-defined>)` or `<functionName>()` if not necessary param
* if we want to use the value of the function, we can use the return value. to declare funtion with return value we can use template like this :
    ~~~
    func <funtionName>(<param-if-need>)<data-type-return-value>{
        var <varName> = <value>
        // code
        return <varName>
    }
    ~~~
    we can return multiple variable like  `return <var1>, <var2>, ....`
* we can also  declare function with predefine return value like:
    ~~~
    func <funcName>(<param>)(<ReturnVarName1> <data-Type-var1>, <ReturnVarName2> <data-Type-var2> )
    //code
    <returnVarName1> := <value>
    <returnVarName2> := <value>

    return
    ~~~

* Go adopt variadic function concept or create function with similar functions infinite. this function create with  the template like this:
    ~~~
    func <funcName>(<parmVarName> ...<data-type>){
        //code
    }
    ~~~
    or you can use slice for save multiple param in a variable. 

    call like this `<funcName>(<val1>, <val2>, <val3> ...`

* Sprintf function is used to return or save in variable some value with string data type
* float64 function is used to convert/casting value to float64
* we can also  combine reguler function paramter with varadic function parameters

    ### Closure Function
    * defition of closure function is a function that can be stored in a variable. when implementing that function, we can create function inside the function or we can create function with return a function.
    * closure is anonymous function
    * declare closure when we save as a variable like this:
        ~~~
        var <varName> = func(<paramVarName> <paramVarName-DataType>)<return-dataType>{
            //code
            return <returnVal>
        }
        ~~~
        if we wan't to call the function, we can write the variabelName with param like this: `resul := <varName>(paramVarName)`
    * Immediately-Invoked Function Expression (IIFE) is closure that will be executed when declaring it. its trademark is the presence of brackets at the end declare closure with the param value. declare like this:

        ~~~
        var <varName> = func(<paramVarName> <paramVarName-DataType>)<return-dataType>{
            //code
            return <returnVal>
        }(<paramValue>)
        ~~~

    * closure as return value like this:
        ~~~
        var <varName> = func(<paramVarName> <paramVarName-DataType>)(<return-dataType>, func() <funcData-Type>){
            //code
            return <returnVal>, func() <funcData-Type>{
                return <returnValFunc>
            }
        }
        ~~~
