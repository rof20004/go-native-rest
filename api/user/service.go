package user

import (
	"database/sql"
	"errors"
	"log"

	"github.com/rof20004/go-native-rest/database"
)

func list() ([]*User, error) {
	var db = database.GetConnection()

	query := "SELECT * FROM user"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var (
		users []*User
		id    sql.NullInt64
		name  sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, &User{ID: id.Int64, Name: name.String})
	}

	return users, nil
}

func create(user *User) error {
	var db = database.GetConnection()

	query := "INSERT INTO user(name) VALUES(?)"

	res, err := db.Exec(query, user.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func read(id int) (*User, error) {
	var db = database.GetConnection()

	query := "SELECT name FROM user WHERE id = ?"

	var (
		name sql.NullString
	)
	err := db.QueryRow(query, id).Scan(&name)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:   int64(id),
		Name: name.String,
	}, nil
}

func update(user *User) error {
	var db = database.GetConnection()

	query := "UPDATE user SET name = ? WHERE id = ?"

	res, err := db.Exec(query, user.Name, user.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	success, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if success == 0 {
		return errors.New("Usuário não encontrado")
	}

	return nil
}
