package patients

import (
	"github.com/google/uuid"
	"time"
)

type PacientDTO struct {
	Name         string `json:"name" validate:"required,min=3,max=100"`
	Age          int    `json:"age" validate:"required,gt=0,lt=150"`
	Gender       string `json:"gender" validate:"required,oneof=male female other"`
	Address      string `json:"address" validate:"required,min=3,max=100"`
	Phone        string `json:"phone" validate:"required,e164"` // Exemplo: +5511999999999
	Email        string `json:"email" validate:"required,email"`
	Observations string `json:"observations" validate:"omitempty,min=3,max=1000"`
}

func (dto *PacientDTO) ToEntity() *Pacient {
	return &Pacient{
		ID:           uuid.New(),
		Name:         dto.Name,
		Age:          dto.Age,
		Gender:       dto.Gender,
		Address:      dto.Address,
		Phone:        dto.Phone,
		Email:        dto.Email,
		Observations: dto.Observations,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
