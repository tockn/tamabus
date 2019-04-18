package domain

type Bus struct {
	BusID      int64   `json:"bus_id"`
	Position   int64   `json:"position"`
	Congestion int64   `json:"congestion"`
	Direction  int64   `json:"direction"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
}

type BusImage struct {
	BusID    int64  `json:"bus_id"`
	Base64   string `json:"base64"`
	FileType string `json:"file_type"`
}
