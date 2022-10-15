package db

import (
	"context"
	"eRecord/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCompany(t *testing.T) {

	err := testQueries.CreateCompany(context.Background(), util.RandomCompany(5))

	require.NoError(t, err)
	return
}

func TestCreateCompany(t *testing.T) {
	createRandomCompany(t)
}
