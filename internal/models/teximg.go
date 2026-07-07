package models

import (
	"database/sql"
	"log"
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
	sqlStmt := `
	INSERT INTO texts (title, content, created)
	VALUES ($1, $2, NOW())
	RETURNING id`

	var id int
	err := m.DB.QueryRow(sqlStmt, title, content).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TeximgModel) GetTex(id int) (Text, error) {
	sqlStmt := `SELECT id, title, content, created, image_ref_id FROM texts WHERE id = $1;`

	var t Text
	var ptr *int64

	row := m.DB.QueryRow(sqlStmt, id)
	switch err := row.Scan(&t.ID, &t.Title, &t.Content, &t.Created, &ptr); err {
	case sql.ErrNoRows:
		return Text{}, err
	case nil:
		if ptr != nil {
			t.ImageRefID = *ptr
		}
		return t, nil
	default:
		return Text{}, err
	}
}

func (m *TeximgModel) GetImg(id int) (ImagesRef, error) {
	sqlStmt := `SELECT id, title, content, created, text_id FROM images_ref WHERE id = $1;`

	var i ImagesRef
	var ptr *int64

	row := m.DB.QueryRow(sqlStmt, id)
	switch err := row.Scan(&i.ID, &i.Title, &i.Content, &i.Created, &ptr); err {
	case sql.ErrNoRows:
		return ImagesRef{}, err
	case nil:
		return i, nil
	default:
		return ImagesRef{}, err
	}
}

func (m *TeximgModel) Latest() ([]Text, error) {

	sqlStmt := `SELECT id, title, content, created FROM texts
		ORDER BY id DESC LIMIT 10`

	texts := make([]Text, 0, 10)

	rows, err := m.DB.Query(sqlStmt)
	if err != nil {
		log.Printf("Error querying DB in Latest Function: %s", err.Error())
		return texts, err
	}

	defer rows.Close()

	for rows.Next() {
		var t Text
		err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.Created)
		if err != nil {
			log.Printf("Error scanning DB retrieval in \"Latest\" Function: %s", err.Error())
			return texts, err
		}

		texts = append(texts, t)
	}

	return texts, err
}
