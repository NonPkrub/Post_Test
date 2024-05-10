package repository

import (
	"Test/domain"
	"fmt"

	"gorm.io/gorm"
)

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(DB *gorm.DB) *postRepository {
	return &postRepository{DB: DB}
}

func (repo *postRepository) Create(post *domain.Post) (*domain.Post, error) {
	tx := repo.DB.Create(post)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return post, nil
}

func (repo *postRepository) FindAllField(query *domain.PostAllReq, pagination *domain.Pagination) ([]*domain.Post, int64, int64, error) {
	posts := []*domain.Post{}
	var totalCount int64
	database := repo.DB
	offset := (pagination.Page - 1) * pagination.PageSize
	err := repo.query(database, query).Offset(offset).Limit(pagination.PageSize).Count(&totalCount).Find(&posts).Error
	if err != nil {
		return nil, 0, 0, err
	}

	totalPage := (totalCount + int64(pagination.PageSize) - 1) / int64(pagination.PageSize)

	return posts, totalCount, totalPage, nil
}

func (repo *postRepository) query(database *gorm.DB, query *domain.PostAllReq) *gorm.DB {
	fmt.Println(database, query)
	if query.Title != "" {
		database = database.Where("title=?", query.Title)
	}

	if !query.CreatedAt.IsZero() {
		database = database.Where("created_at=?", query.CreatedAt)
	}

	if !query.Published {
		database = database.Where("published=?", query.Published)
	}

	return database
}

func (repo *postRepository) FindOne(post *domain.Post) (*domain.Post, error) {
	tx := repo.DB.Where("id=?", post.ID).Find(post)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if post.Published {
		post.ViewCount += 1
		tx := repo.DB.Save(post)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return post, nil
}

func (repo *postRepository) UpdateByID(post *domain.Post) (*domain.Post, error) {
	tx := repo.DB.Model(&domain.Post{}).Where("id = ?", post.ID).Updates(post)
	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = repo.DB.Where("id=?", post.ID).Find(post)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return post, nil
}

func (repo *postRepository) DeleteByID(post *domain.Post) error {
	tx := repo.DB.Where("id =?", post.ID).Delete(post)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
