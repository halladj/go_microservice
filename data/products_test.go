package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{Name: "ahmza", Price: 12, SKU: "aaa-aaa-aaa"}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
