package users

import (
	"fmt"

	"github.com/asepnur/iskandar/src/util/conn"
)

// User ..
type User struct {
	UserID    int
	UserEmail string
	FullName  string
	MSISDN    string
}

// GetMultipleUser ..
func GetMultipleUser() ([]User, error) {
	var res []User
	query := fmt.Sprintf(`
		SELECT
			user_id,
			full_name,
			user_email,
			msisdn
		FROM
			ws_user
		LIMIT 10;
	`)
	rows, err := conn.DB.Query(query)
	if err != nil {

		return res, err
	}

	for rows.Next() {
		u := &User{}
		rows.Scan(&u.UserID, &u.FullName, &u.UserEmail, &u.MSISDN)
		res = append(res, *u)
	}
	return res, nil
}
