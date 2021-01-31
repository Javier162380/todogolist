package cmd

type runCommandOptions struct {
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

type runAPIOptions struct {
	host string
	port uint
}
