package users

import (
	"fmt"
	"log"

	"github.com/asepnur/iskandar/src/util/conn"
	nsq "github.com/bitly/go-nsq"
	"github.com/garyburd/redigo/redis"
)

// User ::
type User struct {
	UserID    int
	UserEmail string
	FullName  string
	MSISDN    string
}

// GetMultipleUser ::
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

// GetVisitor :: to get visitors value
func GetVisitor() (int, error) {
	var el int
	key := "visitor"
	client := conn.Redis.Get()
	el, err := redis.Int(client.Do("GET", key))
	if err != nil {
		return el, err
	}
	return el, nil
}

// IncreaseVisitor ::
func IncreaseVisitor(current string) {
	topic := "180204"
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("devel-go.tkpd:4150", config)
	err := w.Publish(topic, []byte(fmt.Sprintf("%s", current)))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
}
