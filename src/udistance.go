package main

import (
	"fmt"
	"time"
  	"os"
  	"os/signal"
  	"syscall"

	"github.com/stianeikeland/go-rpio"
)

func main() {

	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	PIN_TRIGGER := 4
    PIN_ECHO := 17

	triPin := rpio.Pin(PIN_TRIGGER)
	echoPin := rpio.Pin(PIN_ECHO)

	triPin.Output()
	echoPin.Input()

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		triPin.Low()
    	echoPin.Low()
    	os.Exit(0)
    }()

    defer rpio.Close()

  	triPin.Low()
  	fmt.Println("Waiting for sensor to settle")
  	time.Sleep(1)

	for {
		
		var start_time, end_time int64

		fmt.Println("Calculating distance")

		// initialize sensor
	  	triPin.High()
		time.Sleep(time.Nanosecond)
		triPin.Low()

		// start echoing
		for echoPin.Read() == 0 {
			start_time = time.Now().UnixNano()
		}
		
		// an echo received
		for echoPin.Read() == 1 {
			end_time = time.Now().UnixNano()
		}

		// calculating result 
		// time in nanosec
		// 17000 = (sound speed (320m/s) / 2 converted to cm/s)
		// 1e9 to get result by sec
	    distance := ( float32(end_time - start_time) * 17000 ) / 1e9

	    // print result
		fmt.Printf("Distance: %.2f cm \n", distance)
		time.Sleep(time.Second)

	}

}