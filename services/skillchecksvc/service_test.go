package skillchecksvc_test

import (
	"pathfinderapi/services/skillchecksvc"

	"pathfinderapi/models/requests/skillcheckreq"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRollStats(t *testing.T) {
	// Arrange
	request := skillcheckreq.SkillCheckRequest{
		Modifier: 15,
		Dc:       20,
	}

	// Act
	result := skillchecksvc.CalculatePossibleOutcomes(request)

	// Assert
	assert.Equal(t, 1, result.CriticalFailures, "Critical failures count mismatch")
	assert.Equal(t, 3, result.Failures, "Failures count mismatch")
	assert.Equal(t, 10, result.Successes, "Successes count mismatch")
	assert.Equal(t, 6, result.CriticalSuccesses, "Critical successes count mismatch")
}
