package inputsimulator

import (
	"bufio"

	"go.bug.st/serial"
)

type InputSimulator struct {
	Port serial.Port
}

func (is *InputSimulator) StartSimulating() error {

	portReader := bufio.NewReader(is.Port)
	mSimulator := mouseSimulator{portReader: *portReader}

	for {
		inputType, err := portReader.ReadByte()

		if err != nil {
			return err
		}

		if inputType == 0 {
			continue
		}

		switch inputType {
		case Click:
			mSimulator.Click()

		case Move:
			mSimulator.Move()

		case Press:
			mSimulator.Press()

		case Release:
			mSimulator.Release()

		case End:
			return nil

		case Scroll:
			mSimulator.Scroll()
		}
	}
}
