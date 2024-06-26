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

```
func (d deck, handSize int) return_type {

}
```

You can return more than one item from a function in go using the following syntax

```
func (d deck, handSize int) (return_type_one, return_type_two) {
    return d[:handSize], d[handSize:]
}
```

You can capture multiple return values from a function by doing something like the following:

```
first, second := twoReturnTypes()
```

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

When you define a receiver function as above, you give any variable of type deck access to the receiver
function, in this case print. The variable d is the variable of type deck and it acts as if it is being
passed into the print function that you can then reference from within the function. This is similar to
how we use this or self in other languages.

## Byte Slices
Byte slices is a common data structure in Golang that is basically the equivalent of a C string where each
invidiual character in the string is represented by an ASCII Code.

## Type Conversion
To do type conversion in Golang, we can do something like the following:
`[]byte("Hi there")`
where byte is the type that we want and "Hi there" is what we want to convert.

## Err type
If nothing goes wrong, the value stored in err will be of type 'nil'. If something does go wrong,
the type will be some form of error based on the type of function called. It is common in Go
to do something like the following:

```
data, err := ioutil.ReadFile(filename)
if err != nil {
    // error handle
}

// continue on
```

Using the standard library OS package, you are able to exit a program completely by doing:
```
os.Exit(1)
```
for an error, or 0 for success (exit codes).

## Random Number Generation
The math/rand package contains many different functions for pseudo-random number generation. One of these
is: 
`rand.intN(5)`
Which takes in a non inclusive max value for a random integer. This pseudo-random number generator is seeded
and in order to change it, we can do the following:
```
source := rand.NewSource()
r := rand.New(source)
newPos := r.Intn(len(d))
```
We can use current time in order to seed our random number generator.

## Testing in Go
In order to be able to run tests, we need to first run
`go mod init`
to create a new go.mod file in the current directory. This file is used to specify the specific versions of the
dependencies that you are using in your project. Go tests are *not* like other testing frameworks. Go tests are built
right into the language and are very different. To create a new test in Go, we just create a new file matching .*_test.go.
Go will then recognize this as a test file and you can then run all tests using
`go test`
