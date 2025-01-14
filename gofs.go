package gofs

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Details struct {
	Name       string   `json:"name"`
	Parameters []Schema `json:"parameters"`
}

type Parameters struct {
	Name string
	Type reflect.Type
}

func GetFunctionDetails(f any) (*Details, error) {
	paramType, _, err := getFunctionTypes(f)
	if err != nil {
		return nil, err
	}
	p := destructureParamsFromType(paramType)
	n := getFunctionName(f)
	items := strings.Split(n, ".")
	return &Details{
		items[len(items)-1],
		p,
	}, nil
}

func destructureParamsFromType(t reflect.Type) []Schema {
	fmt.Println(t)
	var schemas []Schema
	params := getParameters(t)
	for _, p := range params {
		var s Schema
		s.Title = p.Name
		switch p.Type.String() {
		case "bool":
			s.Type = TypeBoolean
		case "string", "[]byte", "[]rune":
			s.Type = TypeString
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune":
			s.Type = TypeInteger
		case "float32", "float64", "complex64", "complex128":
			s.Type = TypeNumber
		default:
			s.Type = TypeObject
		}
		schemas = append(schemas, s)
	}
	return schemas
}

func getParameters(t reflect.Type) []Parameters {
	params := make([]Parameters, 0)
	if t.Kind() != reflect.Struct {
		panic("Not a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("json")

		tagParts := strings.Split(tag, ",") // splitting for omitempty, omitzero etc
		tagName := tagParts[0]

		if tagName == "" {
			tagName = strings.ToLower(field.Name)
		}

		params = append(params, Parameters{
			Name: field.Name,
			Type: field.Type,
		})
	}

	return params
}

func getFunctionTypes(f interface{}) (paramType reflect.Type, returnTypes []reflect.Type, err error) {
	fType := reflect.TypeOf(f)

	if fType.Kind() != reflect.Func {
		return nil, nil, fmt.Errorf("not a function")
	}

	numParams := fType.NumIn()

	if numParams != 1 {
		return nil, nil, fmt.Errorf("tool function must have exactly one parameter")
	}

	paramType = fType.In(0)

	if paramType.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("parameter must be a struct")
	}

	numReturns := fType.NumOut()
	returnTypes = make([]reflect.Type, numReturns)
	for i := 0; i < numReturns; i++ {
		returnTypes[i] = fType.Out(i)
	}

	return paramType, returnTypes, nil
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
