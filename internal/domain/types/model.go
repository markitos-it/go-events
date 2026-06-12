package types

import (
	"log"
	"time"
)

type Golden struct {
	Id        string    `json:"id" binding:"required,uuid"`
	Name      string    `json:"name" binding:"required"`
	Content   string    `json:"content" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" binding:"required,datetime" default:"now"`
	UpdatedAt time.Time `json:"updated_at" binding:"required,datetime" default:"now"`
}

func NewGolden(id, name, content string) (*Golden, error) {
	secureId, err := NewGoldenId(id)

	if err != nil {
		log.Printf("❌ DEBUG ERROR (NewGoldenId): %v\n", err)
		return nil, err
	}

	secureName, err := NewGoldenName(name)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (NewGoldenName): %v\n", err)
		return nil, err
	}

	secureContent, err := NewGoldenContent(content)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (NewGoldenContent): %v\n", err)
		return nil, err
	}

	return &Golden{
		Id:        secureId.Value(),
		Name:      secureName.Value(),
		Content:   secureContent.Value(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (h *Golden) GetId() *GoldenId {
	result, _ := NewGoldenId(h.Id)
	return result
}

func (h *Golden) GetContent() *GoldenContent {
	result, _ := NewGoldenContent(h.Content)
	return result
}
