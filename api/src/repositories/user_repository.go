package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
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

	defer line.Close()

	return user.ID, nil
}

// GetAllUsers with name or nickname
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

// GetUser by ID
func (userRepository users) FindById(ID uint64) (models.User, error) {
	var user models.User

	lines, err := userRepository.db.Query("SELECT id, name, nickname, email, created_at, updated_at FROM users WHERE id = @id", sql.Named("id", ID))
	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update a user
func (userRepository users) Update(ID uint64, user models.User) error {
	statement, err := userRepository.db.Prepare("UPDATE users SET name = @name, nickname = @nickname, email = @email WHERE id = @id")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(sql.Named("name", user.Name), sql.Named("nickname", user.Nickname), sql.Named("email", user.Email), sql.Named("id", ID))
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

func (userRepository users) AuthenticationFindByEmail(email string) (models.User, error) {
	var user models.User

	lines, err := userRepository.db.Query("SELECT id, name, nickname, email, password, created_at, updated_at FROM users WHERE email = @email", sql.Named("email", email))
	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (userRepository users) Follow(userID uint64, followerID uint64) error {
	var user models.User

	lines, err := userRepository.db.Query("select user_id from followers where user_id = @user_id and follower_id = @follower_id", sql.Named("user_id", userID), sql.Named("follower_id", followerID))
	if err != nil {
		return err
	}

	defer lines.Close()

	if lines.Next() {
		if err = lines.Scan(&user.ID); err != nil {
			return err
		}
	}

	if user.ID != 0 {
		return errors.New("already following")
	}

	statement, err := userRepository.db.Prepare("INSERT INTO followers (user_id, follower_id) VALUES (@user_id, @follower_id)")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(sql.Named("user_id", userID), sql.Named("follower_id", followerID))
	if err != nil {
		return err
	}

	return nil
}

// select with not exists (select * from followers where user_id = @user_id and follower_id = @follower_id)

// SELECT TOP 0 * INTO TempTable FROM MyTable;

// INSERT INTO TempTable (Column1, Column2) VALUES ('Value1', 'Value2');

// INSERT INTO MyTable SELECT A.* FROM TempTable A LEFT JOIN MyTable B ON A.UniqueField = B.UniqueField WHERE B.UniqueField IS NULL;

// DROP TABLE TempTable;
