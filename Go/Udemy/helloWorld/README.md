# package

```go
package main

import "fmt"

func main() {
	fmt.Println("Hi There")
}
```

There are two types of packages. All the go files which are related to the same package should have a package name same. It helps in grouping together code with a similar purpose.

1. Executable
2. Reusable

## Executable

Generates a file that we can run. If we use the package name as main. Then it will become an executable package. We need a function named main.

```go
package main
```

## Reusable

Code used as helpers. A good place to put reusable logic. If we use any other name other than main. It will be considered as reusable. These are good for making library files.

```go
package <name>
```

# import

This helps in importing standard library and resuable packages. This will give our package access to code written in another package.

```go
import <package_name>
```

# function

The function will have the following format. This function will get executed when we run the program.

```go
func main(arg1 , arg2) {

}
```

# File Organization

| Order                                                           | Notes                                   |
| --------------------------------------------------------------- | --------------------------------------- |
| `package main `                                                 | package declaration                     |
| `import fmt `                                                   | Import other packages that we need      |
| <code>func main() { <br> fmt.Println("Hi There") <br> } </code> | Declare functions, This will have logic |
