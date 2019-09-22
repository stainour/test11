package backupService

import (
	"github.com/stainour/test11/domain"
	"time"
)

type Facade interface {
	Delete(backupsToDelete []domain.Backup)
	GetBackups(toDate time.Time) []domain.Backup
}
