package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	hackerFields = `account, name, hacked, zone, date`
	hackerSets   = `account=:account, name=:name, hacked=:hacked, zone=:zone, date=:date`
	hackerBinds  = `:account, :name, :hacked, :zone, :date`
)

//GetHacker will grab data from storage
func (s *Storage) GetHacker(hacker *model.Hacker) (err error) {
	err = s.db.Get(hacker, fmt.Sprintf("SELECT id, %s FROM hackers WHERE id = ?", hackerFields), hacker.ID)
	if err != nil {
		return
	}
	return
}

//CreateHacker will grab data from storage
func (s *Storage) CreateHacker(hacker *model.Hacker) (err error) {
	if hacker == nil {
		err = fmt.Errorf("Must provide hacker")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO hackers(%s)
		VALUES (%s)`, hackerFields, hackerBinds), hacker)
	if err != nil {
		return
	}
	hackerID, err := result.LastInsertId()
	if err != nil {
		return
	}
	hacker.ID = hackerID
	return
}

//ListHacker will grab data from storage
func (s *Storage) ListHacker(pageSize int64, pageNumber int64) (hackers []*model.Hacker, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM hackers 
		ORDER BY id DESC LIMIT %d OFFSET %d`, hackerFields, pageSize, pageSize*pageNumber))
	if err != nil {
		return
	}

	for rows.Next() {
		hacker := model.Hacker{}
		if err = rows.StructScan(&hacker); err != nil {
			return
		}
		hackers = append(hackers, &hacker)
	}
	return
}

//ListHackerCount will grab data from storage
func (s *Storage) ListHackerCount() (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(id) FROM hackers`)
	if err != nil {
		return
	}
	return
}

//SearchHackerByMessage will grab data from storage
func (s *Storage) SearchHackerByMessage(hacker *model.Hacker) (hackers []*model.Hacker, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM hackers 
		WHERE name like ? ORDER BY id DESC`, hackerFields), "%"+hacker.Hacked+"%")
	if err != nil {
		return
	}

	for rows.Next() {
		hacker := model.Hacker{}
		if err = rows.StructScan(&hacker); err != nil {
			return
		}
		hackers = append(hackers, &hacker)
	}
	return
}

//EditHacker will grab data from storage
func (s *Storage) EditHacker(hacker *model.Hacker) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE hackers SET %s WHERE id = :id`, hackerSets), hacker)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

//DeleteHacker will grab data from storage
func (s *Storage) DeleteHacker(hacker *model.Hacker) (err error) {
	result, err := s.db.Exec(`DELETE FROM hackers WHERE id = ?`, hacker.ID)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}
