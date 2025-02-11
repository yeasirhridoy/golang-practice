package database

import (
	"context"
	"github.com/stretchr/testify/require"
	"practice/utilities"
	"testing"
	"time"
)

func createNewUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:     "test",
		Email:    "user" + utilities.RandomString(5) + "@example.com",
		Password: "password",
	}

	account, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.Password, account.Password)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)

	return account
}

func TestQueries_CreateUser(t *testing.T) {
	createNewUser(t)
}

func TestQueries_GetUser(t *testing.T) {
	newAccount := createNewUser(t)
	fetchedAccount, err := testQueries.GetUser(context.Background(), newAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)

	require.Equal(t, newAccount.ID, fetchedAccount.ID)
	require.Equal(t, newAccount.Name, fetchedAccount.Name)
	require.Equal(t, newAccount.Email, fetchedAccount.Email)
	require.Equal(t, newAccount.Password, fetchedAccount.Password)
	require.WithinDuration(t, newAccount.CreatedAt, fetchedAccount.CreatedAt, time.Second)
	require.WithinDuration(t, newAccount.UpdatedAt, fetchedAccount.UpdatedAt, time.Second)
}

func TestQueries_UpdateUser(t *testing.T) {
	newAccount := createNewUser(t)

	arg := UpdateUserParams{
		ID:       newAccount.ID,
		Name:     "New Name",
		Email:    "updated" + utilities.RandomString(5) + "@example.com",
		Password: "new_password",
	}

	updatedAccount, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, arg.ID, updatedAccount.ID)
	require.Equal(t, arg.Name, updatedAccount.Name)
	require.Equal(t, arg.Email, updatedAccount.Email)
	require.Equal(t, arg.Password, updatedAccount.Password)
	require.WithinDuration(t, newAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)
	require.WithinDuration(t, newAccount.UpdatedAt, updatedAccount.UpdatedAt, time.Second)
}

func TestQueries_DeleteUser(t *testing.T) {
	newAccount := createNewUser(t)
	err := testQueries.DeleteUser(context.Background(), newAccount.ID)
	require.NoError(t, err)

	deletedAccount, err := testQueries.GetUser(context.Background(), newAccount.ID)
	require.Error(t, err)
	require.Empty(t, deletedAccount)
}

func TestQueries_ListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createNewUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
