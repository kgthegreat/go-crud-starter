package repositories

import (
	"log"

	"github.com/kgthegreat/go-crud-starter/database"
	"github.com/kgthegreat/go-crud-starter/models"
)

type TopicRepository interface {
	Create(p *models.Topic) error
	GetAll() ([]*models.Topic, error)
	FindById(id int) (*models.Topic, error)
	DeleteById(id int) error
}

type topicRepository struct {
	*database.SqliteDB
}

func NewTopicRepository(db *database.SqliteDB) TopicRepository {
	return &topicRepository{db}
}

func (topicRepository *topicRepository) Create(topic *models.Topic) error {

	stmt, err := topicRepository.DB.Prepare("INSERT INTO topics(title) values (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(topic.Title)
	if err != nil {
		return err
	}

	tId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	topic.ID = int(tId)

	return nil
}

func (topicRepository *topicRepository) GetAll() ([]*models.Topic, error) {
	var topics []*models.Topic

	rows, err := topicRepository.DB.Query("SELECT * from topics")
	if err != nil {
		log.Print(">>>>>>>>>", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := new(models.Topic)
		err := rows.Scan(&t.ID, &t.Title, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		topics = append(topics, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return topics, nil
}

func (topicRepository *topicRepository) FindById(id int) (*models.Topic, error) {
	topic := models.Topic{}

	err := topicRepository.DB.QueryRow("SELECT * from topics WHERE id = ?", id).Scan(&topic.ID, &topic.Title, &topic.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (topicRepository *topicRepository) DeleteById(id int) error {
	topic := models.Topic{}

	err := topicRepository.DB.QueryRow("DELETE from topics WHERE id = ?", id).Scan(&topic.ID, &topic.Title, &topic.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
