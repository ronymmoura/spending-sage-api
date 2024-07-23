package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ronymmoura/spending-sage-api/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomFixedEntry(t *testing.T) FixedEntry {
	user := createRandomUser(t)
	category := createRandomCategory(t)
	origin := createRandomOrigin(t)

	arg := CreateFixedEntryParams{
		UserID:     user.ID,
		CategoryID: category.ID,
		OriginID:   origin.ID,
		Name:       util.RandomString(10),
		DueDate:    time.Now(),
		PayDay:     time.Now(),
		Amount:     int32(util.RandomInt(1, 1000)),
		Owner:      util.RandomOwner(),
	}

	fixedEntry, err := testStore.CreateFixedEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fixedEntry)

	require.NotZero(t, fixedEntry.ID)
	require.Equal(t, arg.Name, fixedEntry.Name)
	require.NotZero(t, fixedEntry.DueDate)
	require.WithinDuration(t, arg.DueDate, fixedEntry.DueDate, time.Hour*24)
	require.NotZero(t, fixedEntry.PayDay)
	require.WithinDuration(t, arg.PayDay, fixedEntry.PayDay, time.Hour*24)
	require.NotZero(t, fixedEntry.Amount)
	require.Equal(t, arg.Amount, fixedEntry.Amount)
	require.Equal(t, arg.Owner, fixedEntry.Owner)

	return fixedEntry
}

func TestCreateFixedEntry(t *testing.T) {
	createRandomFixedEntry(t)
}

func TestGetFixedEntry(t *testing.T) {
	createdFixedEntry := createRandomFixedEntry(t)

	fixedEntry, err := testStore.GetFixedEntry(context.Background(), createdFixedEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fixedEntry)

	require.NotZero(t, fixedEntry.ID)
	require.Equal(t, createdFixedEntry.ID, fixedEntry.ID)
	require.Equal(t, createdFixedEntry.Name, fixedEntry.Name)
	require.NotZero(t, fixedEntry.DueDate)
	require.WithinDuration(t, createdFixedEntry.DueDate, fixedEntry.DueDate, time.Second)
	require.NotZero(t, fixedEntry.PayDay)
	require.WithinDuration(t, createdFixedEntry.PayDay, fixedEntry.PayDay, time.Second)
	require.NotZero(t, fixedEntry.Amount)
	require.Equal(t, createdFixedEntry.Amount, fixedEntry.Amount)
	require.Equal(t, createdFixedEntry.Owner, fixedEntry.Owner)
}

func TestListFixedEntries(t *testing.T) {
	amount := 7

	var lastCreatedFixedEntry FixedEntry

	for i := 0; i < amount; i++ {
		lastCreatedFixedEntry = createRandomFixedEntry(t)
	}

	arg := SearchFixedEntriesParams{
		UserID:     lastCreatedFixedEntry.UserID,
		OriginID:   sql.NullInt64{Int64: 0, Valid: false},
		CategoryID: sql.NullInt64{Int64: 0, Valid: false},
		Owner:      sql.NullString{String: "", Valid: false},
	}

	fixedEntries, err := testStore.SearchFixedEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fixedEntries)
	require.Equal(t, 1, len(fixedEntries))
}

func TestListFixedEntriesWithOrigin(t *testing.T) {
	amount := 7

	var lastCreatedFixedEntry FixedEntry

	for i := 0; i < amount; i++ {
		lastCreatedFixedEntry = createRandomFixedEntry(t)
	}

	arg := SearchFixedEntriesParams{
		UserID:     lastCreatedFixedEntry.UserID,
		OriginID:   sql.NullInt64{Int64: lastCreatedFixedEntry.OriginID, Valid: false},
		CategoryID: sql.NullInt64{Int64: 0, Valid: false},
		Owner:      sql.NullString{String: "", Valid: false},
	}

	fixedEntries, err := testStore.SearchFixedEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fixedEntries)
	require.Equal(t, 1, len(fixedEntries))
}

