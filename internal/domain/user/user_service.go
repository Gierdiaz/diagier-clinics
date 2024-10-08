package user

import (
	"context"
	"fmt"

	"github.com/Gierdiaz/diagier-clinics/pkg/middleware"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Authenticate(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, email, password string) error
}

type userService struct {
	repository UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repository: repo}
}

func (s *userService) Authenticate(ctx context.Context, email, password string) (string, error) {
	user, err := s.repository.Email(ctx, email)
	if err != nil || user == nil {
		fmt.Println("Erro ao buscar o usuário:", err)
		return "", err
	}

	fmt.Println("Usuário encontrado:", user.Email, " | Senha no banco:", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Erro na comparação de senha:", err)
		return "", err
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		fmt.Println("Erro ao gerar token:", err)
		return "", err
	}

	fmt.Println("Autenticação bem-sucedida, token gerado.")
	return token, nil
}

func (s *userService) Register(ctx context.Context, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erro ao criptografar senha:", err)
		return err
	}

	user := &User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
	}

	fmt.Println("Registrando usuário com email:", user.Email, " | Senha criptografada:", user.Password)

	err = s.repository.Create(ctx, user)
	if err != nil {
		fmt.Println("Erro ao criar usuário no banco de dados:", err)
		return err
	}

	fmt.Println("Usuário registrado com sucesso:", user.Email)
	return nil
}
