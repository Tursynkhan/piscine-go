package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func CloseDoor(Door *Door) bool {
	PrintStr("Door Closing...")
	Door.state = CLOSE
	return true
}

func OpenDoor(Door *Door) {
	PrintStr("Door Opening...")
	if Door.state == false {
		Door.state = true
	}
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	return Door.state == OPEN
}

func IsDoorClose(Door *Door) bool {
	PrintStr("is the Door closed ?")
	return Door.state == CLOSE
}

const (
	OPEN  = true
	CLOSE = false
)

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
