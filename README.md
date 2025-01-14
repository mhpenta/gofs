# gofs

A Go module that converts functions with struct parameters into OpenAPI/JSON Schema descriptions. Fork and reworking of github.com/samjtro/gofs.

## Requirements
- Functions must accept exactly one struct parameter
- Parameter struct fields should use json tags for explicit naming

## Usage

```go
import "github.com/mhpenta/gofs"

type Params struct {
    world string `json:"world"`
}

func helloTool(p Params) {}

details, err := gofs.GetFunctionDetails(helloTool)
/* returns:
{
    "name": "helloTool",
    "parameters": [{
        "title": "world",
        "type": "string"
    }]
}*/

```

## License
GNU General Public License v2
