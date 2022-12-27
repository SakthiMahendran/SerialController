package main

import (
	"fmt"

	"github.com/SakthiMahendran/SerialDriver/arduino"
	"github.com/SakthiMahendran/SerialDriver/inputsimulator"
)

func main() {
	arduino := arduino.Arduino{}

	for {
		fmt.Println("Waiting for Arduino")
		port, portName := arduino.WaitAndConnect()
		fmt.Println("Connected with Arduino at ", portName)

		is := inputsimulator.InputSimulator{
			Port: port,
		}

		fmt.Println("Simulating")
		err := is.StartSimulating()
		fmt.Println("Disconnected due to ", err)

		port.Close()
	}

}
