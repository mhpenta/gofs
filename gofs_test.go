package gofs_test

import (
	"testing"

	"github.com/mhpenta/gofs"
)

type singleParam struct {
	world string `json:"world"`
}

func singleParamFunc(p singleParam) {}

type customParam struct {
	world string `json:"worldParam"`
}

func customParamFunc(p customParam) {}

type multiParam struct {
	hello string `json:"hello"`
	world string `json:"world"`
}

func multiParamFunc(p multiParam) {}

type typeParam struct {
	str   string  `json:"str"`
	num   float64 `json:"num"`
	flag  bool    `json:"flag"`
	count int     `json:"count"`
}

func typeParamFunc(p typeParam) {}

func TestGetFunctionDetails(t *testing.T) {
	tests := []struct {
		name       string
		fn         interface{}
		wantName   string
		wantParams []string
		wantTypes  []gofs.Type
	}{
		{
			name:       "single parameter",
			fn:         singleParamFunc,
			wantName:   "singleParamFunc",
			wantParams: []string{"world"},
			wantTypes:  []gofs.Type{gofs.TypeString},
		},
		{
			name:       "custom json tag",
			fn:         customParamFunc,
			wantName:   "customParamFunc",
			wantParams: []string{"worldParam"},
			wantTypes:  []gofs.Type{gofs.TypeString},
		},
		{
			name:       "multiple parameters",
			fn:         multiParamFunc,
			wantName:   "multiParamFunc",
			wantParams: []string{"hello", "world"},
			wantTypes:  []gofs.Type{gofs.TypeString, gofs.TypeString},
		},
		{
			name:       "different types",
			fn:         typeParamFunc,
			wantName:   "typeParamFunc",
			wantParams: []string{"str", "num", "flag", "count"},
			wantTypes:  []gofs.Type{gofs.TypeString, gofs.TypeNumber, gofs.TypeBoolean, gofs.TypeInteger},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			details, err := gofs.GetFunctionDetails(tt.fn)
			if err != nil {
				t.Fatalf("GetFunctionDetails() error = %v", err)
			}

			if details.Name != tt.wantName {
				t.Errorf("Name = %v, want %v", details.Name, tt.wantName)
			}

			if len(details.Parameters) != len(tt.wantParams) {
				t.Errorf("Got %d parameters, want %d", len(details.Parameters), len(tt.wantParams))
			}

			for i, want := range tt.wantParams {
				if details.Parameters[i].Title != want {
					t.Errorf("Parameter[%d] = %v, want %v", i, details.Parameters[i].Title, want)
				}
				if details.Parameters[i].Type != tt.wantTypes[i] {
					t.Errorf("Parameter[%d] type = %v, want %v", i, details.Parameters[i].Type, tt.wantTypes[i])
				}
			}
		})
	}
}
