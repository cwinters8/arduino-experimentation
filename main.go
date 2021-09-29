package main

import (
	"flag"
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	device := flag.String("device", "", "Serial device path such as /dev/ttyACM0")
	flag.Parse()

	adaptor := firmata.NewAdaptor(*device)
	led := gpio.NewLedDriver(adaptor, "13") // what does the 13 pin refer to?
	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			err := led.Toggle()
			if err != nil {
				log.Fatal("Failed to toggle LED: " + err.Error())
			}
		})
	}
	robot := gobot.NewRobot("bot", []gobot.Connection{adaptor}, []gobot.Device{led}, work)
	robot.Start()
}
