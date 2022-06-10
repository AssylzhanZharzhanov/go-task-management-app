package user

import "gorm.io/gorm"

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository() *PostgresRepository {
	return &PostgresRepository{}
}
