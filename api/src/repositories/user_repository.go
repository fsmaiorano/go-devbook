package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
	"time"
)

// user repository
type users struct {
	db *sql.DB // inject db inside user struct
}

// Creates a new user repository
func NewRepositoryOfUsers(db *sql.DB) *users {
	return &users{db}
}

// Create a new user
func (userRepository users) Create(user models.User) (uint64, error) {
	statement, err := userRepository.db.Prepare("INSERT INTO users (name, nickname, email, password) VALUES (@name, @nickname, @email, @password)")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	user.CreatedAt = time.Now()

	_, err = statement.Exec(sql.Named("name", user.Name), sql.Named("nickname", user.Nickname), sql.Named("email", user.Email), sql.Named("password", user.Password))
	if err != nil {
		return 0, err
	}

	line, err := userRepository.db.Query("SELECT id FROM users WHERE email = @email", sql.Named("email", user.Email))
	if err != nil {
		return 0, err
	}

	if line.Next() {
		if err := line.Scan(&user.ID); err != nil {
			return 0, err
		}
	}

	return user.ID, nil
}

func (userRepository users) FindByNameOrNick(nameOrNick string) ([]models.User, error) {
	var user models.User

	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := userRepository.db.Query("SELECT id, name, nickname, email,  created_at, updated_at FROM users WHERE name LIKE @nameOrNick OR nickname LIKE @nameOrNick", sql.Named("nameOrNick", nameOrNick))
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User
	for lines.Next() {
		if err = lines.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userRepository users) FindById(id uint64) (models.User, error) {
	var user models.User
	line, err := userRepository.db.Query("SELECT id, name, nickname, email, created_at, updated_at FROM users WHERE id = @id", sql.Named("id", id))
	if err != nil {
		return user, err
	}

	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return user, err
		}
	}

	return user, nil
}

func (userRepository users) Update(user models.User) error {
	statement, err := userRepository.db.Prepare("UPDATE users SET name = @name, nickname = @nickname, email = @email, password = @password, updated_at = @updated_at WHERE id = @id")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(sql.Named("name", user.Name), sql.Named("nickname", user.Nickname), sql.Named("email", user.Email), sql.Named("password", user.Password), sql.Named("updated_at", user.UpdatedAt), sql.Named("id", user.ID))
	if err != nil {
		return err
	}

	return nil
}

func (userRepository users) Delete(id uint64) error {
	statement, err := userRepository.db.Prepare("DELETE FROM users WHERE id = @id")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}
