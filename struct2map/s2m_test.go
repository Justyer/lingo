package s2m

import (
	"fmt"
	"testing"
)

func TestKeyByTagToString(t *testing.T) {
	type X struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	var x X
	x.A = 1
	x.B = "s"
	m, err := KeyByTagToString(x, "json")
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("%#v\n", m)
}

func TestKeyByTag(t *testing.T) {
	type X struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	var x X
	x.A = 2
	x.B = "t"
	m := KeyByTag(x, "json")
	fmt.Printf(":%#v\n", m)
}

func BenchmarkKeyByTagToString(b *testing.B) {
	type X struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	var x X
	KeyByTagToString(x, "json")
}
