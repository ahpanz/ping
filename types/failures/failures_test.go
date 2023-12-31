package failures

import (
	"github.com/statping-ng/statping-ng/database"
	"github.com/statping-ng/statping-ng/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInit(t *testing.T) {
	t.Parallel()
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	SetDB(db)
	db.AutoMigrate(&Failure{})

	t.Run("Test Samples", func(t *testing.T) {
		require.Nil(t, Samples())
	})
}
