package main

import "fmt"

type MotarVehicle interface {
	mileage() float64
}

type BMW struct {
	distance float64
	fuel     float64
	avgspeed string
}

type AUDI struct {
	distance float64
	fuel     float64
}

func (b BMW) mileage() float64 {
	return b.distance / b.fuel
}
func (a AUDI) mileage() float64 {
	return a.distance / a.fuel
}

func totalMilage(m []MotarVehicle) {
	tm := 0.0
	for _, v := range m {
		tm = tm + v.mileage()
	}
	fmt.Printf("The Total milage per month %f km/l", tm)
}

func main() {
	b1 := BMW{
		distance: 155.5,
		fuel:     15.5,
		avgspeed: "60",
	}
	a1 := AUDI{
		distance: 254.7,
		fuel:     34.6,
	}
	person := []MotarVehicle{b1, a1}
	totalMilage(person)

}
