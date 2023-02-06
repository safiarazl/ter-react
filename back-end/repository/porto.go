package repository

import (
	"api-cv/model"

	"gorm.io/gorm"
)

type PortoRepository struct {
	db *gorm.DB
}

func NewPortoRepository(db *gorm.DB) PortoRepository {
	return PortoRepository{db}
}

func (p *PortoRepository) AddPorto(porto model.Porto) error {
	return p.db.Create(&porto).Error //TODO: replace this
}

func (p *PortoRepository) ReadPorto() ([]model.Porto, error) {
	var porto []model.Porto
	err := p.db.Select("*").Where("deleted_at IS NULL").Find(&porto).Error // TODO: replace this
	return porto, err
}

func (p *PortoRepository) DeletePorto(id uint) error {
	return p.db.Where("id = ?", id).Delete(&model.Porto{}).Error //TODO: replace this
}

func (p *PortoRepository) UpdatePorto(id uint, porto model.Porto) error {
	return p.db.Model(&porto).Where("id = ?", id).Updates(&porto).Error //TODO: replace this
}