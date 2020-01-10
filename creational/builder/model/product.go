package model

type house struct {
	windowType string
	doorType   string
	floor      int
}

func (h house) GetWindowType() string {
	return h.windowType
}

func (h house) GetDoorType() string {
	return h.doorType
}

func (h house) GetNumFloor() int {
	return h.floor
}
