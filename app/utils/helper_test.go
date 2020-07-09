package utils

import "testing"

func TestIsEmpty(t *testing.T) {
	v := IsEmpty("")
	if v != true {
		t.Errorf("Expect true")
	}

	v2 := IsEmpty(nil)
	if v2 != true {
		t.Errorf("Expect true")
	}

	v3 := IsEmpty(0)
	if v3 != true {
		t.Errorf("Expect true")
	}

	v4 := IsEmpty(" ")
	if v4 != false {
		t.Errorf("Expect false")
	}

	v5 := IsEmpty("eiei")
	if v5 != false {
		t.Errorf("Expect false")
	}
}
