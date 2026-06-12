package database

import (
	"fmt"

	"govent/internal/domain/shared"
	"govent/internal/domain/types"

	"gorm.io/gorm"
)

type GoldenPostgresRepository struct {
	db *gorm.DB
}

func NewGoldenPostgresRepository(db *gorm.DB) GoldenPostgresRepository {
	return GoldenPostgresRepository{db: db}
}

func (r *GoldenPostgresRepository) Create(golden *types.Golden) error {
	return r.db.Create(golden).Error
}

func (r *GoldenPostgresRepository) Delete(id *types.GoldenId) error {
	_, err := r.One(id)
	if err != nil {
		return shared.ErrGoldenNotFound
	}

	return r.db.Delete(&types.Golden{}, "id = ?", id.Value()).Error
}

func (r *GoldenPostgresRepository) Update(golden *types.Golden) error {
	return r.db.Save(golden).Error
}

func (r *GoldenPostgresRepository) One(id *types.GoldenId) (*types.Golden, error) {
	var golden types.Golden
	if err := r.db.First(&golden, "id = ?", id.Value()).Error; err != nil {
		return nil, shared.ErrGoldenNotFound
	}
	return &golden, nil
}

func (r *GoldenPostgresRepository) All() ([]*types.Golden, error) {
	var goldens []*types.Golden
	if err := r.db.Find(&goldens).Error; err != nil {
		return nil, shared.ErrGoldenBadRequest
	}

	return goldens, nil
}

func (r *GoldenPostgresRepository) SearchAndPaginate(
	searchTerm string,
	pageNumber int,
	pageSize int) ([]*types.Golden, error) {
	offset := (pageNumber - 1) * pageSize
	var goldens []*types.Golden
	if err := r.db.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", searchTerm)).
		Order("name").
		Limit(pageSize).
		Offset(offset).
		Find(&goldens).Error; err != nil {
		return nil, err
	}

	return goldens, nil
}
