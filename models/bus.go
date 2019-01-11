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
	for _, id := range ids {
		var lo []*CongestionLog
		err := db.Select(&lo, `
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
		if len(lo) == 0 {
			continue
		}
		logs = append(logs, lo[0])
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

func InsertLog(db *sqlx.DB, bus *domain.Bus) (*CongestionLog, error) {
	pos := calcPosition(bus.Longitude, bus.Latitude)

	// where id = ?して、緯度経度、posをupdate
	res, err := db.Exec(`INSERT INTO congestion_log(latitude, longitude, congestion, position, bus_id) VALUES (?, ?, ?, ?, ?)`,
		bus.Latitude, bus.Longitude, bus.Congestion, pos, bus.BusID)
	if err != nil {
		return nil, err
	}
	newID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	c := CongestionLog{
		ID:         newID,
		Latitude:   bus.Latitude,
		Longitude:  bus.Longitude,
		Congestion: bus.Congestion,
		BusID:      bus.BusID,
	}
	return &c, nil
}

func calcPosition(long float64, lati float64) int64 {
	return 1
}

type BusImage struct {
	BusID int64
	Body  string
}

func (b *BusImage) Insert(db *sqlx.DB) error {
	_, err := db.Exec(`INSERT INTO images(body, bus_id) VALUES (?, ?)`,
		b.Body, b.BusID)
	if err != nil {
		return err
	}
	return nil
}

func TruncateImage(db *sqlx.DB) (err error) {
	_, err := db.Exec(`TRUNCATE TABLE images`)
	return
}
