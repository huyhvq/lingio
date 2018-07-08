package db

import (
	"sort"
	"fmt"
)

type DB struct {
	f []Field
}

type Field struct {
	GameId string
	UserId string
	Level  string
	Score  int
}

func NewDB() (DB) {
	return DB{}
}

func (db *DB) Set(f Field) {
	_, isExisted := db.GetByUserIdAndGameId(f.UserId, f.GameId)
	fmt.Printf("%v\n", isExisted)
	if !isExisted {
		db.f = append(db.f, f)
	}
	fmt.Printf("%v\n", db.f)
}

func (db *DB) GetUserScoresByUserId(uId string) (int) {
	total := 0
	for _, f := range db.f {
		if f.UserId == uId {
			total += f.Score
		}
	}
	return total
}

func (db *DB) GetByUserIdAndGameId(uId string, gId string) (Field, bool) {
	sort.Slice(db.f, func(i, j int) bool {
		return db.f[i].UserId <= db.f[i].UserId
	})
	idx := sort.Search(len(db.f), func(i int) bool {
		return db.f[i].UserId == uId && db.f[i].GameId == gId
	})
	if idx != 0 {
		return db.f[idx], true
	}
	return Field{}, false
}
