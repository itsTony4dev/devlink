package repository

import (
	"devlink/internal/models"

	"gorm.io/gorm"
)

type ResourceRepository struct {
	db *gorm.DB
}

func NewResourceRepository(db *gorm.DB) *ResourceRepository {
	return &ResourceRepository{db: db}
}

func (r *ResourceRepository) GetByID(resourceID uint) (*models.Resource, error) {
	var resource models.Resource
	if err := r.db.First(&resource, resourceID).Error; err != nil {
		return nil, err
	}
	return &resource, nil
}

func (r *ResourceRepository) GetByUserID(userID uint, page, pageSize int) ([]models.Resource, int64, error) {
	var resources []models.Resource
	var total int64

	// Get total count
	if err := r.db.Model(&models.Resource{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * pageSize
	if err := r.db.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Find(&resources).Error; err != nil {
		return nil, 0, err
	}

	return resources, total, nil
}

func (r *ResourceRepository) CreateResource(resource *models.Resource) error {
	return r.db.Create(resource).Error
}

func (r *ResourceRepository) UpdateResource(resource *models.Resource) error {
	return r.db.Save(resource).Error
}

func (r *ResourceRepository) DeleteResource(resourceID uint) error {
	return r.db.Delete(&models.Resource{}, resourceID).Error
}

func (r *ResourceRepository) SearchResources(query string, userID uint, page, pageSize int) ([]models.Resource, int64, error) {
	var resources []models.Resource
	var total int64

	// Build search query
	searchQuery := r.db.Model(&models.Resource{}).
		Where("user_id = ? AND (title LIKE ? OR description LIKE ? OR url LIKE ?)",
			userID,
			"%"+query+"%",
			"%"+query+"%",
			"%"+query+"%")

	// Get total count
	if err := searchQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * pageSize
	if err := searchQuery.Offset(offset).Limit(pageSize).Find(&resources).Error; err != nil {
		return nil, 0, err
	}

	return resources, total, nil
}

func (r *ResourceRepository) GetByTags(tags []string, userID uint, page, pageSize int) ([]models.Resource, int64, error) {
	var resources []models.Resource
	var total int64

	// Build tag search query
	query := r.db.Model(&models.Resource{}).Where("user_id = ?", userID)
	for _, tag := range tags {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&resources).Error; err != nil {
		return nil, 0, err
	}

	return resources, total, nil
} 