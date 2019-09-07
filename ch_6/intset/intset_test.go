package intset

import "testing"

func TestHas(t *testing.T) {
	var x IntSet
	result := x.Has(10)
	if result == true {
		t.Errorf("Has was incorrect, got: %t, wanted: %t.", result, false)
	}
}

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(10)
	result := x.Has(10)
	if result == false {
		t.Errorf("Add was incorect, got: %t, wanted: %t.", result, true)
	}
}

func TestAddMulti(t *testing.T) {
	var x IntSet
	values := [6]int{1, 2, 3, 6, 8, 9}

	var result bool
	for _, value := range values {
		x.Add(value)
		result = x.Has(value)
		if result == false {
			t.Errorf("Add was incorect, got: %t, wanted: %t.", result, true)
		}
	}
}

func TestUnionWithNoDups(t *testing.T) {
	var x, y IntSet
	values := [6]int{10, 11, 12, 1, 2, 3}

	x.Add(10)
	x.Add(11)
	x.Add(12)

	y.Add(1)
	y.Add(2)
	y.Add(3)

	x.UnionWith(&y)

	var result bool
	for _, value := range values {
		result = x.Has(value)
		if result == false {
			t.Errorf("UnionWith was incorrect, got: %t, wanted: %t.", result, true)
		}
	}
}
