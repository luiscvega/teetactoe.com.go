package models

type Campaign struct {
	Id     int
	Name   string
	UserId int
}

//CREATE TABLE campaigns
//(
//id SERIAL PRIMARY KEY,
//name VARCHAR(255),
//user_id INTEGER
//);
