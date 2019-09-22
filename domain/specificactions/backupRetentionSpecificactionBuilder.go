package specificactions

import (
	"errors"
	"sort"
	"time"
)

type setting struct {
	ageInDays        uint
	maxRetainedCount uint
}

type BackupRetentionSpecificationBuilder struct {
	retainBeforeDate time.Time
	specSettings     []setting
}

func NewBackupRetentionSpecificationBuilder(retainBeforeDate time.Time) *BackupRetentionSpecificationBuilder {
	return &BackupRetentionSpecificationBuilder{retainBeforeDate: retainBeforeDate}
}

func (builder *BackupRetentionSpecificationBuilder) Build() (BackupRetentionSpecification, error) {
	length := len(builder.specSettings)
	if length > 0 {

		if length == 1 {
			spec := builder.specSettings[0]
			return NewInitialBackupRetentionSpecification(NewFinishingBackupRetentionSpecification(spec.ageInDays, builder.retainBeforeDate, spec.maxRetainedCount)), nil
		}

		sort.SliceStable(builder.specSettings, func(i, j int) bool {
			return builder.specSettings[i].ageInDays < builder.specSettings[i].ageInDays
		})

		var prevSpecification *innerBackupRetentionSpecification = nil
		var firstSpecification *innerBackupRetentionSpecification = nil

		for i := 0; i < length-1; i++ {
			specSetting := builder.specSettings[i]
			specification := NewInnerBackupRetentionSpecification(specSetting.ageInDays, builder.retainBeforeDate, specSetting.maxRetainedCount)

			if prevSpecification != nil {
				prevSpecification.SetNextSpecification(specification)
			}

			prevSpecification = specification

			if firstSpecification == nil {
				firstSpecification = prevSpecification
			}

		}

		lastSpec := builder.specSettings[length-1]

		prevSpecification.SetNextSpecification(NewFinishingBackupRetentionSpecification(lastSpec.ageInDays, builder.retainBeforeDate, lastSpec.maxRetainedCount))
		return NewInitialBackupRetentionSpecification(firstSpecification), nil
	}

	return nil, errors.New("AddRule() should be invoked at least once")
}

func (builder *BackupRetentionSpecificationBuilder) GetDefaultSpecification() BackupRetentionSpecification {
	builder.AddRule(3, 4)
	builder.AddRule(7, 4)
	builder.AddRule(14, 1)
	specification, _ := builder.Build()
	return specification
}

func (builder *BackupRetentionSpecificationBuilder) AddRule(ageInDays uint, maxRetainedCount uint) {
	builder.specSettings = append(builder.specSettings, setting{
		ageInDays:        ageInDays,
		maxRetainedCount: maxRetainedCount,
	})
}
