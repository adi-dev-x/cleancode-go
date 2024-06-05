package user

import (
	"context"
	"log"
	"myproject/pkg/model"

	"gorm.io/gorm"
)

type Repository interface {
	Register(ctx context.Context, request UserRegisterRequest) error
	Listing(ctx context.Context) ([]model.UserRegisterRequest, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Register(ctx context.Context, request UserRegisterRequest) error {
	query := `INSERT INTO users (FirstName, LastName, Email, Password) VALUES( $1, $2, $3, $4)`

	// result := r.db.Raw(query, map[string]interface{}{
	// 	"first_name": request.FirstName,
	// 	"last_name":  request.LastName,
	// 	"email":      request.Email,
	// 	"password":   request.Password,
	// })
	result := r.db.Exec(query, request.FirstName, request.LastName, request.Email, request.Password)
	log.Println("this is the result ", result)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
func (r *repository) Listing(ctx context.Context) ([]model.UserRegisterRequest, error) {
	query := `SELECT * FROM users;`

	var products []model.UserRegisterRequest
	result := r.db.Raw(query).Scan(&products)
	if err := result.Error; err != nil {
		return nil, err
	}

	return products, nil
}
