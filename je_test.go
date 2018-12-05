package je

import (
	"testing"
)

func TestGet(t *testing.T) {
	str := `{"k1":{"kk1":"vv1","kk2":2},"k2":false,"k3":[{"k31":"v31"},{"k32":"v32"},{"k33":"v33"}]}`

	// get string
	r := Get(str, ".k1.kk1")
	if r != "vv1" {
		t.Fail()
	}

	// get string
	r = Get(str, ".k1.kk2")
	if r != "2" {
		t.Fail()
	}

	// get string
	r = Get(str, ".k2")
	if r != "false" {
		t.Fail()
	}

	// get string
	r = Get(str, ".k3.0.k31")
	if r != "v31" {
		t.Fail()
	}

}

func TestSet(t *testing.T) {
	str := `{"k1":{"kk1":"vv1","kk2":"vv2"},"k2":"v2"}`

	// set struct
	r := Set(str, ".k2", map[string]string{"kk1":"vv1","kk2":"vv0"})
	if r != `{"k1":{"kk1":"vv1","kk2":"vv2"},"k2":{"kk1":"vv1","kk2":"vv0"}}` {
		t.Fail()
	}

	// set array
	r = Set(str, ".k2", []map[string]string{{"kk1":"vv1"},{"kk2":"vv0"}})
	if r != `{"k1":{"kk1":"vv1","kk2":"vv2"},"k2":[{"kk1":"vv1"},{"kk2":"vv0"}]}` {
		t.Fail()
	}

	// set int
	r = Set(str, ".k1.kk2", 2)
	if r != `{"k1":{"kk1":"vv1","kk2":2},"k2":"v2"}` {
		t.Fail()
	}

	// set bool
	r = Set(str, ".k1.kk2", false)
	if r != `{"k1":{"kk1":"vv1","kk2":false},"k2":"v2"}` {
		t.Fail()
	}

	// set string
	r = Set(str, ".k1", "vv0")
	if r != `{"k1":"vv0","k2":"v2"}` {
		t.Fail()
	}

	// set string
	r = Set(str, ".k1.kk2", "2")
	if r != `{"k1":{"kk1":"vv1","kk2":"2"},"k2":"v2"}` {
		t.Fail()
	}

	// set string
	r = Set(str, ".k1.kk2", "false")
	if r != `{"k1":{"kk1":"vv1","kk2":"false"},"k2":"v2"}` {
		t.Fail()
	}

}

func TestGetByte(t *testing.T) {
	str := []byte(`{"k1":{"kk1":"vv1","kk2":"vv2"},"k2":"v2"}`)
	path := []byte(".k1.kk2")
	r := GetB(str, path)
	if string(r) != "vv2" {
		t.Fail()
	}
}

func TestSetByte(t *testing.T) {
	str := []byte(`{"k1":{"kk1":"vv1","kk2":"vv2"},"k2":"v2"}`)
	path := []byte(".k1")
	value := []byte(`{"kk1":"vv1","kk2":"vv0"}`)
	r := SetB(str, path, TypeB(value))
	if string(r) != `{"k1":{"kk1":"vv1","kk2":"vv0"},"k2":"v2"}` {
		t.Fail()
	}
}
