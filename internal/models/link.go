package models

import (
	"context"
	"errors"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

type Link struct {
	Id          int64     `json:"id"`
	LinkId      string    `json:"link_id"`
	OriginalUrl string    `json:"original_url"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
}

type LinkModel struct {
	Db *pgx.Conn
}

func (model *LinkModel) GetByLinkId(linkId string) (Link, error) {
	sql := "SELECT id, link_id, original_url, created_on, updated_on FROM links WHERE link_id=$1"
	args := []any{linkId}

	link := Link{}
	if err := model.Db.QueryRow(context.Background(), sql, args...).Scan(&link.Id, &link.LinkId, &link.OriginalUrl, &link.CreatedOn, &link.UpdatedOn); err != nil {
		if err == pgx.ErrNoRows {
			return link, errors.New("link not found")
		}
		return link, err
	}

	return link, nil
}

func (model *LinkModel) Create(originalUrl string) (string, error) {
	linkIdLength, err := strconv.Atoi(os.Getenv("LINK_ID_LENGTH"))
	if err != nil {
		linkIdLength = 5
	}

	if linkIdLength > 10 {
		linkIdLength = 10
	}

	if linkIdLength < 5 {
		linkIdLength = 5
	}

	// Generating a unique linkId
	linkId := strings.Builder{}
	isUnique := false

	for !isUnique {
		linkId.Reset()
		choices := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz"
		choicesSlice := strings.Split(choices, "")

		for range linkIdLength {
			choice := choicesSlice[rand.Intn(len(choicesSlice))]
			linkId.WriteString(choice)
		}

		var id int64
		if err := model.Db.QueryRow(context.Background(), "SELECT id FROM links WHERE link_id=$1", linkId.String()).Scan(&id); err != nil {
			if err == pgx.ErrNoRows {
				isUnique = true
			} else {
				return "", err
			}
		}
	}

	sql := "INSERT INTO links(link_id, original_url) VALUES($1, $2)"
	args := []any{linkId.String(), originalUrl}

	_, err = model.Db.Exec(context.Background(), sql, args...)
	if err != nil {
		return linkId.String(), err
	}

	return linkId.String(), nil
}
