package specificactions

import (
	"github.com/stainour/test11/domain"
	"time"
)

type innerBackupRetentionSpecification struct {
	NextSpecification backupRetentionSpecification
	abstractBackupRetentionSpecification
}

func (spec *innerBackupRetentionSpecification) ShouldBeRetained(backup domain.Backup) bool {
	creationDate := backup.CreationDate()
	thresholdDate := spec.ThresholdDate()
	nextThresholdDate := spec.NextSpecification.ThresholdDate()
	if (creationDate.Before(thresholdDate) || creationDate.Equal(thresholdDate)) && creationDate.After(nextThresholdDate) && !spec.IsSpecificationMaxOut() {
		spec.NewBackupFound()
		return true
	}
	return spec.NextSpecification.ShouldBeRetained(backup)
}

func NewInnerBackupRetentionSpecification(maxBackupAge uint, retainBeforeDate time.Time, maxRetainedCount uint) *innerBackupRetentionSpecification {
	return &innerBackupRetentionSpecification{
		abstractBackupRetentionSpecification: *NewAbstractBackupRetentionSpecification(maxBackupAge, retainBeforeDate, maxRetainedCount)}
}

func (spec *innerBackupRetentionSpecification) SetNextSpecification(nextSpecification backupRetentionSpecification) {
	spec.NextSpecification = nextSpecification
}

func (spec *innerBackupRetentionSpecification) AreAllRetainedBackupsFound() bool {
	return spec.NextSpecification.AreAllRetainedBackupsFound()
}
