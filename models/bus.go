package models

type Bus struct {
	BusID      int64   `json:"bus_id"`
	Position   int64   `json:"position"`
	Congestion int64   `json:"congestion"`
	Direction  int64   `json:"direction"`
	Longitude  float32 `json:"longitude"`
	Latitude   float32 `json:"latitude"`
}

func GetAll() ([]Bus, error) {
	var buses []Bus
	for i := 0; i < 5; i++ {
		bus := Bus{
			BusID:      int64(i),
			Position:   int64(i),
			Congestion: int64(i),
			Direction:  int64(i),
		}
		buses = append(buses, bus)
	}
	return buses, nil
}

func UpdatePosByID(id int64, bus Bus) (Bus, error) {
	// where id = ?して、緯度経度、posをupdate
	return *new(Bus), nil
}
