package skillcheckenum

type AbilityCheckResultsType int

const (
	CriticalFailure AbilityCheckResultsType = iota // 0
	Failure                                        // 1
	Success                                        // 2
	CriticalSuccess                                // 3
)
