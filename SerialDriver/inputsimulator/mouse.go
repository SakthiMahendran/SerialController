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

func (m *mouseSimulator) leftClick() {
	robotgo.Click("left")
}

func (m *mouseSimulator) rightClick() {
	robotgo.Click("right")
}

func (m *mouseSimulator) scroll() error {
	x, err := m.readInt()
	if err != nil {
		return err
	}

	y, err := m.readInt()
	if err != nil {
		return err
	}

	robotgo.Scroll(x, y)
	return nil
}

func (m *mouseSimulator) scrollRelative() error {
	x, err := m.readInt()
	if err != nil {
		return err
	}

	y, err := m.readInt()
	if err != nil {
		return err
	}

	robotgo.ScrollRelative(x, y)
	return nil
}

func (m *mouseSimulator) move() error {
	x, err := m.readInt()
	if err != nil {
		return err
	}

	y, err := m.readInt()
	if err != nil {
		return err
	}

	robotgo.Move(x, y)
	return nil
}

func (m *mouseSimulator) moveRelative() error {
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

func (m *mouseSimulator) readInt() (int, error) {
	str, err := m.portReader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	num, _ := strconv.Atoi(strings.TrimSpace(str))

	return num, err
}
