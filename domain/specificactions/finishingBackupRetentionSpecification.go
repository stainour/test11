package specificactions

import (
	"github.com/stainour/test11/domain"
	"time"
)

type finishingBackupRetentionSpecification struct {
	abstractBackupRetentionSpecification
}

func NewFinishingBackupRetentionSpecification(maxBackupAge uint, retainBeforeDate time.Time, maxRetainedCount uint) backupRetentionSpecification {
	return &finishingBackupRetentionSpecification{
		abstractBackupRetentionSpecification: *NewAbstractBackupRetentionSpecification(maxBackupAge, retainBeforeDate, maxRetainedCount),
	}
}

func (spec *finishingBackupRetentionSpecification) AreAllRetainedBackupsFound() bool {
	return spec.IsSpecificationMaxOut()
}

func (spec *finishingBackupRetentionSpecification) ShouldBeRetained(backup domain.Backup) bool {
	if spec.AreAllRetainedBackupsFound() {
		return false
	}

	if backup.CreationDate().Before(spec.ThresholdDate()) || backup.CreationDate().Equal(spec.ThresholdDate()) {
		spec.NewBackupFound()
		return true
	}
	return false
}
