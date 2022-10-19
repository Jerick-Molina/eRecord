package db

import (
	"context"
	"eRecord/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCompany(t *testing.T) string {
	uniqueId := util.RandomChars(10)
	err := testQueries.CreateCompany(context.Background(), "Test_"+util.RandomCompany(5), uniqueId)

	require.NoError(t, err)
	return uniqueId
}

func TestCreateCompany(t *testing.T) {
	createRandomCompany(t)
}

func TestFindCompanyWithUniqueId(t *testing.T) {
	uniqueId := createRandomCompany(t)

	id, err := testQueries.FindCompanyWithUniqueId(context.Background(), uniqueId)
	var invalidNum int = -1

	require.NoError(t, err)
	fmt.Print(id)
	require.GreaterOrEqual(t, id, invalidNum, "Errr!?")
}
