package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/joekingsleyMukundi/bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomAmount(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotZero(t, entry.ID)
	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	require.NotZero(t, entry.CreatedAt)
	return entry
}
func TestCreateEnry(t *testing.T) {
	createRandomEntry(t)
}
func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}
func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)
	args := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomInt(0, 100),
	}
	updateEntry, err := testQueries.UpdateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, updateEntry)
	require.Equal(t, entry.ID, updateEntry.ID)
	require.Equal(t, entry.AccountID, updateEntry.AccountID)
	require.Equal(t, args.Amount, updateEntry.Amount)

	require.WithinDuration(t, entry.CreatedAt, updateEntry.CreatedAt, time.Second)
}
func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
func TestListentries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	args := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, len(entries), 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
