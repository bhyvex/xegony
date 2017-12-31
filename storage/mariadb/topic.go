package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) GetTopic(topicID int64) (topic *model.Topic, err error) {
	topic = &model.Topic{}
	err = s.db.Get(topic, "SELECT id, icon, title, forum_id FROM topic WHERE id = ?", topicID)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateTopic(topic *model.Topic) (err error) {
	if topic == nil {
		err = fmt.Errorf("Must provide topic")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO topic(title, icon)
		VALUES (:title, :icon)`, topic)
	if err != nil {
		return
	}
	topicID, err := result.LastInsertId()
	if err != nil {
		return
	}
	topic.Id = topicID
	return
}

func (s *Storage) ListTopic(forumID int64) (topics []*model.Topic, err error) {
	rows, err := s.db.Queryx(`SELECT id, title, icon, forum_id FROM topic WHERE forum_id = ? ORDER BY id DESC`, forumID)
	if err != nil {
		return
	}

	for rows.Next() {
		topic := model.Topic{}
		if err = rows.StructScan(&topic); err != nil {
			return
		}
		topics = append(topics, &topic)
	}
	return
}

func (s *Storage) EditTopic(topicID int64, topic *model.Topic) (err error) {
	topic.Id = topicID
	result, err := s.db.NamedExec(`UPDATE topic SET icon=:icon, title=:title, forum_id=:forum_id WHERE id = :id`, topic)
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

func (s *Storage) DeleteTopic(topicID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM topic WHERE id = ?`, topicID)
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

func (s *Storage) createTableTopic() (err error) {
	_, err = s.db.Exec(`CREATE TABLE topic (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  title varchar(32) NOT NULL DEFAULT '',
  owner_id int(11) unsigned NOT NULL,
  forum_id int(11) unsigned NOT NULL,
  last_modified timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  icon varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
