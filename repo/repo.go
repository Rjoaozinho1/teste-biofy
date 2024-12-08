package repo

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	Nome      string    `json:"nome"`
	SessionID uuid.UUID `json:"session_id"`
	Mensagem  string    `json:"mensagem"`
	CreatedAt time.Time `json:"created_at"`
}

type ReqLogin struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type ResLogin struct {
	Token string `json:"token"`
}
