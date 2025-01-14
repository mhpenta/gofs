# gofs

## usage

```go
import "github.com/samjtro/gofs"

func hello(world string) {}

gofs.Get(hello)
/* returns {
    Name: "hello",
    Parameters: []genai.Schema{
        Type: genai.TypeString,
        Description: world,
    },
}*/
```
