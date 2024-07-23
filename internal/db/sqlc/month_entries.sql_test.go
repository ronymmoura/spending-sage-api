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

func createRandomMonthEntry(t *testing.T) MonthEntry {
	month := createRandomMonth(t)
	category := createRandomCategory(t)
	origin := createRandomOrigin(t)

	arg := CreateMonthEntryParams{
		MonthID:    month.ID,
		CategoryID: category.ID,
		OriginID:   origin.ID,
		Name:       util.RandomString(10),
		DueDate:    time.Now(),
		PayDate:    time.Now(),
		Amount:     int32(util.RandomInt(1, 1000)),
		Owner:      util.RandomOwner(),
	}

	monthEntry, err := testStore.CreateMonthEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, monthEntry)

	require.NotZero(t, monthEntry.ID)
	require.Equal(t, arg.Name, monthEntry.Name)
	require.NotZero(t, monthEntry.DueDate)
	require.WithinDuration(t, arg.DueDate, monthEntry.DueDate, time.Hour*24)
	require.NotZero(t, monthEntry.PayDate)
	require.WithinDuration(t, arg.PayDate, monthEntry.PayDate, time.Hour*24)
	require.NotZero(t, monthEntry.Amount)
	require.Equal(t, arg.Amount, monthEntry.Amount)
	require.Equal(t, arg.Owner, monthEntry.Owner)

	return monthEntry
}

func TestCreateMonthEntry(t *testing.T) {
	createRandomMonthEntry(t)
}

func TestGetMonthEntry(t *testing.T) {
	createdMonthEntry := createRandomMonthEntry(t)

	monthEntry, err := testStore.GetMonthEntry(context.Background(), createdMonthEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, monthEntry)

	require.NotZero(t, monthEntry.ID)
	require.Equal(t, createdMonthEntry.ID, monthEntry.ID)
	require.Equal(t, createdMonthEntry.Name, monthEntry.Name)
	require.NotZero(t, monthEntry.DueDate)
	require.WithinDuration(t, createdMonthEntry.DueDate, monthEntry.DueDate, time.Second)
	require.NotZero(t, monthEntry.PayDate)
	require.WithinDuration(t, createdMonthEntry.PayDate, monthEntry.PayDate, time.Second)
	require.NotZero(t, monthEntry.Amount)
	require.Equal(t, createdMonthEntry.Amount, monthEntry.Amount)
	require.Equal(t, createdMonthEntry.Owner, monthEntry.Owner)
}

func TestListMonthEntries(t *testing.T) {
	amount := 7

	var lastCreatedMonthEntry MonthEntry

	for i := 0; i < amount; i++ {
		lastCreatedMonthEntry = createRandomMonthEntry(t)
	}

	arg := SearchMonthEntriesParams{
		MonthID:    lastCreatedMonthEntry.MonthID,
		OriginID:   sql.NullInt64{Int64: 0, Valid: false},
		CategoryID: sql.NullInt64{Int64: 0, Valid: false},
		Owner:      sql.NullString{String: "", Valid: false},
	}

	monthEntries, err := testStore.SearchMonthEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, monthEntries)
	require.Equal(t, 1, len(monthEntries))
}

func TestListMonthEntriesWithOrigin(t *testing.T) {
	amount := 7

	var lastCreatedMonthEntry MonthEntry

	for i := 0; i < amount; i++ {
		lastCreatedMonthEntry = createRandomMonthEntry(t)
	}

	arg := SearchMonthEntriesParams{
		MonthID:    lastCreatedMonthEntry.MonthID,
		OriginID:   sql.NullInt64{Int64: lastCreatedMonthEntry.OriginID, Valid: false},
		CategoryID: sql.NullInt64{Int64: 0, Valid: false},
		Owner:      sql.NullString{String: "", Valid: false},
	}

	monthEntries, err := testStore.SearchMonthEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, monthEntries)
	require.Equal(t, 1, len(monthEntries))
}

