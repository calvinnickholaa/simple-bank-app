package sqlc

import (
	"context"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:  util.RandomMoney(),
	}
	Entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Entry)

	require.Equal(t, arg.AccountID, Entry.AccountID)
	require.Equal(t, arg.Amount, Entry.Amount)

	require.NotZero(t, Entry.ID)
	require.NotZero(t, Entry.CreatedAt)

	return Entry
}
func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	Entry1 := createRandomEntry(t, account)
	Entry2, err := testQueries.GetEntry(context.Background(), Entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Entry2)

	require.Equal(t, Entry1.ID, Entry2.ID)
	require.Equal(t, Entry1.AccountID, Entry2.AccountID)
	require.Equal(t, Entry1.Amount, Entry2.Amount)
	require.WithinDuration(t, Entry1.CreatedAt.Time, Entry2.CreatedAt.Time, time.Second)
}

func TestListEntry(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
