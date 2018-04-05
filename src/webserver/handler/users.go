package handler

import (
	"net/http"

	us "github.com/asepnur/iskandar/src/module/users"
	"github.com/asepnur/iskandar/src/webserver/template"
	"github.com/julienschmidt/httprouter"
)

// User struct :: save value
type User struct {
	UserID    int    `json:"id"`
	UserEmail string `json:"email"`
	FullName  string `json:"name"`
	MSISDN    string `json:"mdisdn"`
}

// SelectUserHandler ..
func SelectUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	_, err = us.GetVisitor()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

// SelectUserFilterHandler ..
func SelectUserFilterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	return
}
