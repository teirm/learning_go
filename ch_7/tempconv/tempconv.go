package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius(5.0 * (f - 32) / 9.0) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.0) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.0) }
func KToF(k Kelvin) Fahrenheit  { return CToF(KToC(k)) }
func FToK(f Fahrenheit) Kelvin  { return CToK(FToC(f)) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (k Kelvin) String() string  { return fmt.Sprintf("%gK", k) }

type celsiusFlag struct{ Celsius }
type kelvinFlag struct{ Kelvin }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Kelvin = CToK(Celsius(value))
		return nil
	case "F", "°F":
		f.Kelvin = FToK(Fahrenheit(value))
		return nil
	case "K":
		f.Kelvin = Kelvin(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}
