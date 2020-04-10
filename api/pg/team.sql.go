// Code generated by sqlc. DO NOT EDIT.
// source: team.sql

package pg

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTeam = `-- name: CreateTeam :one
INSERT INTO team (organization_id, created_at, name) VALUES ($1, $2, $3) RETURNING team_id, created_at, name, organization_id
`

type CreateTeamParams struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at"`
	Name           string    `json:"name"`
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	row := q.db.QueryRowContext(ctx, createTeam, arg.OrganizationID, arg.CreatedAt, arg.Name)
	var i Team
	err := row.Scan(
		&i.TeamID,
		&i.CreatedAt,
		&i.Name,
		&i.OrganizationID,
	)
	return i, err
}

const deleteTeamByID = `-- name: DeleteTeamByID :exec
DELETE FROM team WHERE team_id = $1
`

func (q *Queries) DeleteTeamByID(ctx context.Context, teamID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTeamByID, teamID)
	return err
}

const getAllTeams = `-- name: GetAllTeams :many
SELECT team_id, created_at, name, organization_id FROM team
`

func (q *Queries) GetAllTeams(ctx context.Context) ([]Team, error) {
	rows, err := q.db.QueryContext(ctx, getAllTeams)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.TeamID,
			&i.CreatedAt,
			&i.Name,
			&i.OrganizationID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeamByID = `-- name: GetTeamByID :one
SELECT team_id, created_at, name, organization_id FROM team WHERE team_id = $1
`

func (q *Queries) GetTeamByID(ctx context.Context, teamID uuid.UUID) (Team, error) {
	row := q.db.QueryRowContext(ctx, getTeamByID, teamID)
	var i Team
	err := row.Scan(
		&i.TeamID,
		&i.CreatedAt,
		&i.Name,
		&i.OrganizationID,
	)
	return i, err
}

const getTeamsForOrganization = `-- name: GetTeamsForOrganization :many
SELECT team_id, created_at, name, organization_id FROM team WHERE organization_id = $1
`

func (q *Queries) GetTeamsForOrganization(ctx context.Context, organizationID uuid.UUID) ([]Team, error) {
	rows, err := q.db.QueryContext(ctx, getTeamsForOrganization, organizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.TeamID,
			&i.CreatedAt,
			&i.Name,
			&i.OrganizationID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
