package gcm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateNumeric(t *testing.T) {
	assert.Nil(t, ValidateDurationBetween(1*time.Hour, 5*time.Hour)(6*time.Minute))
}
