package models

import "time"

type ImagesBt struct {
	ID      int64
	Title   string
	Content []byte
	Created time.Time
	TextID  int64
}

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

/*
type Text struct {
	bun.BaseModel `bun:"table:text"`

	ID         int64     `bun:",pk,autoincrement"`
	Title      string    `bun:",notnull"`
	Content    string    `bun:",notnull"`
	Created    time.Time `bun:",notnull,default:current_timestamp"`
	ImageBtID  int64     `bun:","`
	ImageRefID int64     `bun:","`
}

type ImagesRef struct {
	bun.BaseModel `bun:"table:imagesref"`

	ID      int64     `bun:",pk,autoincrement"`
	Title   string    `bun:",notnull"`
	Content string    `bun:",notnull"`
	Created time.Time `bun:",notnull,default:current_timestamp"`
	TextID  int64     `bun:",notnull"`
}

type ImagesBt struct {
	bun.BaseModel `bun:"table:imagesbt"`

	ID      int64     `bun:",pk,autoincrement"`
	Title   string    `bun:",notnull"`
	Content []byte    `bun:",notnull,type:bytea"`
	Created time.Time `bun:",notnull,default:current_timestamp"`
	TextID  int64     `bun:",notnull"`
}
*/
