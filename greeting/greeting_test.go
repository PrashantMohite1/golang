package greeting

import (
	"regexp"
	"testing"
)

func TestPrintName(t *testing.T) {
	name := "prashant"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, error := PrintName(name)

	if !want.MatchString(message) || error != nil {
		t.Errorf(`PrintName(name) = %q, %v, want match for %#q, nil`, message, error, want)
	}

}
