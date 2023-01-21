package storage

import (
	"database/sql"
	"os"

	"github.com/gustagcosta/go-api/types"
	_ "github.com/lib/pq"
)

type PgStorage struct {
	db *sql.DB
}

func (s *PgStorage) Connect() error {
	db, err := sql.Open("postgres", os.Getenv("PG_CONNECTION_STRING"))
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *PgStorage) GetDogs() (*[]types.Dog, error) {
	s.Connect()
	var dogs []types.Dog

	rows, err := s.db.Query(`SELECT * FROM dogs`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var d types.Dog

		err = rows.Scan(&d.ID, &d.Name, &d.Age)

		if err != nil {
			return nil, err
		}

		dogs = append(dogs, d)
	}

	defer s.db.Close()

	return &dogs, nil
}

func (s *PgStorage) StoreDog(name string, age int) (int, error) {
	s.Connect()
	var id int

	sql := `INSERT INTO dogs (name, age) VALUES ($1, $2) RETURNING id`

	err := s.db.QueryRow(sql, name, age).Scan(&id)
	if err != nil {
		return 0, err
	}

	defer s.db.Close()

	return id, nil
}

func (s *PgStorage) GetDog(id int) (*types.Dog, error) {
	s.Connect()
	var dog types.Dog

	rows, err := s.db.Query(`SELECT * FROM dogs WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&dog.ID, &dog.Name, &dog.Age)

		if err != nil {
			return nil, err
		}
	}

	if dog.ID == 0 {
		return nil, nil
	}

	defer s.db.Close()

	return &dog, nil
}

func (s *PgStorage) UpdateDog(newDog *types.Dog) error {
	dog, err := s.GetDog(newDog.ID)
	if err != nil {
		return err
	}

	s.Connect()

	dog.Name = newDog.Name
	dog.Age = newDog.Age

	_, err = s.db.Exec(`UPDATE dogs SET name = $1, age = $2 WHERE id = $3`, dog.Name, dog.Age, dog.ID)
	if err != nil {
		return err
	}

	defer s.db.Close()

	return nil
}

func (s *PgStorage) DeleteDog(id int) error {
	s.Connect()

	err := s.db.QueryRow(`DELETE FROM dogs WHERE id = $1`, id)
	if err != nil {
		return err.Err()
	}

	defer s.db.Close()

	return nil
}
