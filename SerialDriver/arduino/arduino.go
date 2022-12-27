package arduino

import (
	"bufio"
	"time"

	"go.bug.st/serial"
)

type Arduino struct {
}

func (a *Arduino) WaitAndConnect() (serial.Port, string) {
	defaultMode := serial.Mode{
		BaudRate: 115200,
	}

	portInfoChan := make(chan PortInfo)

	for {
		portNames, err := serial.GetPortsList()

		if err != nil {
			continue
		}

		for _, p := range portNames {
			port, err := serial.Open(p, &defaultMode)

			if err != nil {
				continue
			}

			go a.tryHandShake(p, port, portInfoChan)
		}

		select {
		case portInfo := <-portInfoChan:
			// close(portInfoChan)
			return portInfo.port, portInfo.portName
		case <-time.After(time.Second * 5):
			continue
		}
	}
}

func (a *Arduino) tryHandShake(portName string, port serial.Port, portInfoChan chan PortInfo) {
	signalChan := make(chan struct{})

	portReader := bufio.NewReader(port)

	go func() {
		portReader.ReadString('\n')

		if s, _ := portReader.ReadString('\n'); s == "Hello, I am Arduino\n" {
			signalChan <- struct{}{}
		}

	}()

	select {
	case <-signalChan:
		portInfoChan <- PortInfo{portName: portName, port: port}
	case <-time.After(time.Second * 5):
		port.Close()
	}
}
