package repository

import (
	"api-cv/model"

	"gorm.io/gorm"
)

type EducationRepository struct {
	db *gorm.DB
}

func NewEducationRepository(db *gorm.DB) EducationRepository {
	return EducationRepository{db}
}

func (p *EducationRepository) AddEducation(education model.Education) error {
	return p.db.Create(&education).Error //TODO: replace this
}

func (p *EducationRepository) ReadEducation() ([]model.Education, error) {
	var education []model.Education
	err := p.db.Select("*").Where("deleted_at IS NULL").Find(&education).Error // TODO: replace this
	return education, err
}

func (p *EducationRepository) DeleteEducation(id uint) error {
	return p.db.Where("id = ?", id).Delete(&model.Education{}).Error //TODO: replace this
}

func (p *EducationRepository) UpdateEducation(id uint, education model.Education) error {
	return p.db.Model(&education).Where("id = ?", id).Updates(&education).Error //TODO: replace this
}