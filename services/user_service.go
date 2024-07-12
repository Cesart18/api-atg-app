package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/cesart18/ranking_app/config"
	"github.com/cesart18/ranking_app/db"
	"github.com/cesart18/ranking_app/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Signup(string, string) (string, error)
	Login(string, string) (string, error)
	Logout(models.RevokedToken) (string, error)
}

type UserService struct{}

func (us *UserService) Signup(username, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	user := models.User{Username: username, Password: string(hash)}
	r := db.DB.Create(&user)

	if r.Error != nil {
		var pgError *pgconn.PgError
		if errors.As(r.Error, &pgError) && pgError.Code == "23505" {
			return "", fmt.Errorf("usuario con el nombre %s ya existe", user.Username)
		}
		return "", r.Error
	}
	return "Usuario creado exitosamente", nil
}

func (us *UserService) Login(username, password string) (string, error) {

	var user models.User
	db.DB.First(&user, "username = ?", username)
	if user.ID == 0 {
		return "", fmt.Errorf("usuario o contraseña invalidos")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.Secret))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (us *UserService) Logout(token models.RevokedToken) (string, error) {

	r := db.DB.Create(&token)

	if r.Error != nil {
		return "", r.Error
	}
	return "Sesion cerrada correctamente", nil
}
