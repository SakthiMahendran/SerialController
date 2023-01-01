package inputsimulator

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
)

type mouseSimulator struct {
	portReader bufio.Reader
}

func (m *mouseSimulator) Click() error {
	btn, err := m.portReader.ReadByte()
	if err != nil {
		return err
	}

	switch btn {
	case MOUSE_LEFT:
		robotgo.Click("left")

	case MOUSE_MIDDLE:
		robotgo.Click("center")

	case MOUSE_RIGHT:
		robotgo.Click("right")
	}

	return nil
}

func (m *mouseSimulator) Move() error {
	x, err := m.readInt()
	if err != nil {
		return err
	}

	y, err := m.readInt()
	if err != nil {
		return err
	}

	robotgo.MoveRelative(x, y)
	return nil
}

func (m *mouseSimulator) Press() error {
	btn, err := m.portReader.ReadByte()
	if err != nil {
		return nil
	}

	switch btn {
	case MOUSE_LEFT:
		robotgo.Toggle("left", "down")

	case MOUSE_MIDDLE:
		robotgo.Toggle("center", "down")

	case MOUSE_RIGHT:
		robotgo.Toggle("right", "down")
	}

	return nil
}

func (m *mouseSimulator) Release() error {
	btn, err := m.portReader.ReadByte()
	if err != nil {
		return nil
	}

	switch btn {
	case MOUSE_LEFT:
		robotgo.Toggle("left", "up")

	case MOUSE_MIDDLE:
		robotgo.Toggle("center", "up")

	case MOUSE_RIGHT:
		robotgo.Toggle("right", "up")
	}

	return nil
}

func (m *mouseSimulator) Scroll() error {
	x, err := m.readInt()
	if err != nil {
		return err
	}

	robotgo.ScrollRelative(x, 0)
	return nil
}

func (m *mouseSimulator) readInt() (int, error) {
	str, err := m.portReader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	num, _ := strconv.Atoi(strings.TrimSpace(str))

	return num, err
}
