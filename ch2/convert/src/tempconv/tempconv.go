package tempconv

import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
)

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g F", f) }
