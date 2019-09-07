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
	xValues := []int{10, 11, 12}
	yValues := []int{1, 2, 3}

	for _, val := range xValues {
		x.Add(val)
	}

	for _, val := range yValues {
		y.Add(val)
	}
	x.UnionWith(&y)

	var result bool
	allValues := append(xValues, yValues...)
	for _, value := range allValues {
		result = x.Has(value)
		if result == false {
			t.Errorf("UnionWith was incorrect, got: %t, wanted: %t.", result, true)
		}
	}
}

func TestUnionWithDups(t *testing.T) {
	var x, y IntSet
	xValues := []int{10, 11, 12}
	yValues := []int{1, 10, 12, 2, 3}

	for _, val := range xValues {
		x.Add(val)
	}

	for _, val := range yValues {
		y.Add(val)
	}

	x.UnionWith(&y)

	var result bool
	allValues := append(xValues, yValues...)
	for _, value := range allValues {
		result = x.Has(value)
		if result == false {
			t.Errorf("UnionWith was incorrect, got %t, wanted %t for value: %d.",
				result, true, value)
		}
	}
}

func TestLen(t *testing.T) {
	var x IntSet
	xValues := []int{10, 11, 12, 13, 59}

	for _, val := range xValues {
		x.Add(val)
	}

	setLen := x.Len()
	if setLen != len(xValues) {
		t.Errorf("Len was incorrect, got %d, wanted %d.", setLen, len(xValues))
	}
}

func TestElems(t *testing.T) {
	var x IntSet
	xValues := []int{10, 11, 12, 13, 59}

	for _, val := range xValues {
		x.Add(val)
	}

	result := x.Elems()
	for i, val := range xValues {
		if result[i] != val {
			t.Errorf("Elems was incorrect at index %d, got %d, wanted %d.", i, result[i], val)
		}
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	xValues := []int{1, 2, 3, 4, 5, 6, 7}
	x.AddAll(xValues...)

	var result bool
	for _, val := range xValues {
		result = x.Has(val)
		if result != true {
			t.Errorf("AddAll was incorrect for value %d, got %t, wanted %t.", val, result, true)
		}
	}
}
