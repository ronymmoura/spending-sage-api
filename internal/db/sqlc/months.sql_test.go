package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func createRandomMonth(t *testing.T) Month {
	user := createRandomUser(t)

	arg := CreateMonthParams{
		UserID: user.ID,
		Date:   time.Now(),
	}

	month, err := testStore.CreateMonth(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, month)

	require.NotZero(t, month.ID)
	require.NotZero(t, month.Date)
	require.Equal(t, arg.UserID, month.UserID)

	return month
}

func createRandomMonthForUser(t *testing.T, user User) Month {
	arg := CreateMonthParams{
		UserID: user.ID,
		Date:   time.Now(),
	}

	month, err := testStore.CreateMonth(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, month)

	require.NotZero(t, month.ID)
	require.NotZero(t, month.Date)
	require.Equal(t, arg.UserID, month.UserID)

	return month
}

func TestCreateMonth(t *testing.T) {
	createRandomMonth(t)
}

func TestGetMonth(t *testing.T) {
	createdMonth := createRandomMonth(t)

	month, err := testStore.GetMonth(context.Background(), createdMonth.ID)
	require.NoError(t, err)
	require.NotEmpty(t, month)

	require.NotZero(t, month.ID)
	require.NotZero(t, month.Date)
	require.Equal(t, createdMonth.UserID, month.UserID)
}

func TestListMonths(t *testing.T) {
	amount := 7

	user := createRandomUser(t)

	for i := 0; i < amount; i++ {
		createRandomMonthForUser(t, user)
	}

	limit := 5

	arg := ListMonthsParams{
		UserID: user.ID,
		Limit:  int32(limit),
		Offset: 0,
	}

	months, err := testStore.ListMonths(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, months)
	require.Equal(t, limit, len(months))

	for _, month := range months {
		require.Equal(t, user.ID, month.UserID)
	}

	arg = ListMonthsParams{
		UserID: user.ID,
		Limit:  int32(limit),
		Offset: 1 * int32(limit),
	}

	months, err = testStore.ListMonths(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, months)
	require.Equal(t, 2, len(months))

	for _, month := range months {
		require.Equal(t, user.ID, month.UserID)
	}
}

func TestListMonthsEmpty(t *testing.T) {
	user := createRandomUser(t)

	arg := ListMonthsParams{
		UserID: user.ID,
		Limit:  5,
		Offset: 0,
	}

	months, err := testStore.ListMonths(context.Background(), arg)
	require.NoError(t, err)
	require.Empty(t, months)
}

func TestDeleteMonth(t *testing.T) {
	createdMonth := createRandomMonth(t)

	err := testStore.DeleteMonth(context.Background(), createdMonth.ID)
	require.NoError(t, err)

	month, err := testStore.GetMonth(context.Background(), createdMonth.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, month)
}
