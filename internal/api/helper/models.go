package helper

import (
	"encoding/json"
	"time"

	"github.com/naijasonezeiru/go-phish-backend/internal/database"
)

type User struct {
	ID        int32           `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Username  string          `json:"username"`
	Jwt       string          `json:"jwt"`
	Victims   json.RawMessage `json:"victims"`
}

func DatabaseUserToUser(dbUser database.GetUserByUsernameRow, jwt string) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Username:  dbUser.Username,
		Victims:   dbUser.Victims,
		Jwt:       jwt,
	}
}
