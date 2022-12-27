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
		case mouseLeftClick:
			mSimulator.leftClick()

		case mouseRightClick:
			mSimulator.rightClick()

		case mouseScroll:
			mSimulator.scroll()

		case mouseScrollRelative:
			mSimulator.scrollRelative()

		case mouseMove:
			mSimulator.move()

		case mouseMoveRelative:
			mSimulator.moveRelative()
		}
	}
}
