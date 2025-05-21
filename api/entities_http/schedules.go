package entities

type Schedules struct {
	Schedules []Schedule `json:"schedules"`
}

type Schedule struct {
	Windows []Window `json:"windows"`

	ScheduleId string `json:"scheduleId"`

	ScheduleType ScheduleType `json:"scheduleType"`
}

type Window struct {
	// TYPO
	// in UTC, describes when and at what cadence this window starts, in the quartz flavor of cron
	// examples: This schedule is begins at 7:00:00am UTC everyday between Monday and Friday 0 0 7 ? *
	// MON-FRI * This schedule begins every 5 minutes starting at 12:00:00pm UTC until 8:00:00pm UTC
	// everyday 0 0/5 12-20 * * ? * This schedule begins at 12:00:00pm UTC on March 2nd 2023 0 0 12 2 3 ? 2023
	CronExpression string `json:"cronExpression"`

	// describes the duration
	DurationMillis string `json:"durationMillis"`
}

type ScheduleType int

const (
	SCHEDULE_TYPE_INVALID = iota
	SCHEDULE_TYPE_ZONE_ENABLED
	SCHEDULE_TYPE_ZONE_TEMP_ENABLED
)
