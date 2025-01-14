package gofs

// Type contains the list of OpenAPI data types as defined by
// https://spec.openapis.org/oas/v3.0.3#data-types
type Type int32

const (
	// TypeUnspecified means not specified, should not be used.
	TypeUnspecified Type = 0
	// TypeString means string type.
	TypeString Type = 1
	// TypeNumber means number type.
	TypeNumber Type = 2
	// TypeInteger means integer type.
	TypeInteger Type = 3
	// TypeBoolean means boolean type.
	TypeBoolean Type = 4
	// TypeArray means array type.
	TypeArray Type = 5
	// TypeObject means object type.
	TypeObject Type = 6
)

// Schema is the `Schema` object allows the definition of input and output data types.
// These types can be objects, but also primitives and arrays.
//
// Represents a select subset of an [OpenAPI 3.0 schema
// object](https://spec.openapis.org/oas/v3.0.3#schema).
type Schema struct {
	// See Title in the properties part of the spec, https://spec.openapis.org/oas/v3.0.3#properties
	Title string `json:"title"`

	// Required. Data type.
	Type Type `json:"type"`
	// Optional. The format of the data. This is used only for primitive
	// datatypes. Supported formats:
	//
	//	for NUMBER type: float, double
	//	for INTEGER type: int32, int64
	Format string `json:"format,omitempty"`
	// Optional. A brief description of the parameter. This could contain examples
	// of use. Parameter description may be formatted as Markdown.
	Description string `json:"description,omitempty"`
	// Optional. Indicates if the value may be null.
	Nullable bool `json:"nullable,omitempty"`
	// Optional. Possible values of the element of Type.STRING with enum format.
	// For example we can define an Enum Direction as :
	// {type:STRING, format:enum, enum:["EAST", NORTH", "SOUTH", "WEST"]}
	Enum []string `json:"enum,omitempty"`
	// Optional. Schema of the elements of Type.ARRAY.
	Items *Schema `json:"items,omitempty"`
	// Optional. Properties of Type.OBJECT.
	Properties map[string]*Schema `json:"properties,omitempty"`
	// Optional. Required properties of Type.OBJECT.
	Required []string `json:"required,omitempty"`
}
