package greeting

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Gladys"
	msg, err := Hello(name)

	want := regexp.MustCompile(`\b` + name + `\b`)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
