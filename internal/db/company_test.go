package db

import (
	"context"
	"eRecord/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCompany(t *testing.T) string {
	uniqueCode := util.RandomChars(10)
	err := testQueries.CreateCompany(context.Background(), "Test_"+util.RandomCompany(5), uniqueCode)

	require.NoError(t, err)
	return uniqueCode
}

func findcompanyWithUniqueId(t *testing.T, invCode string) int {

	id, err := testQueries.FindCompanyWithUniqueId(context.Background(), invCode)
	var invalidNum int = -1

	require.NoError(t, err)
	require.GreaterOrEqual(t, id, invalidNum, "Errr!?")

	return id
}
func TestCreateCompany(t *testing.T) {
	createRandomCompany(t)
}

func TestFindCompanyWithUniqueId(t *testing.T) {
	InvCode := createRandomCompany(t)
	compId := findcompanyWithUniqueId(t, InvCode)

	require.NotEmpty(t, compId)

}

func TestValidateInvitationToken(t *testing.T) {

}
