# Interfaces
An interface type in Go is like a definition. It defines and describes the exact
methods that some other type must have. It makes it so that we are able to use the same
function to perform actions on different data types.
```
type bot interface {
	getGreeting() string
}
```
The above will allow us to then call the getGreeting() with our own types that will correspond
to the individual receiver.

You can also add other functions to the interface. This will make it so that the bot interface
will only be applied to the types that implement every single function listed in the interface
defintion. For example
```
type bot interface {
    getGreeting(string, int) (string, error)
    getBotVersion() float64
}
```
To use this interface, your type would have to have the exact function definitions listed above.

There are concrete types and interface types. With concrete types you are able to instantiate
a variable of this type. These include data structures such as map, struct, int,
string, englishBot. With interface types, you would not be able to create a variable of this
type. For example, the bot interface.

Interfaces are implicit, you do not have to make a type definition "implement" it.

If you specify an interface as a value inside of a struct, it means that the struct can have 
any value assigned to it as long as it matches the interface. Interfaces are able to be combined,
or another way of saying it is you can contain other interfaces inside a primary interface.