func TestListMonthEntriesWithCategory(t *testing.T) {
	amount := 7

	var lastCreatedMonthEntry MonthEntry

	for i := 0; i < amount; i++ {
		lastCreatedMonthEntry = createRandomMonthEntry(t)
	}

	arg := SearchMonthEntriesParams{
		MonthID:    lastCreatedMonthEntry.MonthID,
		OriginID:   sql.NullInt64{Int64: 0, Valid: false},
		CategoryID: sql.NullInt64{Int64: lastCreatedMonthEntry.CategoryID, Valid: false},
		Owner:      sql.NullString{String: "", Valid: false},
	}

	monthEntries, err := testStore.SearchMonthEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, monthEntries)
	require.Equal(t, 1, len(monthEntries))
}

func TestListMonthEntriesWithOwner(t *testing.T) {
	amount := 7

	var lastCreatedMonthEntry MonthEntry

	for i := 0; i < amount; i++ {
		lastCreatedMonthEntry = createRandomMonthEntry(t)
	}

	arg := SearchMonthEntriesParams{
		MonthID:    lastCreatedMonthEntry.MonthID,
		OriginID:   sql.NullInt64{Int64: 0, Valid: false},
		CategoryID: sql.NullInt64{Int64: 0, Valid: false},
		Owner:      sql.NullString{String: lastCreatedMonthEntry.Owner, Valid: false},
	}

	monthEntries, err := testStore.SearchMonthEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, monthEntries)
	require.Equal(t, 1, len(monthEntries))
}

func TestListMonthEntriesWithEverything(t *testing.T) {
	amount := 7

	var lastCreatedMonthEntry MonthEntry

	for i := 0; i < amount; i++ {
		lastCreatedMonthEntry = createRandomMonthEntry(t)
	}

	arg := SearchMonthEntriesParams{
		MonthID:    lastCreatedMonthEntry.MonthID,
		OriginID:   sql.NullInt64{Int64: lastCreatedMonthEntry.OriginID, Valid: false},
		CategoryID: sql.NullInt64{Int64: lastCreatedMonthEntry.CategoryID, Valid: false},
		Owner:      sql.NullString{String: lastCreatedMonthEntry.Owner, Valid: false},
	}

	monthEntries, err := testStore.SearchMonthEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, monthEntries)
	require.Equal(t, 1, len(monthEntries))
}

func TestEditMonthEntry(t *testing.T) {
	createdMonthEntry := createRandomMonthEntry(t)

	arg := EditMonthEntryParams{
		ID:         createdMonthEntry.ID,
		CategoryID: createdMonthEntry.CategoryID,
		OriginID:   createdMonthEntry.OriginID,
		Name:       util.RandomString(10),
		DueDate:    createdMonthEntry.DueDate,
		PayDate:    createdMonthEntry.PayDate,
		Amount:     createdMonthEntry.Amount,
		Owner:      createdMonthEntry.Owner,
	}

	monthEntry, err := testStore.EditMonthEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, monthEntry)

	require.NotZero(t, monthEntry.ID)
	require.Equal(t, arg.Name, monthEntry.Name)
	require.NotZero(t, monthEntry.DueDate)
	require.WithinDuration(t, arg.DueDate, monthEntry.DueDate, time.Hour*24)
	require.NotZero(t, monthEntry.PayDate)
	require.WithinDuration(t, arg.PayDate, monthEntry.PayDate, time.Hour*24)
	require.NotZero(t, monthEntry.Amount)
	require.Equal(t, arg.Amount, monthEntry.Amount)
	require.Equal(t, arg.Owner, monthEntry.Owner)
}

func TestDeleteMonthEntry(t *testing.T) {
	createdMonthEntry := createRandomMonthEntry(t)

	err := testStore.DeleteMonthEntry(context.Background(), createdMonthEntry.ID)
	require.NoError(t, err)

	monthEntry, err := testStore.GetMonthEntry(context.Background(), createdMonthEntry.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, monthEntry)
}
