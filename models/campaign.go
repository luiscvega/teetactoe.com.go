package models

type Campaign struct {
	Id     int64
	Name   string
	UserId int64
}

//CREATE TABLE campaigns
//(
//id SERIAL PRIMARY KEY,
//name VARCHAR(255),
//user_id INTEGER
//);
