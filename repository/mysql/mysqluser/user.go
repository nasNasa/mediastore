package mysqluser

import (
	"context"
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
	res := d.conn.Conn().QueryRowContext(ctx, `select * from users where phone_number = ?`, phoneNumber)
	fmt.Println("result", res)
	user, err := scaner(res)
	if err != nil {
		return userEntity.User{}, err
	}
	return user, nil
}

func scaner(scanner mysql.Scanner) (userEntity.User, error) {
	var createdAt time.Time
	user := userEntity.User{}
	err := scanner.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &createdAt)
	if err != nil {
		return userEntity.User{}, err
	}
	return user, nil
}
