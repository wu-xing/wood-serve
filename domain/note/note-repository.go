package note

import "github.com/wu-xing/wood-serve/common"

type Repository struct{}

func (repository *Repository) saveNote(note Note) bool {
	db := common.GetDB()
	return db.Connection.NewRecord(note)
}
