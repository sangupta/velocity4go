# velocity4go

`velocity4go` is an implementation for Apache Velocity templating
for [Go](https://go.dev). It is inspired from [EscapeVelocity](https://github.com/google/escapevelocity), 
a  subset re-write of original [Velocity](https://velocity.apache.org).

### In Development - basic functionality not yet complete.

TODO items:
* Template parsing
    * ~if/elseif/else directive~
    * foreach directive
    * set directive
    * custom directives: include/user-defined
    * macros
* Evaluation
    * ~parameter evaluation~
    * foreach evaluation
    * set evaluation
    * if/elseif/else evaluation
    * custom directive evaluation
    * macro evaluation
* Unit tests

## Hacking

```go
package main

import (
    "fmt"
    velocity "sangupta.com/velocity4go"
)

func main() {
    templateString := "Hello World"
    template := velocity.ParseTemplate(templateString)

    variables := make(map[string]interface{}, 0)

    // .... add variables here

    text := velocity.Evaluate(template, variables)
    fmt.Println("Generated text: " + text)
}
```

## Author(s)

* [@sangupta](https://github.com/sangupta)

## Dependencies

None.

## License

MIT License.Copyright (c) 2022, Sandeep Gupta.
