# velocity4go

`velocity4go` is an implementation for Apache Velocity templating
for [Go](https://go.dev). It is inspired from [EscapeVelocity](https://github.com/google/escapevelocity), 
a  subset re-write of original [Velocity](https://velocity.apache.org).

### In Development - basic functionality not yet complete.

## Hacking

```go

import velocity "sangupta.com/velocity4go"

templateString := "Hello World"
template := velocity.ParseTemplate(templateString)

variables := make(map[string]interface{}, 0)

// .... add variables here

text := velocity.Evaluate(template, variables)
fmt.Println("Generated text: " + text)
```

## Author(s)

* [@sangupta](https://github.com/sangupta)

## Dependencies

None.

## License

MIT License.Copyright (c) 2022, Sandeep Gupta.
