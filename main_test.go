package main

import "testing"

func TestXxx(t *testing.T) {
	m := map[string]any{
		"foo": "FFF",
		"bar": "BBB",
	}
	marshal(m)
	t.Error("")
}
