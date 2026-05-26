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
	ImageRefID int64
}

type TeximgModel struct {
	DB *sql.DB
}

func (m *TeximgModel) InsertImg(title string, content string) (int, error) {
	return 0, nil
}

func (m *TeximgModel) InsertTex(title string, content string) (int, error) {
	// row := db.QueryRow(`insert into users(name, age) values('Scrooge McDuck', 93) returning id`)
	// var userid int
	// err := row.Scan(&userid)

	return 0, nil
}

func (m *TeximgModel) GetTex(id int) (Text, error) {
	sqlStmt := `SELECT id, title, content, created, image_ref_id FROM texts WHERE id = $1;`

	var t Text

	row := m.DB.QueryRow(sqlStmt, id)
	switch err := row.Scan(&t.ID, &t.Title, &t.Content, &t.Created, &t.ImageRefID); err {
	case sql.ErrNoRows:
		return Text{}, err
	case nil:
		return t, nil
	default:
		return Text{}, err
	}
}

func (m *TeximgModel) GetImg(id int) (ImagesRef, error) {
	return ImagesRef{}, nil
}

func (m *TeximgModel) Latest() ([]Text, error) {
	return nil, nil
}
