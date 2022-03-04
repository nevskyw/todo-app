package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nevskyw/todo-app"
	"github.com/nevskyw/todo-app/pkg/repository"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH" // набор случайных байтов для JWT подписей
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandartClaims
	UserId int `json:"user_id"`
}

type AuthService struct { // AuthService - структура, которую в конструкторе будем принимать repository для работы с базой
	repo repository.Authorization
}

// NewAuthService...
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser...
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	// перед записью пользователей в БД мы будем хешировать пароль
	user.Password = generatePasswordHash(user.Password)
	return repo.CreateUser(user)
}

// generatePasswordHash...
// функция хеширования пороля!
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

// GenerateToken...
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// GetUser...
	// Получаем пользователя из БД
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	// NewWithClaims...
	// Генерируем токен, если такой пользователь существует
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{ // передаем метод для подписи (jwt.SigningMethodHS256) и  JSON объект с набором различных полей (&tokenClaims)
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

// ParseToken...
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	// ParseWithClaims... - функция, принимающая JWT токен
	// token.Method... - функция, которая возвращает ключ-подпись или ошибку
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
