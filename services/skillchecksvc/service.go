package skillchecksvc

import (
	"pathfinderapi/models/enums/skillcheckenum" // Importando o pacote do enum
	"pathfinderapi/models/requests/skillcheckreq"
	"pathfinderapi/models/responses/skillcheckres"
)

func CheckOutcome(totalRoll int, dc int) skillcheckenum.AbilityCheckResultsType {

	switch true {
	case totalRoll >= dc+10:
		return skillcheckenum.CriticalSuccess
	case totalRoll >= dc:
		return skillcheckenum.Success
	case totalRoll > dc-10:
		return skillcheckenum.Failure
	default:
		return skillcheckenum.CriticalFailure
	}
}

func AdjustDegree(outcome skillcheckenum.AbilityCheckResultsType, roll int) skillcheckenum.AbilityCheckResultsType {

	switch true {
	case roll == 1 && outcome > skillcheckenum.CriticalFailure:
		return outcome - 1
	case roll == 20 && outcome < skillcheckenum.CriticalSuccess:
		return outcome + 1
	default:
		return outcome
	}
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
