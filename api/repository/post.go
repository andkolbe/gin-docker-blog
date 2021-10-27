package repository

import (
	"github.com/andkolbe/gin-docker-blog/infrastructure"
	"github.com/andkolbe/gin-docker-blog/models"
)

// this layer is the one that interacts and performs CRUD operations on the database

type PostReposity struct {
	db infrastructure.Database
}

// NewPostReposity : fetching database
func NewPostRepository(db infrastructure.Database) PostReposity {
	return PostReposity{
		db: db,
	}
}

// Save -> method for saving post to database
func (p PostReposity) Save(post models.Post) error {
	return p.db.DB.Create(&post).Error
}

// FindAll -> method for fetching all posts from db
func (p PostReposity) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0

	queryBuilder := p.db.DB.Order("created_at desc").Model(&models.Post{})

	// search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			p.db.DB.Where("post.title LIKE ? ", queryKeyword),
		)
	}

	err := queryBuilder.Where(post).Find(&posts).Count(&totalRows).Error 

	return &posts, totalRows, err
}

// Update -> method for updating Post
func (p PostReposity) Update(post models.Post) error {
	return p.db.DB.Save(&post).Error
}

// Find -> method for fetching post by id
func (p PostReposity) Find(post models.Post) (models.Post, error) {
	var posts models.Post
	err := p.db.DB.Debug().Model(&models.Post{}).Where(&post).Take(&posts).Error

	return posts, err 
}

// Delete Post
func (p PostReposity) Delete(post models.Post) error {
	return p.db.DB.Delete(&post).Error
}
 