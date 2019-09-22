package domain

import "time"

type Backup struct {
	creationDate time.Time
	id           int
}

func NewBackup(creationDate time.Time, id int) Backup {
	return Backup{creationDate: creationDate, id: id}
}

func (b Backup) CreationDate() time.Time {
	return b.creationDate
}

func (b Backup) Id() int {
	return b.id
}
