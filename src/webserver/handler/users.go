package handler

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	us "github.com/asepnur/iskandar/src/module/users"
	"github.com/asepnur/iskandar/src/webserver/template"
	"github.com/julienschmidt/httprouter"
)

const (
	sessionPrefix     = "session:"
	character         = "!QAZ@WSX#EDC$RFV%TGB^YHN&UJM*IK<(OL>)P:?_{+}|1qaz2wsx3edc4rfv5tgb6yhn7ujm8ik,9ol.0p-[=]"
	listPrefixSession = "session:list:"
)

type (
	Config struct {
		SessionKey string `json:"sessionkey"`
	}
)

var (
	c                  Config
	errSessionNotlogin = errors.New("SessionNotLogin")
	charMaxIndex       = len(character)
)

type User struct {
	UserID    int    `json:"id"`
	UserEmail string `json:"email"`
	FullName  string `json:"name"`
	MSISDN    string `json:"mdisdn"`
}

// SelectUserHandler ..
func SelectUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cookie, err := SetSession()
	http.SetCookie(w, cookie)
	var resp []User
	users, err := us.GetMultipleUser()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	for _, el := range users {
		resp = append(resp, User{
			UserID:    el.UserID,
			UserEmail: el.UserEmail,
			FullName:  el.FullName,
			MSISDN:    el.MSISDN,
		})
	}
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// SelectUserFilterHandler ..
func SelectUserFilterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cookie, _ := r.Cookie(c.SessionKey)
	session := strings.Trim(cookie.Value, " ")
	fmt.Println(session)
	return
}

func SetSession() (*http.Cookie, error) {
	var cookie string
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 32; i++ {
		cookie = cookie + string(character[rand.Intn(charMaxIndex)])
	}
	return &http.Cookie{
		Name:    "_ci",
		Expires: time.Now().AddDate(0, 1, 0),
		Value:   cookie,
		Path:    "/",
	}, nil
}
