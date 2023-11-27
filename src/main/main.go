package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/thoas/go-funk"
)

//cars
//	.filter(car => car.driver.license)
//	.filter(car => car.carName.contains("Holden"))
//	.map(car => car.driver)
//	.map(driver => driver.name)

func main() {
	//car := Car{driver: Person{name: "Jaque", licence: true}}
	//motorbike := Motorbike{rider: Person{name: "Brendon", licence: false}}
	//pullOverVehicle(car)
	//pullOverVehicle(motorbike)
	cars := []Car{}
	for range [10]int{} {
		car := Car{driver: Person{name: gofakeit.Name(), licence: gofakeit.Bool()}, carCompany: gofakeit.CarMaker()}
		cars = append(cars, car)
	}
	fmt.Println("cars", cars)
	licencedCars := funk.Filter(cars, func(c Car) bool {
		return c.driver.licence
	})
	newlySuspendedDrivers :=
		funk.Map(
			funk.Map(licencedCars, func(c Car) Person { return c.driver }),
			func(p Person) Person { return p.suspendLicense() },
		)
	//newlySuspendedDriverNames := funk.Map(newlySuspendedDrivers, func(p Person) string { return p.name })

	licencedNames := funk.Map(licencedCars, driversName)
	nonLicencedCars := funk.Filter(cars, func(c Car) bool {
		return !c.driver.licence
	})
	nonlicencedNames := funk.Map(nonLicencedCars, driversName)
	fmt.Println("newlySuspendedDriverNames", newlySuspendedDrivers)
	fmt.Println("licenced cars", licencedCars)
	fmt.Println("licenced names", licencedNames)
	fmt.Println("non licenced cars", nonLicencedCars)
	fmt.Println("non licenced names", nonlicencedNames)
	fmt.Println("cars", cars)
}

//func suspendLicense(p Person) Person {
//	newPerson := p
//	newPerson.licence = false
//	return newPerson
//}

func driversName(c Car) string {
	return c.driver.name
}

func pullOverVehicle(vehicle Vehicle) {
	fmt.Println(vehicle.Stop())
	fmt.Println(vehicle.Operator().name)
	if vehicle.Operator().licence {
		fmt.Println(vehicle.Start())
	} else {
		fmt.Println("You're under arrest!")
	}
}

type Vehicle interface {
	Start() string
	Stop() string
	Operator() Person
}

type Car struct {
	driver     Person
	carCompany string
}

func (c Car) Start() string {
	return "car starting"
}

func (c Car) Stop() string {
	return "car stoping"
}

func (c Car) Operator() Person {
	return c.driver
}

type Motorbike struct {
	rider Person
}

func (m Motorbike) Start() string {
	return "motorbike starting"
}

func (m Motorbike) Stop() string {
	return "motorbike stoping"
}

func (m Motorbike) Operator() Person {
	return m.rider
}

type Person struct {
	name    string
	licence bool
}

func (p Person) suspendLicense() Person {
	newPerson := p
	newPerson.licence = false
	return newPerson
}
