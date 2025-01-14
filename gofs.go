package gofs

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

type Details struct {
	Name       string
	Parameters []genai.Schema
}

func Get(f any) any {
	sig := fmt.Sprintf("%T", f)
	p := destructureParams(sig)
	n := getFunctionName(f)
	items := strings.Split(n, ".")
	return Details{
		items[len(items)-1],
		p,
	}
}

func destructureParams(s string) []genai.Schema {
	var (
		start   int
		end     int
		schemas []genai.Schema
	)
	for a, b := range s {
		if b == '(' {
			start = a
		}
		if b == ')' {
			end = a
		}
	}
	params := strings.Split(s[start:end], ", ")
	for _, a := range params {
		var t genai.Type
		if a[0:1] == "[]" && (a[2:5] != "rune" && a[2:5] != "byte") {
			t = genai.TypeArray
		} else {
			switch a {
			case "bool":
				t = genai.TypeBoolean
			case "string", "[]byte", "[]rune":
				t = genai.TypeString
			case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune":
				t = genai.TypeInteger
			case "float32", "float64", "complex64", "complex128":
				t = genai.TypeNumber
			default:
				t = genai.TypeObject
			}
		}
		schemas = append(schemas, genai.Schema{
			Type: t,
		})
	}
	return schemas
}

// credit: https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
