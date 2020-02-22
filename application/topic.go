package application

import (
	"github.com/trastanechora/pi-engine/config"
	"github.com/trastanechora/pi-engine/domain"
	"github.com/trastanechora/pi-engine/infrastructure/persistence"
)

// GetTopic returns a topic by id
func GetTopic(id int) (*domain.Topic, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewTopicRepositoryWithRDB(conn)
	return repo.Get(id)
}

// GetAllTopic return all topics
func GetAllTopic() ([]domain.Topic, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewTopicRepositoryWithRDB(conn)
	return repo.GetAll()
}

// AddTopic saves new topic
func AddTopic(name string, slug string) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewTopicRepositoryWithRDB(conn)
	u := &domain.Topic{
		Name: name,
		Slug: slug,
	}
	return repo.Save(u)
}

// RemoveTopic do remove topic by id
func RemoveTopic(id int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewTopicRepositoryWithRDB(conn)
	return repo.Remove(id)
}

// UpdateTopic do update topic by id
func UpdateTopic(p domain.Topic, id int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewTopicRepositoryWithRDB(conn)
	p.ID = uint(id)

	return repo.Update(&p)
}
