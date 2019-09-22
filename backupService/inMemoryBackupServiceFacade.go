package backupService

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/stainour/test11/domain"
	"math/rand"
	"time"
)

type InMemoryBackupServiceFacade struct {
	backups map[int]domain.Backup
}

func (facade *InMemoryBackupServiceFacade) Delete(backupsToDelete []domain.Backup) {
	for _, backup := range backupsToDelete {
		delete(facade.backups, backup.Id())
	}
}

func (facade *InMemoryBackupServiceFacade) GetBackups(toDate time.Time) []domain.Backup {
	var backups []domain.Backup
	linq.From(facade.backups).SelectT(func(keyValue linq.KeyValue) domain.Backup {
		return keyValue.Value.(domain.Backup)
	}).WhereT(func(backup domain.Backup) bool {
		return backup.CreationDate().Before(toDate)
	}).ToSlice(&backups)
	return backups
}

func NewInMemoryBackupServiceFacade() *InMemoryBackupServiceFacade {
	facade := &InMemoryBackupServiceFacade{backups: map[int]domain.Backup{}}

	now := time.Now()
	linq.Range(1, 30).SelectT(func(i int) domain.Backup {
		return domain.NewBackup(now.Add(time.Duration(-i*12)*time.Hour), i)
	}).OrderByT(func(backup domain.Backup) int {
		return rand.Int()
	}).ForEachT(func(backup domain.Backup) {
		facade.backups[backup.Id()] = backup
	})

	return facade
}
