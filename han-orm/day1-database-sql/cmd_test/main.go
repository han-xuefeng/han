package main

import (
	"geeorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main()  {

	engine, _ := geeorm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	s.Raw("SELECT Name FROM User")
	row := s.QueryRow()
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}