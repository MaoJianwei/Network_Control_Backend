package lib

type MaoPort struct {
	id uint16
	name string

	device MaoDevice
	link MaoLink
}