// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type Championship struct {
	ID   int64
	Name string
}

type Team struct {
	ID             int64
	Name           string
	Country        string
	Class          string
	ChampionshipID sql.NullInt64
}
