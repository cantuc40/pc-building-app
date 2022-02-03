package Motherboard

type Motherboard struct {
	name           string
	company        string
	socket         string
	chipset        string
	pci_slots      int8
	ram_slots      int8
	internal_ports int8
	external_ports int8
	price          float64
}
