package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"time"
)

func main() {
	gbot := gobot.NewGobot()
	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")

	sensor := gpio.NewAnalogSensorDriver(beagleboneAdaptor, "sensor", "P9_33")
	led := gpio.NewLedDriver(beagleboneAdaptor, "led", "P9_14")

	work := func() {
		gobot.Every(0.1*time.Second, func() {
			val := sensor.Read()
			brightness := uint8(gpio.ToPwm(val))
			fmt.Println("sensor", val)
			fmt.Println("brightness", brightness)
			led.Brightness(brightness)
		})
	}

	gbot.Robots = append(gbot.Robots,
		gobot.NewRobot("sensorBot", []gobot.Connection{beagleboneAdaptor}, []gobot.Device{sensor, led}, work))
	gbot.Start()
}
