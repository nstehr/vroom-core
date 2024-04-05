-- name: ListTeams :many
SELECT t.name, t.country, t.class, c.name as championship FROM team t, championship c where t.championship_id = c.id;

-- name: CreateTeam :one
INSERT INTO team (
  name, country, class, championship_id
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;