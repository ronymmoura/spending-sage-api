package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ronymmoura/spending-sage-api/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomPaymentHistory(t *testing.T) FixedEntryPaymentHistory {
	entry := createRandomFixedEntry(t)

	arg := CreateFixedEntryPaymentHistoryParams{
		EntryID: entry.ID,
		Amount:  int32(util.RandomInt(1, 1000)),
		Date:    time.Now(),
	}

	paymentHistory, err := testStore.CreateFixedEntryPaymentHistory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, paymentHistory)

	require.NotZero(t, paymentHistory.ID)
	require.Equal(t, arg.EntryID, paymentHistory.EntryID)
	require.Equal(t, arg.Amount, paymentHistory.Amount)
	require.WithinDuration(t, arg.Date, paymentHistory.Date, time.Second)

	return paymentHistory
}

func TestCreatePaymentHistory(t *testing.T) {
	createRandomPaymentHistory(t)
}

func TestGetPaymentHistory(t *testing.T) {
	createdHistory := createRandomPaymentHistory(t)

	paymentHistory, err := testStore.GetFixedEntryPaymentHistory(context.Background(), createdHistory.ID)
	require.NoError(t, err)
	require.NotEmpty(t, paymentHistory)

	require.NotZero(t, paymentHistory.ID)
	require.Equal(t, createdHistory.EntryID, paymentHistory.EntryID)
	require.Equal(t, createdHistory.Amount, paymentHistory.Amount)
	require.WithinDuration(t, createdHistory.Date, paymentHistory.Date, time.Second)
}

func TestListPaymentHistory(t *testing.T) {
	amount := 7

	var lastCreatedPaymentHistory FixedEntryPaymentHistory

	for i := 0; i < amount; i++ {
		lastCreatedPaymentHistory = createRandomPaymentHistory(t)
	}

	paymentHistory, err := testStore.ListFixedEntryPaymentHistory(context.Background(), lastCreatedPaymentHistory.EntryID)
	require.NoError(t, err)
	require.NotEmpty(t, paymentHistory)
	require.Equal(t, 1, len(paymentHistory))
}

func TestDeletePaymentHistory(t *testing.T) {
	createdHistory := createRandomPaymentHistory(t)

	err := testStore.DeleteFixedEntryPaymentHistory(context.Background(), createdHistory.ID)
	require.NoError(t, err)

	paymentHistory, err := testStore.GetFixedEntryPaymentHistory(context.Background(), createdHistory.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, paymentHistory)
}
