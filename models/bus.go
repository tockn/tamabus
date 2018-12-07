package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/tockn/tamabus/domain"
)

type Bus struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type CongestionLog struct {
	ID         int64   `db:"id"`
	Latitude   float64 `db:"latitude"`
	Longitude  float64 `db:"longitude"`
	Congestion int64   `db:"congestion"`
	BusID      int64   `db:"bus_id"`
}

func GetAll(db *sqlx.DB) ([]*domain.Bus, error) {
	var ids []int64
	if err := db.Select(&ids, `SELECT id FROM buses`); err != nil {
		return nil, err
	}
	var logs []*CongestionLog
	for id := range ids {
		var log *CongestionLog
		err := db.Get(&log, `
			SELECT
				id, latitude, longitude, congestion, bus_id
			FROM
				congestion_log
		  	WHERE
		  		bus_id = ?
		  	ORDER BY
		  		created_at
		  	DESC
		  	LIMIT 1`, id)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	bs := make([]*domain.Bus, len(logs))
	for i, c := range logs {
		p := calcPosition(c.Latitude, c.Longitude)
		b := domain.Bus{
			BusID:      c.BusID,
			Position:   p,
			Congestion: c.Congestion,
			Direction:  0,
			Latitude:   c.Latitude,
			Longitude:  c.Longitude,
		}
		bs[i] = &b
	}
	return bs, nil
}

func UpdatePosByID(id int64, bus Bus) (Bus, error) {
	// where id = ?して、緯度経度、posをupdate
	return *new(Bus), nil
}

func calcPosition(long float64, lati float64) int64 {

}
