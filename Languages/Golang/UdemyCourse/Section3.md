# Section 3 - Deeper Into Go

## Variables
There are multiple different ways to define variables in Go with different use cases. Go
is a statically typed language, but it does have some other properties that allow it to
act more like a dynamically typed language.

1. `var variable_name string = "value"`   => statically typed. Value is not required to be
                                           defined here
2. `variable_name := "value"`             => inferred type. Once the type is specified, it
                                           cannot be changed later on

The ":=" is called the walrus operator and is used for inferring type. You only use this at
variable declaration, not when reassigning a variable.

## Basic Go Types
- bool
- string
- int
- float64

## Functions and Return Types
Basic syntax of a function:

func () return_type {

}

## Slices and For Loops
### Array
Fixed size storage of data.

`var my_arr[10] int`

### Slice
Is a type of an array that can grow or shrink dynamically. All elements of a slice must
be of the same type.

`my_slice := []string{}`

To append to a slice, you can do the following:

`new_slice = append(my_slice, "new element")`

The append function does not overwrite the slice provided, but rather returns a new slice
that we can then assign to a variable.

### For loops
Range based for loop
```
for i, card := range cards {
    fmt.Println(i, card)
}
```

## OO Approach vs Go Approach
Go is not an object oriented language. There are no ideas of classes in Go. Instead, we use
the basic Go types and extend the functionality of them by defining new custom types. For example,
we can do the following:

`type deck []string`

We can then create functions that have a receiver- a function that takes in a custom type- to add
functionality to our type. 

## Receiver Functions
Below is how you can define a receiver function for a custom type
```
func (d deck) print() {
	
}
```