package application

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ericles-Miller/crudUserInMemory/internal/database/pgstore"
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("user not found")

type Application struct {
	queries *pgstore.Queries
}

func New(queries *pgstore.Queries) *Application {
	return &Application{queries: queries}
}

func (a *Application) FindAll(ctx context.Context) ([]pgstore.User, error) {
	return a.queries.FindAll(ctx)
}

func (a *Application) FindById(ctx context.Context, id uuid.UUID) (pgstore.User, error) {
	user, err := a.queries.FindById(ctx, id)
	
	if errors.Is(err, sql.ErrNoRows) {
		return pgstore.User{}, ErrNotFound
	}

	return user, err
}

func (a *Application) Insert(ctx context.Context, body UserBody) (pgstore.User, error) {
	if err := ValidateUser(body); err != nil {
		return pgstore.User{}, err
	}

	return a.queries.Insert(ctx, pgstore.InsertParams{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Biography: body.Biography,
	})
}

func (a *Application) Update(ctx context.Context, id uuid.UUID, body UserBody) (pgstore.User, error) {
	if err := ValidateUser(body); err != nil {
		return pgstore.User{}, err
	}

	user, err := a.queries.Update(ctx, pgstore.UpdateParams{
		ID:        id,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Biography: body.Biography,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return pgstore.User{}, ErrNotFound
	}

	return user, err
}

func (a *Application) Delete(ctx context.Context, id uuid.UUID) (pgstore.User, error) {
	user, err := a.queries.Delete(ctx, id)
	
	if errors.Is(err, sql.ErrNoRows) {
		return pgstore.User{}, ErrNotFound
	}

	return user, err
}
