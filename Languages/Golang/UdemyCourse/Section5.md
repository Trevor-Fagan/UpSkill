# Section 5 - Maps
Similar to other languages, a map is a hash map/dictionary. All keys in a map and all values in a map
must be of the same type. The keys and the values can be different types.

There are multiple different ways to declare a map in Go, such as
```
colors := map[string]string{
    "red": "ff0000",
    "green": "dsa231",
}

var colors map[string]string

colors := make(map[string]string)
```

Delete from a map
`delete(mapName, "key")`