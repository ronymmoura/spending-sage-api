package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/ronymmoura/spending-sage-api/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	name := util.RandomString(10)

	category, err := testStore.CreateCategory(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.NotZero(t, category.ID)
	require.Equal(t, name, category.Name)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	createdCategory := createRandomCategory(t)

	category, err := testStore.GetCategory(context.Background(), createdCategory.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, createdCategory.ID, category.ID)
	require.Equal(t, createdCategory.Name, category.Name)
}

func TestListCategories(t *testing.T) {
	amount := 5

	for i := 0; i < amount; i++ {
		createRandomCategory(t)
	}

	categories, err := testStore.ListCategories(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, categories)
}

func TestDeleteCategory(t *testing.T) {
	createdCategory := createRandomCategory(t)

	err := testStore.DeleteCategory(context.Background(), createdCategory.ID)
	require.NoError(t, err)

	category, err := testStore.GetCategory(context.Background(), createdCategory.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, category)
}
