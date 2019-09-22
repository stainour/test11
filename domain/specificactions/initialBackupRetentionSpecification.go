package specificactions

import (
	"github.com/stainour/test11/domain"
	"sort"
)

type initialBackupRetentionSpecification struct {
	currentSatisfied  int
	NextSpecification backupRetentionSpecification
}

func (spec *initialBackupRetentionSpecification) GetRetainedBackups(backups []domain.Backup) []domain.Backup {
	sortedBackups := make([]domain.Backup, len(backups))
	copy(sortedBackups, backups)
	sort.Slice(sortedBackups, func(i, j int) bool {
		return sortedBackups[i].CreationDate().After(sortedBackups[j].CreationDate())
	})

	retainedBackups := []domain.Backup{}

	for _, backup := range sortedBackups {
		if spec.ShouldBeRetained(backup) {
			retainedBackups = append(retainedBackups, backup)
		}
		if spec.NextSpecification.AreAllRetainedBackupsFound() {
			break
		}
	}
	return retainedBackups
}

func NewInitialBackupRetentionSpecification(nextSpecification backupRetentionSpecification) *initialBackupRetentionSpecification {
	return &initialBackupRetentionSpecification{NextSpecification: nextSpecification}
}

func (spec *initialBackupRetentionSpecification) ShouldBeRetained(backup domain.Backup) bool {
	if backup.CreationDate().After(spec.NextSpecification.ThresholdDate()) {
		spec.currentSatisfied++
		return true
	}
	return spec.NextSpecification.ShouldBeRetained(backup)
}
