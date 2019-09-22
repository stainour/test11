package main

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/stainour/test11/backupService"
	"github.com/stainour/test11/domain"
	"github.com/stainour/test11/domain/specificactions"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Minute)

	serviceFacade := backupService.NewInMemoryBackupServiceFacade()
	for _ = range ticker.C {
		now := time.Now()
		specification := specificactions.NewBackupRetentionSpecificationBuilder(now).GetDefaultSpecification()

		backups := serviceFacade.GetBackups(now)
		retainedBackups := specification.GetRetainedBackups(backups)

		var backupsToDelete []domain.Backup

		linq.From(backups).ExceptBy(linq.From(retainedBackups), func(i interface{}) interface{} {
			return i.(domain.Backup).Id()
		}).ToSlice(&backupsToDelete)

		serviceFacade.Delete(backupsToDelete)
	}
}
