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
func (u *School) GetAll() ([]*School, error) {
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
