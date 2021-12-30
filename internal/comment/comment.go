package comment

import "gorm.io/gorm"

// Service - struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - define our comment struct
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

//Commenter - the interface for our comment service
type Commenter interface {
	GetComment(ID int) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID int, newComment Comment) (Comment, error)
	DeleteComment(ID int) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

// GetComment - retrieves comments by ID from the database
func (s *Service) GetComment(ID int) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetCommentsBySlug - retrieves comments by slug from the database
func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment

	if result := s.DB.Where("slug = ?", slug).Find(&comments); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

// PostComment - inserts a comment into the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Create(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment - updates a comment by ID in the database
func (s *Service) UpdateComment(ID int, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)

	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// DeleteComment - deletes a comment by ID from the database
func (s *Service) DeleteComment(ID int) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments - gets all comments from the database
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}
