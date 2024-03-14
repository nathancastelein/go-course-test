package mocking

import "database/sql"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

type DatabaseService struct {
	db *sql.DB
}

func NewDatabaseService(db *sql.DB) *DatabaseService {
	return &DatabaseService{
		db: db,
	}
}

func (service *DatabaseService) ListUsers() ([]*User, error) {
	rows, err := service.db.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (service *DatabaseService) FindUserByID(id int) (*User, error) {
	rows, err := service.db.Query("SELECT id, first_name, last_name FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &User{}
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, sql.ErrNoRows
}
