package specificactions

import (
	"github.com/stainour/test11/domain"
	"time"
)

type backupRetentionSpecification interface {
	AreAllRetainedBackupsFound() bool
	ShouldBeRetained(backup domain.Backup) bool
	ThresholdDate() time.Time
}

type BackupRetentionSpecification interface {
	GetRetainedBackups([]domain.Backup) []domain.Backup
}
