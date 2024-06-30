# Section 4 - Organizing Data with Structs
A struct in Go is a collection of properties that are related together, very similar to other
programming languages. In Go, structs offer similar functionality to classes in other languages.

## Defining Structs
In order to define a struct, we first have define all of the different properties for within the struct.
For a person struct, we could do something similar to the following:
```
type person struct {
    firstName string
	lastName string
}
```
There are multiple ways to instantiate a new variable of type person as defined above. You can do any of the following.
```
newPerson := person{"first", "last"}
newPerson := person{firstName: "first", lastName: "last"}
var newPerson person
```
You can see all of the different variables within a struct and their values by doing the following:
`fmt.Printf("%+v", alex)`

## Embedded Structs
You are able to embed structs inside of other structs.
```
type contactInfo struct {
    email string
}

type person struct {
    name string
    age int
    contact contactInfo
}
```

## Receiver Functions on Structs
Similar to other type definitions, we can have receiver functions on structs with the usual format.
```
func (p person) print () {
    fmt.Println(p.firstName, p.lastName)
	fmt.Println(p.contact.email)
	fmt.Println(p.contact.zipCode)
}
```

## Pointers
Pointers in Go are relatively straight forward and fairly simple. Just like in other languages,
we can access the memory address that a variable is at by using:
`&variable`
and we can get the value of a memory address by using (dereference the pointer)
`*pointer`
When using pointers with functions, we do not have to get the address of the variable we are passing
into a function before passing it in, Go will handle it automatically.
```
myItem := "hey there"
myFunc(myItem)

func myFunc(val *string) {
    fmt.Println(val)
}
```

* Gotcha: In the following we would not expect s[0] to be changed, but here it is*
```
func main () {
    mySlice := []string{"Hi", "There", "How"}
    updateSlice(mySlice)
    fmt.Println(mySlice)
}

func updateSlice(s []string) {
    s[0] = "Bye"
}
```
The cause of this is because of how slices are implemented. The slice data structure is a pointer to
an underlying array. When we pass a slice to a function, the underlying array that it is pointing to
is still the same.

### Difference between arrays and slices
Arrays are a very primitive data structure that can't be resized and are rarely used directly.
Slices can grow and shrink dynamically and is used almost all of the time and is an extension
of the array data structure.

Slices are implemented with three different properties
1. ptr to head
2. capacity
3. length

The actual data (where the ptr to head points to) is actually stored within an array.

## Data structures that are passed by value vs by reference
### Value Types (Need to use pointers with these)
1. int
2. float
3. string
4. bool
5. structs

### Reference Types (No need to use pointers)
1. slices
2. maps
3. channels
4. pointers
5. functions