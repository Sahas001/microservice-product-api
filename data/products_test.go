package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "Sahas",
		Price: 5,
		SKU: "abc-def-ghi",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
