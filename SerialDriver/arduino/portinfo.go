package arduino

import "go.bug.st/serial"

type PortInfo struct {
	portName string
	port     serial.Port
}
