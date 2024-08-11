package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/cesart18/ranking_app/db"
	"github.com/cesart18/ranking_app/models"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type PlayerServiceInterface interface {
	CreatePlayer(*models.Player) (string, error)
	GetPlayers() ([]models.Player, error)
	GetPlayerById(int) (models.Player, error)
	UpdatePlayer(int, string) (string, error)
	AddDoublePoint(int, int) (string, error)
	AddSinglePoint(int, int) (string, error)
	AddMatch([]int, string) (string, error)
	ToggleMembership(int) (string, error)
	TogglePayedBalls(int) (string, error)
	DeletePlayer(int) (string, error)
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

	var players []models.Player

	r := db.DB.Preload("MatchPlayers").Preload("MatchPlayers.Match.MatchPlayers").Find(&players)
	if r.Error != nil {
		return []models.Player{}, r.Error
	}

	return players, nil

}

func (us *PlayerService) GetPlayerById(i int) (models.Player, error) {
	var user models.Player
	r := db.DB.Preload("MatchPlayers").Preload("MatchPlayers.Match.MatchPlayers").First(&user, i)

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

func (us *PlayerService) ToggleMembership(id int) (string, error) {
	var player models.Player
	r := db.DB.First(&player, id)
	if r.Error != nil {
		return "", r.Error
	}
	player.IsMembershipValid = !player.IsMembershipValid

	if r := db.DB.Save(&player); r.Error != nil {
		return "", r.Error
	}

	msg := fmt.Sprintf("Membresia cambiada a %v", player.IsMembershipValid)
	return msg, nil
}
func (us *PlayerService) TogglePayedBalls(id int) (string, error) {
	var player models.Player
	r := db.DB.First(&player, id)
	if r.Error != nil {
		return "", r.Error
	}
	player.IsPayedBalls = !player.IsPayedBalls

	if r := db.DB.Save(&player); r.Error != nil {
		return "", r.Error
	}
	msg := fmt.Sprintf("Pago de pelotas cambiada a %v", player.IsPayedBalls)
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

func (us *PlayerService) AddMatch(ids []int, score string) (string, error) {
	if len(ids) == 0 {
		return "", fmt.Errorf("debe proporcionar al menos dos jugadores")
	}
	if len(ids) < 2 || len(ids) > 4 {
		return "", fmt.Errorf("debe proporcionar los jugadores completos")
	}

	var players []models.Player
	if tx := db.DB.Find(&players, ids); tx.Error != nil {
		return "", tx.Error
	}

	var matchPlayers []models.MatchPlayer

	var winners []bool

	// Determinar ganadores según el número de IDs
	if len(ids) == 2 {
		winners = []bool{true, false} // El primer jugador gana, el segundo pierde
	} else if len(ids) == 4 {
		winners = []bool{true, true, false, false} // Los dos primeros ganan, los dos últimos pierden
	}

	for _, player := range players {
		index := IndexOf(ids, int(player.ID))
		matchPlayer := models.MatchPlayer{
			PlayerID: player.ID,
			Winner:   winners[index], // Los primeros jugadores son los ganadores
		}
		matchPlayers = append(matchPlayers, matchPlayer)

	}

	match := models.Match{
		Score:        score,
		Date:         formatDate(time.Now()),
		MatchType:    getMatchType(len(ids)),
		MatchPlayers: matchPlayers,
	}

	if tx := db.DB.Create(&match); tx.Error != nil {
		return "", tx.Error
	}
	for _, matchPlayer := range matchPlayers {
		if match.MatchType == "single" {
			if matchPlayer.Winner {
				us.AddSinglePoint(int(matchPlayer.PlayerID), 2)
			} else {
				us.AddSinglePoint(int(matchPlayer.PlayerID), -1)
			}
		} else {
			if matchPlayer.Winner {
				us.AddDoublePoint(int(matchPlayer.PlayerID), 2)
			} else {
				us.AddDoublePoint(int(matchPlayer.PlayerID), -1)
			}
		}
	}

	return "juego creado exitosamente", nil
}

func getMatchType(numPlayers int) string {
	if numPlayers == 2 {
		return "single"
	} else {
		return "double"
	}
}

func IndexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1 // Retorna -1 si no se encuentra el valor
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02") // Formato: YYYY-MM-DD
}
