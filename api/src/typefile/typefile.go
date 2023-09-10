package typefile

import (
	"time"
)

type Staff struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type Staffs []Staff

var columnsStaff = []string{"id", "email", "password", "name"}

type Event struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Img             string    `json:"img"`
	Date            time.Time `json:"date"`
	Venue           string    `json:"venue"`
	Castid          int       `json:"castid"`
	EventCategoryId int       `json:"eventcategoryid"`
	Description     string    `json:"description"`
}
type Events []Event

var columnsEvent = []string{"id", "name", "img", "date", "venue", "castid", "eventcategoryid", "description"}

type EventCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type EventCategories []EventCategory

var columnsEventCategories = []string{"id", "name"}

type Cast struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Casts []Cast

var columnCast = []string{"id", "name"}

var Columns = map[string][]string{
	"staff":         columnsStaff,
	"event":         columnsEvent,
	"eventcategory": columnsEventCategories,
	"cast":          columnCast,
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (u User) CreateUserMap() map[string]string {
	userMap := make(map[string]string)
	userMap["Name"] = u.Name
	userMap["Email"] = u.Email
	userMap["Password"] = u.Password
	return userMap
}
