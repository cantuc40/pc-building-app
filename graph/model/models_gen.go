// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CPU struct {
	Name       string  `json:"name"`
	ID         string  `json:"id"`
	Company    *string `json:"company"`
	Cores      int     `json:"cores"`
	Clockspeed string  `json:"clockspeed"`
	Sockets    int     `json:"sockets"`
	Price      float64 `json:"price"`
}

type Case struct {
	Name     string  `json:"name"`
	ID       string  `json:"id"`
	Company  *string `json:"company"`
	Windowed *bool   `json:"windowed"`
	Material string  `json:"material"`
	Price    float64 `json:"price"`
}

type DeleteUser struct {
	ID string `json:"id"`
}

type GraphicsCard struct {
	Name      string  `json:"name"`
	ID        string  `json:"id"`
	Company   *string `json:"company"`
	Gpu       string  `json:"gpu"`
	Memory    int     `json:"memory"`
	CoreClock string  `json:"core_clock"`
	Price     float64 `json:"price"`
}

type Memory struct {
	Name           string  `json:"name"`
	ID             string  `json:"id"`
	Company        *string `json:"company"`
	Frequency      string  `json:"frequency"`
	CasLatency     string  `json:"cas_latency"`
	MemoryChannels string  `json:"memory_channels"`
	Price          float64 `json:"price"`
}

type Monitor struct {
	Name       string  `json:"name"`
	ID         string  `json:"id"`
	Company    *string `json:"company"`
	Resolution string  `json:"resolution"`
	Hz         int     `json:"hz"`
	Price      float64 `json:"price"`
}

type Motherboard struct {
	Name          string  `json:"name"`
	ID            string  `json:"id"`
	Company       *string `json:"company"`
	FormFactor    string  `json:"form_factor"`
	Sockets       int     `json:"sockets"`
	Chipsets      int     `json:"chipsets"`
	PciSlots      int     `json:"pci_slots"`
	RAMSlots      int     `json:"ram_slots"`
	InternalPorts int     `json:"internal_ports"`
	ExternalPorts int     `json:"external_ports"`
	Price         float64 `json:"price"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type OperatingSystem struct {
	Name    string  `json:"name"`
	ID      string  `json:"id"`
	Company *string `json:"company"`
	Price   int     `json:"price"`
}

type Pc struct {
	Name            string           `json:"name"`
	ID              string           `json:"id"`
	Motherboard     *Motherboard     `json:"motherboard"`
	CPU             *CPU             `json:"cpu"`
	RAM             *Memory          `json:"ram"`
	Storage         *Storage         `json:"storage"`
	Powersupplie    *PowerSupply     `json:"powersupplie"`
	GraphicCard     *GraphicsCard    `json:"graphic_card"`
	Case            *Case            `json:"case"`
	Monitor         *Monitor         `json:"monitor"`
	OperatingSystem *OperatingSystem `json:"operating_system"`
	Cost            int              `json:"cost"`
}

type Part struct {
	Name    string  `json:"name"`
	ID      string  `json:"id"`
	Company string  `json:"company"`
	Price   float64 `json:"price"`
}

type PartsDb struct {
	Motherboards     []*Motherboard     `json:"motherboards"`
	Cpus             []*CPU             `json:"cpus"`
	Rams             []*Memory          `json:"rams"`
	Storages         []*Storage         `json:"storages"`
	Powersupplies    []*PowerSupply     `json:"powersupplies"`
	GraphicCards     []*GraphicsCard    `json:"graphic_cards"`
	Cases            []*Case            `json:"cases"`
	Monitors         []*Monitor         `json:"monitors"`
	OperatingSystems []*OperatingSystem `json:"operating_systems"`
}

type PowerSupply struct {
	Name       string  `json:"name"`
	ID         string  `json:"id"`
	Company    *string `json:"company"`
	FormFactor string  `json:"form_factor"`
	Modularity string  `json:"modularity"`
	Price      float64 `json:"price"`
}

type Storage struct {
	Name     string  `json:"name"`
	ID       string  `json:"id"`
	Company  *string `json:"company"`
	Format   string  `json:"format"`
	Capacity string  `json:"capacity"`
	Speed    string  `json:"speed"`
	Price    float64 `json:"price"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Pcs      []*Pc  `json:"pcs"`
}
