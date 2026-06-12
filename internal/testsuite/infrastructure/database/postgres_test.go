package database_test

import (
	"testing"
	"time"

	"govent/internal/domain/types"
	"govent/internal/infrastructure/database"

	"govent/internal/testsuite/infrastructure/testdb"
	internal_test "govent/internal/testsuite/internal"

	"github.com/stretchr/testify/require"
)

func TestGoldenCreate(t *testing.T) {
	var golden = internal_test.NewRandomGolden()
	err := testdb.GetRepository().Create(golden)
	require.NoError(t, err)

	var result types.Golden
	err = testdb.GetDB().First(&result, "id = ?", golden.Id).Error
	require.NoError(t, err)
	require.Equal(t, golden.Id, result.Id)
	require.Equal(t, golden.Name, result.Name)
	require.WithinDuration(t, golden.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, golden.UpdatedAt, result.UpdatedAt, time.Second)

	testdb.GetDB().Delete(&result)
}

func TestGoldenDelete(t *testing.T) {
	var golden = internal_test.NewRandomGolden()
	_ = testdb.GetRepository().Create(golden)

	repository := database.NewGoldenPostgresRepository(testdb.GetDB())

	id, _ := types.NewGoldenId(golden.Id)
	err := repository.Delete(id)
	require.NoError(t, err)
}

func TestGoldenOne(t *testing.T) {
	var golden = internal_test.NewRandomGolden()
	_ = testdb.GetRepository().Create(golden)

	repository := database.NewGoldenPostgresRepository(testdb.GetDB())

	result, err := repository.One(golden.GetId())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, golden.Id, result.Id)
	require.Equal(t, golden.Name, result.Name)

	id, _ := types.NewGoldenId(golden.Id)
	err = repository.Delete(id)
	require.NoError(t, err)
}

func TestGoldenUpdate(t *testing.T) {
	var golden = internal_test.NewRandomGolden()
	_ = testdb.GetRepository().Create(golden)

	repository := database.NewGoldenPostgresRepository(testdb.GetDB())

	golden.Name = golden.Name + "Updated"
	err := repository.Update(golden)
	require.NoError(t, err)

	var result types.Golden
	err = testdb.GetDB().First(&result, "id = ?", golden.Id).Error
	require.NoError(t, err)
	require.Equal(t, golden.Name, result.Name)

	id, _ := types.NewGoldenId(golden.Id)
	err = repository.Delete(id)
	require.NoError(t, err)
}
