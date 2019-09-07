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
	values := []int{1, 2, 3, 6, 8, 9}

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
	x_values := []int{10, 11, 12}
	y_values := []int{1, 2, 3}

	for _, val := range x_values {
		x.Add(val)
	}

	for _, val := range y_values {
		y.Add(val)
	}
	x.UnionWith(&y)

	var result bool
	all_values := append(x_values, y_values...)
	for _, value := range all_values {
		result = x.Has(value)
		if result == false {
			t.Errorf("UnionWith was incorrect, got: %t, wanted: %t.", result, true)
		}
	}
}
