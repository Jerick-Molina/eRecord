package db

import (
	"context"
	"fmt"

	"eRecord/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) CreateAccountParams {

	arg := CreateAccountParams{
		FirstName: util.RandomName(),
		LastName:  "testSubject",

		Role:      util.RandomRole(),
		Password:  util.RandomPassword(10),
		CompanyId: int(util.RandomNumber(10)),
	}
	fullName := fmt.Sprintf("%s%s", arg.FirstName, arg.LastName)
	arg.Email = fmt.Sprintf("%s%d@testSubject.com", fullName, arg.CompanyId)
	err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	return arg
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestEmailExistValidationTrue(t *testing.T) {
	acc := createRandomAccount(t)
	err := testQueries.EmailExistValidation(context.Background(), acc.Email)

	require.NoError(t, err)
}

func TestEmailExistValidationFalse(t *testing.T) {
	acc := createRandomAccount(t)
	acc.Email = acc.Email + "nomail"
	err := testQueries.EmailExistValidation(context.Background(), acc.Email)

	require.Error(t, err)
}

//Should return the essentials to provide a access key .
func TestSignInValidationTrue(t *testing.T) {
	randAccount := createRandomAccount(t)
	fmt.Println(randAccount)
	account, err := testQueries.SignInValidation(context.Background(), randAccount.Email, randAccount.Password)

	require.NoError(t, err)

	require.NotEmpty(t, account.Id)

	require.Equal(t, randAccount.CompanyId, account.CompanyId)
	require.Equal(t, randAccount.Role, account.Role)
}

func TestSignInValidationFalse(t *testing.T) {
	randAccount := createRandomAccount(t)
	randAccount.Password = randAccount.Password + "123"
	account, err := testQueries.SignInValidation(context.Background(), randAccount.Email, randAccount.Password)

	require.Error(t, err)

	require.Empty(t, account)
}
