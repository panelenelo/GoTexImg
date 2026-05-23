package models

import (
	"database/sql"
	"time"
)

type ImagesRef struct {
	ID      int64
	Title   string
	Content string
	Created time.Time
	TextID  int64
}

type Text struct {
	ID         int64
	Title      string
	Content    string
	Created    time.Time
	ImageBtID  int64
	ImageRefID int64
}

type TeximgModel struct {
	DB *sql.DB
}
