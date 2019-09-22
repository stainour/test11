package specificactions

import (
	"time"
)

type abstractBackupRetentionSpecification struct {
	maxBackupAge         uint
	retainBeforeDate     time.Time
	thresholdDate        time.Time
	maxRetainedCount     uint
	currentRetainedCount uint
}

func (spec *abstractBackupRetentionSpecification) ThresholdDate() time.Time {
	return spec.thresholdDate
}

func (spec *abstractBackupRetentionSpecification) NewBackupFound() {
	spec.currentRetainedCount++
}

func (spec *abstractBackupRetentionSpecification) IsSpecificationMaxOut() bool {
	return spec.currentRetainedCount == spec.maxRetainedCount
}

func NewAbstractBackupRetentionSpecification(maxBackupAge uint, retainBeforeDate time.Time, maxRetainedCount uint) *abstractBackupRetentionSpecification {
	return &abstractBackupRetentionSpecification{
		maxBackupAge:     maxBackupAge,
		retainBeforeDate: retainBeforeDate,
		maxRetainedCount: maxRetainedCount,
		thresholdDate:    retainBeforeDate.AddDate(0, 0, int(-maxBackupAge)),
	}
}
