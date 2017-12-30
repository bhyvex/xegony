package mariadb

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Storage struct {
	db *sqlx.DB
}

func (s *Storage) Initialize(config string) (err error) {
	if s.db != nil {
		return
	}
	if config == "" {
		user := os.Getenv("API_DB_USERNAME")
		if len(user) == 0 {
			user = "eqemu"
		}
		pass := os.Getenv("API_DB_PASSWORD")
		if len(pass) == 0 {
			pass = "eqemu"
		}
		host := os.Getenv("API_DB_HOSTNAME")
		if len(host) == 0 {
			host = "127.0.0.1"
		}
		port := os.Getenv("API_DB_PORT")
		if len(port) == 0 {
			port = "3306"
		}
		dbname := os.Getenv("API_DB_NAME")
		if len(dbname) == 0 {
			dbname = "eqemu"
		}
		config = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, pass, host, port, dbname)
	}
	if s.db, err = sqlx.Open("mysql", config); err != nil {
		return
	}
	return
}

func (s *Storage) InsertTestData() (err error) {
	_, err = s.db.Exec(`INSERT INTO user (id, name, email, password, account_id)
	   VALUES
	   	(1, 'Test', '', '$2a$10$YV0PiWDMiuXL4e77.jv8leD3NpDCk.v41aXPn7Yyi7fBWwBa0XzzC', 1);`)
	if err != nil {
		return
	}

	_, err = s.db.Exec(`INSERT INTO account (id, name, status)
	   VALUES
	   	(1, 'Shin', 200);`)
	if err != nil {
		return
	}
	return
}

func (s *Storage) DropTables() (err error) {
	s.Initialize("")

	_, err = s.db.Exec(`SET FOREIGN_KEY_CHECKS = 0`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`drop table if exists character_data`)
	if err != nil {
		return
	}
	tables := []string{
		"account",
		"activity",
		"base",
		"bazaar",
		"character_data",
		"faction",
		"forum",
		"goal",
		"item",
		"lootdrop",
		"lootdropentry",
		"loottable",
		"loottablentry",
		"npc",
		"npcloot",
		"post",
		"task",
		"topic",
		"user",
		"zone",
	}
	for _, table := range tables {
		_, err = s.db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %s`, table))
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("Failed to delete %s", table))
			return
		}
	}
	_, err = s.db.Exec(`SET FOREIGN_KEY_CHECKS = 1`)
	if err != nil {
		return
	}
	return
}

func (s *Storage) VerifyTables() (err error) {
	if err = s.Initialize(""); err != nil {
		return
	}

	err = s.createTableAccount()
	if err != nil && !isExistErr(err) {
		return
	}
	err = s.createTableBazaar()
	if err != nil && !isExistErr(err) {
		return
	}
	err = s.createTableCharacter()
	if err != nil && !isExistErr(err) {
		return
	}
	err = s.createTableForum()
	if err != nil && !isExistErr(err) {
		return
	}
	err = s.createTableNpcLoot()
	if err != nil && !isExistErr(err) {
		return
	}
	err = s.createTablePost()
	if err != nil && !isExistErr(err) {
		return
	}
	err = s.createTableTopic()
	if err != nil && !isExistErr(err) {
		return
	}
	err = s.createTableUser()
	if err != nil && !isExistErr(err) {
		return
	}
	return
}

func isExistErr(err error) bool {
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == 1050 {
			return true
		}
	}
	return false
}
