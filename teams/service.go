package teams

import (
	"context"
	"database/sql"

	"github.com/nstehr/vroom-core/db"
)

type Service struct {
	queries db.Queries
}

func NewService(database *sql.DB) *Service {
	queries := db.New(database)
	return &Service{queries: *queries}
}

func (s *Service) GetTeams(ctx context.Context) ([]Team, error) {
	rows, err := s.queries.ListTeams(ctx)
	if err != nil {
		return nil, err
	}
	var teams []Team
	for _, row := range rows {
		t := Team{Name: row.Name, Country: row.Country, Class: row.Class, Championship: row.Championship}
		teams = append(teams, t)
	}
	return teams, nil
}

func (s *Service) CreateTeam(ctx context.Context, team Team) error {
	params := db.CreateTeamParams{Name: team.Name, Country: team.Country, Class: team.Class, ChampionshipID: sql.NullInt64{Int64: int64(1), Valid: true}}
	_, err := s.queries.CreateTeam(ctx, params)
	return err
}
