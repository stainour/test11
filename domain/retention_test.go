package domain_test

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/stainour/test11/domain"
	"github.com/stainour/test11/domain/specificactions"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

var retainBeforeDate = time.Date(2019, 5, 31, 0, 0, 0, 0, time.UTC)

func TestOneLatestBackupOneRetained(t *testing.T) {
	backups := []domain.Backup{domain.NewBackup(time.Date(2019, 5, 30, 0, 0, 0, 0, time.UTC), 1)}
	retainedBackups := applySpecification(backups)
	assert.EqualValues(t, backups, retainedBackups)
}

func TestNoBackupNoRetained(t *testing.T) {
	backups := []domain.Backup{}
	retainedBackups := applySpecification(backups)
	assert.EqualValues(t, backups, retainedBackups)
}

func TestOneOldestBackupOneRetained(t *testing.T) {
	backups := []domain.Backup{domain.NewBackup(time.Date(2019, 5, 10, 0, 0, 0, 0, time.UTC), 1)}
	retainedBackups := applySpecification(backups)
	assert.EqualValues(t, backups, retainedBackups)
}

func TestOneBackupEveryDayTwelveRetained(t *testing.T) {
	assertions := assert.New(t)

	var backups []domain.Backup
	linq.Range(1, 31).SelectT(func(i int) domain.Backup {
		return domain.NewBackup(time.Date(2019, 5, i, 0, 0, 0, 0, time.UTC), 1)
	}).ToSlice(&backups)

	retainedBackups := applySpecification(backups)

	sort.Slice(backups, func(i, j int) bool {
		return backups[i].CreationDate().After(backups[j].CreationDate())
	})
	assertions.Len(retainedBackups, 12)

	assertions.EqualValues(retainedBackups[:11], backups[:11])
	assertions.Equal(retainedBackups[11], backups[14])
}

func TestTwoBackupsEveryDaySeventeenRetained(t *testing.T) {
	assertions := assert.New(t)

	var backups []domain.Backup
	linq.Range(2, 62).SelectT(func(i int) domain.Backup {
		return domain.NewBackup(time.Date(2019, 5, i/2, 5, i/2, 0, 0, time.UTC), 1)
	}).ToSlice(&backups)

	retainedBackups := applySpecification(backups)
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].CreationDate().After(backups[j].CreationDate())
	})

	assertions.Len(retainedBackups, 17)
	assertions.EqualValues(retainedBackups[:12], backups[:12])

	assertions.Equal(retainedBackups[12], backups[16])
	assertions.Equal(retainedBackups[13], backups[17])
	assertions.Equal(retainedBackups[14], backups[18])
	assertions.Equal(retainedBackups[15], backups[19])
	assertions.Equal(retainedBackups[16], backups[30])
}

func TestTwoOldestBackupOneRetained(t *testing.T) {
	assertions := assert.New(t)

	backups := []domain.Backup{
		domain.NewBackup(time.Date(2019, 5, 10, 0, 0, 0, 0, time.UTC), 1),
		domain.NewBackup(time.Date(2019, 5, 8, 0, 0, 0, 0, time.UTC), 2)}

	retainedBackups := applySpecification(backups)
	assertions.Len(retainedBackups, 1)
	assertions.Equal(retainedBackups[0], backups[0])
}

func applySpecification(backups []domain.Backup) []domain.Backup {
	specification := getSpecification()
	return specification.GetRetainedBackups(backups)
}

func getSpecification() specificactions.BackupRetentionSpecification {
	builder := specificactions.NewBackupRetentionSpecificationBuilder(retainBeforeDate)
	return builder.GetDefaultSpecification()
}
