package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Tavi",
		Price: 20,
		SKU:   "abs-fsdfds-asd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
