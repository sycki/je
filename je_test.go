package main

import (
	"testing"
)

func TestGet(t *testing.T) {
	str := `{"k1":{"kk1":"vv1","kk2":"vv2"},"k2":"v2"}`
	r := Get(str, ".k1.kk2")
	if r != `"vv2"` {
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	str := `{"k1": {"kk1": "vv1", "kk2": "vv2"}, "k2": "v2"}`
	r := Set(str, ".k1.kk2", "vv0")
	if r != `{"k1":{"kk1":"vv1","kk2":"vv0"},"k2":"v2"}` {
		t.Fail()
	}
}
