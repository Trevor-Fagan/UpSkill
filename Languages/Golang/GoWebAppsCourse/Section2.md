# Section 2 - Overview of the Go Language

## Switch Statements
```
switch myVar {
    case "cat":
        fmt.Println("cat")
    
    case "dog":
        fmt.Println("dog")
    
    default:
        fmt.Println("Unknown")
}
```

## Go mod init
Conventionally you prefix the package with the GitHub repository it is going to live in.

## Table Tests
```
var tests = []struct {
    name string
    dividend float32
    divisor float32
    expected float32
    isErr bool
}{
    {"valid-data", 100.0, 10.0, 10.0, false},
    {"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
    for _, tt := range tests {
        got, err := divide(tt.dividend, tt.divisor)
        if tt.isErr {
            if err == nil {
                t.Error("Expected an error")
            }
        } else {
            if err != nil {
                t.Error("Got an error")
            }
        }

        if got != tt.expected {
            t.Error("Did not get expected value")
        }
    }
}
```

### Check test coverage
`go test -coverprofile=coverage.out && go tool cover -html coverage.out`