package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/ronymmoura/spending-sage-api/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrigin(t *testing.T) Origin {
	arg := CreateOriginParams{
		Name: util.RandomString(5),
		Type: util.RandomString(5),
	}

	origin, err := testStore.CreateOrigin(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, origin)

	require.NotZero(t, origin.ID)
	require.Equal(t, arg.Name, origin.Name)
	require.Equal(t, arg.Type, origin.Type)

	return origin
}

func TestCreateOrigin(t *testing.T) {
	createRandomOrigin(t)
}

func TestGetOrigin(t *testing.T) {
	createdOrigin := createRandomOrigin(t)

	origin, err := testStore.GetOrigin(context.Background(), createdOrigin.ID)
	require.NoError(t, err)
	require.NotEmpty(t, origin)

	require.Equal(t, createdOrigin.ID, origin.ID)
	require.Equal(t, createdOrigin.Name, origin.Name)
	require.Equal(t, createdOrigin.Type, origin.Type)
}

func TestListOrigins(t *testing.T) {
	amount := 5

	for i := 0; i < amount; i++ {
		createRandomOrigin(t)
	}

	origins, err := testStore.ListOrigins(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, origins)
}

func TestDeleteOrigin(t *testing.T) {
	createdOrigin := createRandomOrigin(t)

	err := testStore.DeleteOrigin(context.Background(), createdOrigin.ID)
	require.NoError(t, err)

	origin, err := testStore.GetOrigin(context.Background(), createdOrigin.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, origin)
}
