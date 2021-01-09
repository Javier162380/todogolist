package cmd

type runOptions struct {
	debugMode        bool
	dryRun           bool
	TaskName         string
	TaskLabel        string
	TaskDate         string
	TaskEndDate      string
	Periodicity      string
	TaskTimeInvested int
	TaskPeriodicity  string
	Limit            int
}
