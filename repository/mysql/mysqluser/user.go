package mysqluser

import (
	"context"
	"database/sql"
	"fmt"
	"mediaStorer/entity/userEntity"
	"mediaStorer/repository/mysql"
	"time"
)

func (d *DB) RegisterToDb(user userEntity.User) (userEntity.User, error) {
	// Todo create param for structs
	res, err := d.conn.Conn().Exec(`insert into users(name, phone_number,email, password) values(?, ?, ?,?)`,
		user.Name, user.PhoneNumber, user.Email, user.Password)
	if err != nil {
		return userEntity.User{}, fmt.Errorf("can't execute command: %w", err)
	}

	id, _ := res.LastInsertId()
	user.Id = uint(id)

	return user, nil
}

func (d *DB) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (userEntity.User, error) {
	fmt.Println("phonenumber", phoneNumber)

	res := d.conn.Conn().QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	fmt.Println("here", res)

	user, err := scaner(res)
	if err != nil {
		if err == sql.ErrNoRows {
			return userEntity.User{}, fmt.Errorf("no rows")
		}

		// TODO - log unexpected error for better observability
		return userEntity.User{}, fmt.Errorf("cant scan")
	}

	return user, nil
}

func scaner(scanner mysql.Scanner) (userEntity.User, error) {
	var createdAt time.Time
	user := userEntity.User{}
	err := scanner.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &createdAt)

	return user, err
}
