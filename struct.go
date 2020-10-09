package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16 // min 0 max 65535
	break_pedal uint16
	steerint_wheel int16 // min -32k max 32k
	top_speed_kmh float64
}

// value receiver method
// makes a copy of the struct
// if the struct is small then more efficient
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

// value receiver method
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

// pointer receiver method
// can modify the original struct
// if the struct is large then more efficient
func (c *car) new_top_speed(new_speed float64) {
	c.top_speed_kmh = new_speed
}

func newer_top_speed(c car, new_speed float64) car {
	c.top_speed_kmh = new_speed
	return c
}

func main() {
	a_car := car{
		gas_pedal: 65000,
		break_pedal: 0,
		steerint_wheel: 12561,
		top_speed_kmh: 225.0,
	}

	fmt.Println(a_car)
	fmt.Println(a_car.gas_pedal)

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car.new_top_speed(500)

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car = newer_top_speed(a_car, 225)

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}