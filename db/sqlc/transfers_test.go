package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/joekingsleyMukundi/bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfers(t *testing.T) Transfer {
	saccount := createRandomAccount(t)
	raccount := createRandomAccount(t)
	args := CreateTransferParams{
		FromAccountID: saccount.ID,
		ToAccountID:   raccount.ID,
		Amount:        util.RandomAmount(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}
func TestCreateTransfers(t *testing.T) {
	createRandomTransfers(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfers(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfers(t)
	args := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomAmount(),
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, args.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}
func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfers(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}
func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfers(t)
	}
	args := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, len(transfers), 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
