package mysqluser

import (
	"fmt"
	"mediaStorer/entity/userEntity"
)

func (d *DB) RegisterToDb(user userEntity.User) (userEntity.User, error) {
	// Todo create param for structs
	res, err := d.conn.Conn().Exec(`insert into users(name, phone_number, password) values(?, ?, ?)`,
		user.Name, user.Email, user.Password)
	if err != nil {
		return userEntity.User{}, fmt.Errorf("can't execute command: %w", err)
	}

	id, _ := res.LastInsertId()
	user.Id = uint(id)

	return user, nil
}
