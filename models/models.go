package models

type Case struct {
	name     string
	company  string
	windowed bool
	material string
	price    float64
}

type CPU struct {
	name       string
	company    string
	cores      int8
	clockspeed string
	socket     string
	material   string
	price      float64
}

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

type Memory struct {
	frequency       string
	cas_latency     string
	memory_channels string
}

type Storage struct {
	type_    string
	capacity string
	speed    string
}

type PowerSupply struct {
	form_factor string
	modularity  string
}

type GraphicsCard struct {
	gpu        string
	memory     string
	core_clock string
}

type Monitor struct {
	resolution    string
	hz            int8
	response_time int8
	display_lag   string
	panel_type    string
}
