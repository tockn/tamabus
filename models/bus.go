package models

type Bus struct {
	BusID      int `json:"bus_id"`
	Position   int `json:"position"`
	Congestion int `json:"congestion"`
	Direction  int `json:"direction"`
}

func GetAll() ([]Bus, error) {
	var buses []Bus
	for i := 0; i < 5; i++ {
		bus := Bus{
			BusID:      i,
			Position:   i,
			Congestion: i,
			Direction:  i,
		}
		buses = append(buses, bus)
	}
	return buses, nil
}
