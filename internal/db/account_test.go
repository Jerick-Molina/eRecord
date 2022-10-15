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

	err := testQueries.EmailExistValidation(context.Background(), "JohntestSubject4@testSubject.com")

	require.NoError(t, err)
}
func TestEmailExistValidationFalse(t *testing.T) {
	var email = "JohntestSubjec@testSubject.com"

	err := testQueries.EmailExistValidation(context.Background(), email)

	require.Error(t, err)
}
