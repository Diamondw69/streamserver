package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Links struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
	Status bool   `json:"status"`
	Device string `json:"device"`
}
type LinkModel struct {
	DB *sql.DB
}

func (m LinkModel) Insert(links *Links) error {
	query := `
INSERT INTO links (name, link,status, device)
VALUES ($1, $2, $3, $4)
RETURNING id`
	args := []any{links.Name, links.Link, links.Status, links.Device}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&links.ID)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (m LinkModel) GetByName(name string) (*Links, error) {
	query := `
SELECT id, name,link, device
FROM links
WHERE name = $1`
	var links Links
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, name).Scan(
		&links.ID,
		&links.Name,
		&links.Link,
		&links.Device,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &links, nil
}

func (m LinkModel) GetByStatus(status bool) (*Links, error) {
	query := `
SELECT id, name,link, device
FROM links
WHERE status = $1`
	var links Links
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, status).Scan(
		&links.ID,
		&links.Name,
		&links.Link,
		&links.Device,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &links, nil
}
func (m LinkModel) GetByStatusAndDevice(status bool, device string) (*Links, error) {
	query := `
SELECT id, name,link, device
FROM links
WHERE status = $1 and device = $2`
	var links Links
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, status, device).Scan(
		&links.ID,
		&links.Name,
		&links.Link,
		&links.Device,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &links, nil
}

func (m LinkModel) UpdateAllLinks(device string) error {
	query := `
UPDATE links
set status=false
where device=$1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_ = m.DB.QueryRowContext(ctx, query, device)
	return nil
}
func (m LinkModel) GetAllLinks() (*[]Links, error) {
	query := `
SELECT *
FROM links`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []Links
	for rows.Next() {
		var item Links
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Link,
			&item.Status,
			&item.Device,
		)
		if err != nil {
			return nil, err // Update this to return an empty Metadata struct.
		}
		links = append(links, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err // Update this to return an empty Metadata struct.
	}

	return &links, nil
}
func (m LinkModel) GetAllWorkingLinks() (*[]Links, error) {
	query := `
SELECT *
FROM links
where status=true`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []Links
	for rows.Next() {
		var item Links
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Link,
			&item.Status,
			&item.Device,
		)
		if err != nil {
			return nil, err // Update this to return an empty Metadata struct.
		}
		links = append(links, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err // Update this to return an empty Metadata struct.
	}

	return &links, nil
}
func (m LinkModel) UpdateLink(name string) error {
	query := `
UPDATE links
SET status = true
WHERE name = $1`
	args := []any{
		name,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
