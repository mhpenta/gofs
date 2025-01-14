# gofs

Fork and reworking of github.com/samjtro/gofs

License: GNU GENERAL PUBLIC LICENSE Version 2

Requires functions to be passed a single struct as a parameter. In that context, the code is able to deconstruct the function and parameters into a LLM tool call.

## usage

```go
import "github.com/samjtro/gofs"

func hello(world string) {}

gofs.GetFunctionDetails(hello)
/* returns {
    Name: "hello",
    Parameters: []Schema{
        Type: TypeString,
    },
}*/
```
