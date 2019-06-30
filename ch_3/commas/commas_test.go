package commas

import "testing"

func errorMessage(result string, expected string, t *testing.T) {
	t.Errorf("Comma placement was incorrect, got: %s, want %s.", result, expected)
}

func TestComma1(t *testing.T) {
	expected := "12,345"

	s := Comma1("12345")
	if s != expected {
		errorMessage(s, expected, t)
	}
}

func TestComma2(t *testing.T) {
	expected := "12,345"

	s := Comma2("12345")
	if s != expected {
		errorMessage(s, expected, t)
	}
}

func TestComma3(t *testing.T) {
	expected := "12,345"

	s := Comma3("12345")
	if s != expected {
		errorMessage(s, expected, t)
	}
}

func TestComma3Decimal(t *testing.T) {
	expected := "12,345.123"

	s := Comma3("12345.123")
	if s != expected {
		errorMessage(s, expected, t)
	}
}

func TestComma3OnlyDecimal(t *testing.T) {
	expected := ".123"
	s := Comma3(".123")
	if s != expected {
		errorMessage(s, expected, t)
	}
}

func TestComma3PositiveNoDecimal(t *testing.T) {
	expected := "+12,345"
	s := Comma3("+12345")
	if s != expected {
		errorMessage(s, expected, t)
	}
}

func TestComma3NegativeNoDecimal(t *testing.T) {
	expected := "-12,345"
	s := Comma3("-12345")
	if s != expected {
		errorMessage(s, expected, t)
	}
}

func TestComma3PositiveDecimal(t *testing.T) {
	expected := "+12,345.123"
	s := Comma3("+12345.123")
	if s != expected {
		errorMessage(s, expected, t)
	}
}

func TestComma3NegativeDecimal(t *testing.T) {
	expected := "-12,345.123"
	s := Comma3("-12345.123")
	if s != expected {
		errorMessage(s, expected, t)
	}
}
