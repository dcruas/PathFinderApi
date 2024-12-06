package skillchecksvc

import (
	"pathfinderapi/models/requests/skillcheckreq"
	"pathfinderapi/models/responses/skillcheckres"
)

func CalculatePossibleOutcomes(request skillcheckreq.SkillCheckRequest) skillcheckres.SkillCheckOutcomesResponse {
	const totalRolls = 20
	modifier := request.Modifier
	dc := request.Dc

	results := skillcheckres.SkillCheckOutcomesResponse{
		CriticalFailures:  0,
		Failures:          0,
		Successes:         0,
		CriticalSuccesses: 0,
	}

	for roll := 1; roll <= totalRolls; roll++ {
		totalRoll := roll + modifier

		// Natural 1
		if roll == 1 {
			if totalRoll >= dc+10 {
				results.Successes++
			} else if totalRoll >= dc {
				results.Failures++
			} else {
				results.CriticalFailures++
			}
			continue
		}

		// Natural 20
		if roll == 20 {
			if totalRoll >= dc {
				results.CriticalSuccesses++
			} else if totalRoll >= dc-10 {
				results.Successes++
			} else {
				results.Failures++
			}
			continue
		}

		// Normal calculation for other rolls
		if totalRoll < dc-10 {
			results.CriticalFailures++
		} else if totalRoll < dc {
			results.Failures++
		} else if totalRoll < dc+10 {
			results.Successes++
		} else {
			results.CriticalSuccesses++
		}
	}

	return results
}
