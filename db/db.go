package db

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
	_, exists := db.GetByUserIdAndGameId(f.UserId, f.GameId)
	if !exists {
		db.f = append(db.f, f)
	}
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

func (db *DB) GetUserMedalByUserId(uId string) (int) {
	medal := 0
	userGame := make(map[string][]Field)

	for _, f := range db.f {
		if f.UserId == uId && f.Score > 10 {
			userGame[f.Level] = append(userGame[f.Level], f)
		}
	}

	for _, f := range userGame {
		if len(f) >= 3 {
			medal++
		}
	}

	return medal
}

func (db *DB) GetByUserIdAndGameId(uId string, gId string) (Field, bool) {
	for _, f := range db.f {
		if f.UserId == uId && f.GameId == gId {
			return f, true
		}
	}
	return Field{}, false
}
