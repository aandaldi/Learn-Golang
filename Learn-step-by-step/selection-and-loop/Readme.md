# Selection and Loop

## Selection
* in Go, selection of condition any 2 ways, there are `if else` and `switch`.

### A. [`if`, `if else`, and `else`](https://github.com/aandaldi/Learn-Golang/blob/aan/Learn-step-by-step/selection-and-loop/if_else.go)
* declare with operator and parenthesis
* the format like:
    
        if <aCondition>{
            //code
        } else if <bCondition> {
            //code
        } else {
            //code
        }
* you can do nested selection like this:

        if <aCondition>{
            if <bCondition>{
                //code
            }
        } else {
            //code
        }

    or 

        if <aCondition> && <bCondition>{   // (&&) for use "and", (||) for (or)
                //code
        } else {
            //code
        }

    [Example](https://github.com/aandaldi/Learn-Golang/blob/aan/Learn-step-by-step/selection-and-loop/if_else.go)
    <br/>
<br/>

### B. [`switch` - `case`](https://github.com/aandaldi/Learn-Golang/blob/aan/Learn-step-by-step/selection-and-loop/switch_case.go)
* Switch is selection which is focused on one variable, then cek the value, simple example: is value of x are 1,2,3,4 or etc?
* declare with case number and parenthesis
* the format like:
    
        <var> := <val>

        switch <var>{
        case <val1>:
            //code
        case <val2>:
            //code
        case <val3>:{
            //code
            //code
        }    
        case <val4>, <val5>, <val6>:
            //code
        default :
            //code    
        }
    default value will be run when not any case is true. default this like else in if-else selection.
* in go, switch just focuse on one variable and one case, if some case are true, then code in insede of the case will be run, then direct to break althought there not any `break` command

* keyword `fallthrough` will be direct `switch-case` to check another case

* we can write switch like `if-else` style, like this:
    
        <var> := <val>

        switch {
        case <var> == <val1>:
            //code
        case (<var> > <val3>) && (<var> < <val7>):
            //code
        case <val3>:{
            //code
            //code
        }    
        case <val4>, <val5>, <val6>:
            //code
        default :
            //code    
        }

    [Example](https://github.com/aandaldi/Learn-Golang/blob/aan/Learn-step-by-step/selection-and-loop/switch_case.go) <br/>
<br/>


## Loop
* in Go, selection of condition any 2 ways, there are `if else` and `switch`.

### A. [`for`](https://github.com/aandaldi/Learn-Golang/blob/aan/Learn-step-by-step/selection-and-loop/for_loop.go)
* declare with counter variable and parenthesis, the command like this:
        
        for <varCounter> := 0; <varCounter> < <limitLoop>; <increament>{
            //code
        }
    
* another way to declare `for` just with condition like `while` concept in another language
        
        for <condition>{
            //code
            <increment>
        }

* alse any way to declare `for` like `do while`:

        for {
            //argument
            <increment>
            if <condition>{
                break
            }
        }

* keyword `break` use to force stop the loop. keyword `continue` for direct to next loop

* you can create label for every loop command, so can use break or continue from or to label. eg:

        <label>:
        for <varCounter> := 0; <varCounter> < <limitLoop>; <increament>{
            if <condition>{
                break <label>
            }
            //code
        }

[Example](https://github.com/aandaldi/Learn-Golang/blob/aan/Learn-step-by-step/selection-and-loop/for_loop.go) <br/>
<br/>