# Section 2 - A Simple Start
## Common Go Commands
- `go run`      => compile and run program, can be more than one
- `go build`    => compile program, will produce executable. Can be more than one
- `go fmt`      => formats all code in each file in current directory
- `go install`  => compiles and installs a package
- `go get`      => downloads the raw source code of a package
- `go test`     => Runs any tests associated with the current project

## Packages
A related set of programs that work together to perform some function. What you think
of when you think of packages. Every Go file must declare the package they are part of
as the first line of the file. 

There are two types of packages in Go- an executable type and a reusable type. The 
executable type generates an executable whereas the reusable type is code used as helpers
that often contain reusable logic. The name of the package is what determines the type.
"main" is used for all executable packages, and anything else is a reusable type.