func TestListFixedEntriesWithCategory(t *testing.T) {
	amount := 7

	var lastCreatedFixedEntry FixedEntry

	for i := 0; i < amount; i++ {
		lastCreatedFixedEntry = createRandomFixedEntry(t)
	}

	arg := SearchFixedEntriesParams{
		UserID:     lastCreatedFixedEntry.UserID,
		OriginID:   sql.NullInt64{Int64: 0, Valid: false},
		CategoryID: sql.NullInt64{Int64: lastCreatedFixedEntry.CategoryID, Valid: false},
		Owner:      sql.NullString{String: "", Valid: false},
	}

	fixedEntries, err := testStore.SearchFixedEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fixedEntries)
	require.Equal(t, 1, len(fixedEntries))
}

func TestListFixedEntriesWithOwner(t *testing.T) {
	amount := 7

	var lastCreatedFixedEntry FixedEntry

	for i := 0; i < amount; i++ {
		lastCreatedFixedEntry = createRandomFixedEntry(t)
	}

	arg := SearchFixedEntriesParams{
		UserID:     lastCreatedFixedEntry.UserID,
		OriginID:   sql.NullInt64{Int64: 0, Valid: false},
		CategoryID: sql.NullInt64{Int64: 0, Valid: false},
		Owner:      sql.NullString{String: lastCreatedFixedEntry.Owner, Valid: false},
	}

	fixedEntries, err := testStore.SearchFixedEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fixedEntries)
	require.Equal(t, 1, len(fixedEntries))
}

func TestListFixedEntriesWithEverything(t *testing.T) {
	amount := 7

	var lastCreatedFixedEntry FixedEntry

	for i := 0; i < amount; i++ {
		lastCreatedFixedEntry = createRandomFixedEntry(t)
	}

	arg := SearchFixedEntriesParams{
		UserID:     lastCreatedFixedEntry.UserID,
		OriginID:   sql.NullInt64{Int64: lastCreatedFixedEntry.OriginID, Valid: false},
		CategoryID: sql.NullInt64{Int64: lastCreatedFixedEntry.CategoryID, Valid: false},
		Owner:      sql.NullString{String: lastCreatedFixedEntry.Owner, Valid: false},
	}

	fixedEntries, err := testStore.SearchFixedEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fixedEntries)
	require.Equal(t, 1, len(fixedEntries))
}

func TestEditFixedEntry(t *testing.T) {
	createdFixedEntry := createRandomFixedEntry(t)

	arg := EditFixedEntryParams{
		ID:         createdFixedEntry.ID,
		CategoryID: createdFixedEntry.CategoryID,
		OriginID:   createdFixedEntry.OriginID,
		Name:       util.RandomString(10),
		DueDate:    createdFixedEntry.DueDate,
		PayDay:     createdFixedEntry.PayDay,
		Amount:     createdFixedEntry.Amount,
		Owner:      createdFixedEntry.Owner,
	}

	fixedEntry, err := testStore.EditFixedEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fixedEntry)

	require.NotZero(t, fixedEntry.ID)
	require.Equal(t, arg.Name, fixedEntry.Name)
	require.NotZero(t, fixedEntry.DueDate)
	require.WithinDuration(t, arg.DueDate, fixedEntry.DueDate, time.Hour*24)
	require.NotZero(t, fixedEntry.PayDay)
	require.WithinDuration(t, arg.PayDay, fixedEntry.PayDay, time.Hour*24)
	require.NotZero(t, fixedEntry.Amount)
	require.Equal(t, arg.Amount, fixedEntry.Amount)
	require.Equal(t, arg.Owner, fixedEntry.Owner)
}

func TestDeleteFixedEntry(t *testing.T) {
	createdFixedEntry := createRandomFixedEntry(t)

	err := testStore.DeleteFixedEntry(context.Background(), createdFixedEntry.ID)
	require.NoError(t, err)

	fixedEntry, err := testStore.GetFixedEntry(context.Background(), createdFixedEntry.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, fixedEntry)
}
