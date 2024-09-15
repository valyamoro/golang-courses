package main

import "fmt"

type Celsius float64

func FahrenheitToCelsius(fahrenheit float64) (result Celsius) {
	result = Celsius((fahrenheit - 32) * 5 / 9)

	return
}

func KelvinToCelsius(kelvin float64) (result Celsius) {
	result = Celsius(kelvin - 273.15)

	return
}

func TaskTwentyOne() {
	fahrenheit := 9.0
	kelvin := 15.0

	fmt.Println(FahrenheitToCelsius(fahrenheit))
	fmt.Println(KelvinToCelsius(kelvin))
}
