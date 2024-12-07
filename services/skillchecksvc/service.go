package skillchecksvc

import (
	"pathfinderapi/models/enums/skillcheckenum" // Importando o pacote do enum
	"pathfinderapi/models/requests/skillcheckreq"
	"pathfinderapi/models/responses/skillcheckres"
)

func CheckOutcome(totalRoll int, dc int) skillcheckenum.AbilityCheckResultsType {
	if totalRoll >= dc+10 {
		return skillcheckenum.CriticalSuccess
	} else if totalRoll >= dc {
		return skillcheckenum.Success
	} else if totalRoll > dc-10 {
		return skillcheckenum.Failure
	}

	return skillcheckenum.CriticalFailure
}

func AdjustDegree(outcome skillcheckenum.AbilityCheckResultsType, roll int) skillcheckenum.AbilityCheckResultsType {
	if roll == 1 && outcome > skillcheckenum.CriticalFailure {
		return outcome - 1
	} else if roll == 20 && outcome < skillcheckenum.CriticalSuccess {
		return outcome + 1
	}

	return outcome
}

func CalculateOutcomes(roll int, modifier int, dc int, results skillcheckres.SkillCheckOutcomesResponse) skillcheckres.SkillCheckOutcomesResponse {
	outcome := CheckOutcome(roll+modifier, dc)
	check := AdjustDegree(outcome, roll)
	switch check {
	case skillcheckenum.CriticalFailure:
		results.CriticalFailures++
	case skillcheckenum.Failure:
		results.Failures++
	case skillcheckenum.Success:
		results.Successes++
	case skillcheckenum.CriticalSuccess:
		results.CriticalSuccesses++
	}

	if roll >= 20 {
		return results
	}

	return CalculateOutcomes(roll+1, modifier, dc, results)
}

func CalculatePossibleOutcomes(request skillcheckreq.SkillCheckRequest) skillcheckres.SkillCheckOutcomesResponse {
	modifier := request.Modifier
	dc := request.Dc

	results := skillcheckres.SkillCheckOutcomesResponse{
		CriticalFailures:  0,
		Failures:          0,
		Successes:         0,
		CriticalSuccesses: 0,
	}

	return CalculateOutcomes(1, modifier, dc, results)
}
