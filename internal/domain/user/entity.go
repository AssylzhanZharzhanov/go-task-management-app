package domain

const (
	UsersTableName = "users"
)

type UserID int64

type User struct {
	ID        UserID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	FirstName string `json:"first_name" gorm:"not null;column:first_name"`
	LastName  string `json:"last_name" gorm:"not null;column:last_name"`
	Email     string `json:"email" gorm:"not null;column:email"`
	Password  string `json:"password" gorm:"not null;column:password"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at"`
}
