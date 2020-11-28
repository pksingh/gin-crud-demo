package appProps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	t.Run("GetAllWithoutLoad", func(t *testing.T) {
		assert.Nil(t, GetAll())
	})

	t.Run("GetAllSuccess", func(t *testing.T) {
		assert.NoError(t, Load("./resources"))
		assert.NotNil(t, GetAll())
	})
}
