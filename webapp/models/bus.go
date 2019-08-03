package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tockn/tamabus/webapp/domain"
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
//	var ids []int64
//	if err := db.Select(&ids, `
//SELECT
//	id
//FROM
//	buses`); err != nil {
//		return nil, err
//	}
//	var logs []*CongestionLog
//	for _, id := range ids {
//		var lo []*CongestionLog
//		err := db.Select(&lo, `
//SELECT
//	id, latitude, longitude, congestion, bus_id
//FROM
//	congestion_log
//WHERE
//	bus_id = ? AND complete = 1
//ORDER BY
//	created_at
//DESC
//	LIMIT 1`, id)
//		if err != nil {
//			return nil, err
//		}
//		if len(lo) == 0 {
//			continue
//		}
//		logs = append(logs, lo[0])
//	}

	//bs := make([]*domain.Bus, len(logs))
	//for i, c := range logs {
	//	p := calcPosition(c.Latitude, c.Longitude)
	//	b := domain.Bus{
	//		BusID:      c.BusID,
	//		Position:   p,
	//		Congestion: c.Congestion,
	//		Direction:  0,
	//		Latitude:   c.Latitude,
	//		Longitude:  c.Longitude,
	//	}
	//	bs[i] = &b
	//}

	bs := make([]*domain.Bus, 2)
	for i, b := range bs {
		rand.Seed(time.Now().UnixNano())
		b = domain.Bus{
			BusID: i,
			Position: rand.Intn(6)
			Congestion: rand.Intn(5),
			Direction: 0,
			Latitude: 0,
			Longitude: 0,
		}
	}
	return bs, nil
}

func InsertLog(db *sqlx.DB, bus *domain.Bus) (*CongestionLog, error) {
	pos := calcPosition(bus.Longitude, bus.Latitude)

	// where id = ?して、緯度経度、posをupdate
	res, err := db.Exec(`
INSERT INTO 
	congestion_log(
		latitude, 
		longitude, 
		congestion, 
		position, 
		bus_id
	) 
VALUES
	(?, ?, ?, ?, ?)`,
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
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(6))
}

type BusImage struct {
	BusID  int64
	Path   string
	Base64 string
}

func (b *BusImage) Insert(db *sqlx.DB) error {
	_, err := db.Exec(`
INSERT INTO
	images(
		base64, bus_id
	)
VALUES
	(?, ?)`,
		b.Base64, b.BusID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllBusImages(db *sqlx.DB) ([]*domain.BusImage, error) {
	var ids []int64
	if err := db.Select(&ids, `
SELECT
	id
FROM
	buses`); err != nil {
		return nil, err
	}

	bis := make([]*BusImage, 0, len(ids))
	for _, id := range ids {
		bi := &BusImage{}
		if err := db.DB.QueryRow(`
SELECT
	bus_id, base64
FROM 
	images
WHERE
	bus_id = ?
ORDER BY
	created_at
	DESC
LIMIT 1;
`, id).Scan(&bi.BusID, &bi.Base64); err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, err
		}
		fmt.Println(bi)
		bis = append(bis, bi)
	}

	dbis := make([]*domain.BusImage, 0, len(bis))
	for _, b := range bis {
		dbi := &domain.BusImage{
			BusID:  b.BusID,
			Base64: b.Base64,
		}
		dbis = append(dbis, dbi)
	}
	return dbis, nil
}

func (c *CongestionLog) UpdateCongestion(db *sqlx.DB) error {
	if _, err := db.Exec(`
update
	congestion_log
set
	congestion = ?, 
	complete = 1 
where 
	bus_id = ? 
  and 
  	complete = 0`, c.Congestion, c.BusID); err != nil {
		return err
	}
	return nil
}

func TruncateImage(db *sqlx.DB) {
}
