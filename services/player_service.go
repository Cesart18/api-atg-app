package services

import (
	"errors"
	"fmt"

	"github.com/cesart18/ranking_app/db"
	"github.com/cesart18/ranking_app/models"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type PlayerServiceInterface interface {
	CreatePlayer(u *models.Player) (string, error)
	GetPlayers() ([]models.Player, error)
	GetPlayerById(id int) (models.Player, error)
	UpdatePlayer(id int, nName string) (string, error)
	AddDoublePoint(id int, p int) (string, error)
	AddSinglePoint(id int, p int) (string, error)
	DeletePlayer(id int) (string, error)
}

type PlayerService struct{}

func (us *PlayerService) CreatePlayer(u *models.Player) (string, error) {
	r := db.DB.Create(&u)
	if r.Error != nil {
		var pgError *pgconn.PgError
		if errors.As(r.Error, &pgError) && pgError.Code == "23505" {
			return "", fmt.Errorf("jugador con el nombre %s ya existe", u.Name)
		}
		return "", r.Error
	}
	return "Usuario creado exitosamente", nil
}

func (us *PlayerService) GetPlayers() ([]models.Player, error) {

	var users []models.Player
	r := db.DB.Find(&users)

	if r.Error != nil {
		return []models.Player{}, r.Error
	}

	return users, nil

}

func (us *PlayerService) GetPlayerById(i int) (models.Player, error) {
	var user models.Player
	r := db.DB.First(&user, i)

	if r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return models.Player{}, fmt.Errorf("usuario con el id: %d no existe", i)
		}
		return models.Player{}, r.Error
	}
	return user, nil
}

func (us *PlayerService) UpdatePlayer(id int, newName string) (string, error) {
	var u models.Player
	r := db.DB.First(&u, id)
	oldName := u.Name
	if r.Error != nil {
		return "", r.Error
	}
	u.Name = newName

	if r := db.DB.Save(&u); r.Error != nil {
		return "", r.Error
	}
	msg := fmt.Sprintf("Nombre cambiado de %s a %s", oldName, newName)
	return msg, nil
}

func (us *PlayerService) AddDoublePoint(id int, p int) (string, error) {
	var user models.Player
	if r := db.DB.First(&user, id); r.Error != nil {
		return "", r.Error
	}

	if p < 0 {
		if (user.DoublePoints + p) < 0 {
			return "", fmt.Errorf("puntos del usuario %d, los puntos del usuario no pueden ser menor a 0", user.DoublePoints)
		}
	}

	user.DoublePoints += p

	if r := db.DB.Save(&user); r.Error != nil {
		return "", r.Error
	}

	msg := fmt.Sprintf("Puntos totales %d", user.DoublePoints)
	return msg, nil
}

func (us *PlayerService) AddSinglePoint(id int, p int) (string, error) {
	var user models.Player
	if r := db.DB.First(&user, id); r.Error != nil {
		return "", r.Error
	}

	if p < 0 {
		if (user.SinglePoints + p) < 0 {
			return "", fmt.Errorf("puntos del usuario %d, los puntos del usuario no pueden ser menor a 0", user.DoublePoints)
		}
	}

	user.SinglePoints += p

	if r := db.DB.Save(&user); r.Error != nil {
		return "", r.Error
	}

	msg := fmt.Sprintf("Puntos totales %d", user.SinglePoints)
	return msg, nil
}

func (us *PlayerService) DeletePlayer(i int) (string, error) {
	var u models.Player
	r := db.DB.Delete(&u, i)
	if r.Error != nil {
		return "", r.Error
	}
	return "Usuario eliminado exitosamente", nil
}
