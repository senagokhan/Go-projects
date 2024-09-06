package services

import (
	"Project-Management-System/config"
	"Project-Management-System/models"
	_ "errors"
	_ "golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.UserModel) (models.UserModel, error) {
	db := config.GetDB()

	result, err := db.Exec("INSERT INTO users (username, email, password_hash, role) VALUES (?, ?, ?, ?)",
		user.Username, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		return models.UserModel{}, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return models.UserModel{}, err
	}
	user.UserId = int(userID)

	return user, nil
}

// GetUserByEmail retrieves a user by email.
func GetUserByEmail(email string) (models.UserModel, error) {
	db := config.GetDB()

	var user models.UserModel
	err := db.QueryRow("SELECT user_id, username, email, password_hash, role FROM users WHERE email = ?", email).
		Scan(&user.UserId, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}

func GetUserById(userId string) (models.UserModel, error) {
	db := config.GetDB()

	var user models.UserModel
	err := db.QueryRow("SELECT user_id, username, email, password_hash, role FROM users WHERE user_id = ?", userId).
		Scan(&user.UserId, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}

func UpdateUser(userId string, updatedUser models.UserModel) error {
	db := config.GetDB()

	query := "UPDATE users SET username = ?, email = ?, role = ?"
	args := []interface{}{updatedUser.Username, updatedUser.Email, updatedUser.Role}

	if updatedUser.PasswordHash != "" {
		query += ", password_hash = ?"
		args = append(args, updatedUser.PasswordHash)
	}

	query += " WHERE user_id = ?"
	args = append(args, userId)

	_, err := db.Exec(query, args...)
	return err
}

func DeleteUser(userId string) error {
	db := config.GetDB()

	_, err := db.Exec("DELETE FROM users WHERE user_id = ?", userId)
	return err
}
