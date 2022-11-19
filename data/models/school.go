package models

import (
	"context"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

type School struct {
	ID        int       `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAll returns a slice of all users, sorted by last name
func (s *School) GetAll() ([]*School, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from schools`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schools []*School

	for rows.Next() {
		var school School
		err := rows.Scan(
			&school.ID,
			&school.Name,
			&school.CreatedAt,
			&school.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		schools = append(schools, &school)
	}

	return schools, nil
}

func (s *School) GetOne(id int) (*School, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from schools where id = $1`

	var school School
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&school.ID,
		&school.Name,
		&school.CreatedAt,
		&school.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &school, nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (s *School) Insert(school School) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into schools (name, created_at, updated_at)
		values ($1, $2, $3) returning id`

	err := db.QueryRowContext(ctx, stmt,
		school.Name,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (s *School) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update schools set
		name = $1,
		updated_at = $2
		where id = $3
	`

	_, err := db.ExecContext(ctx, stmt,
		s.Name,
		time.Now(),
		s.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *School) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from schools where id = $1`

	_, err := db.ExecContext(ctx, stmt, s.ID)
	if err != nil {
		return err
	}

	return nil
}
