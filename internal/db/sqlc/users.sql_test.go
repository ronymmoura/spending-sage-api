package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ronymmoura/spending-sage-api/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:    util.RandomEmail(),
		FullName: util.RandomOwner(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)

	user, err := testStore.GetUser(context.Background(), createdUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, createdUser.ID, user.ID)
	require.Equal(t, createdUser.Email, user.Email)
	require.Equal(t, createdUser.FullName, user.FullName)
	require.NotZero(t, user.CreatedAt)
	require.WithinDuration(t, createdUser.CreatedAt, user.CreatedAt, time.Second)
}

func TestGetUserByEmail(t *testing.T) {
	createdUser := createRandomUser(t)

	user, err := testStore.GetUserByEmail(context.Background(), createdUser.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, createdUser.ID, user.ID)
	require.Equal(t, createdUser.Email, user.Email)
	require.Equal(t, createdUser.FullName, user.FullName)
	require.NotZero(t, user.CreatedAt)
	require.WithinDuration(t, createdUser.CreatedAt, user.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	createdUser := createRandomUser(t)

	err := testStore.DeleteUser(context.Background(), createdUser.ID)
	require.NoError(t, err)

	user, err := testStore.GetUser(context.Background(), createdUser.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, user)
}
