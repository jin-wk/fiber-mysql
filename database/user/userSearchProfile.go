package database

import (
	"github.com/jin-wk/fiber-mysql/database"
	"github.com/jin-wk/fiber-mysql/models"
)

// SearchProfile : search user profile by id
func SearchProfile(ID string) (models.User, error) {
	var user models.User
	err := database.DB.QueryRow("SELECT * FROM users WHERE id = ?", ID).Scan(
		&user.ID,
		&user.UserID,
		&user.Password,
		&user.Name,
		&user.Phone,
		&user.Address,
		&user.BackNumber,
		&user.PreferredPosition,
	)
	if err != nil {
		return user, err
	}
	// defer rows.Close()
	// user := models.User{}
	// for rows.Next() {
	// 	if err := rows.Scan(
	// 		&user.ID,
	// 		&user.UserID,
	// 		&user.Password,
	// 		&user.Name,
	// 		&user.Phone,
	// 		&user.Address,
	// 		&user.BackNumber,
	// 		&user.PreferredPosition,
	// 	); err != nil {
	// 		return user, err
	// 	}
	// }
	return user, nil
}
