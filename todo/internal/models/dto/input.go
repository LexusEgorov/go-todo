package dto

import "time"

type Register struct {
	TgID     int    `json:"tgId"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserUpdate struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TaskUpdate struct {
	ID       int        `json:"id"`
	Title    string     `json:"title"`
	Text     string     `json:"text"`
	Status   TaskStatus `json:"status"`
	Deadline time.Time  `json:"deadline"`
}
