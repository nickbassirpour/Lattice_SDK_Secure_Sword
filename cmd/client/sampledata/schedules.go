package sampledata

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func SampleSchedules() *components.Schedules {
	schedules := []*components.Schedule{
		sampleSchedule(),
	}
	return &components.Schedules{
		Schedules: schedules,
	}
}

func sampleSchedule() *components.Schedule {
	return &components.Schedule{
		ScheduleId:   Pointer("10026"),
		ScheduleType: Pointer(components.ScheduleType_SCHEDULE_TYPE_ZONE_ENABLED),
		Windows:      sampleWindows(),
	}
}

func sampleWindows() []*components.Window {
	windows := []*components.Window{}
	window := &components.Window{
		CronExpression: Pointer("0 0 7 ? * MON-FRI *"),
		DurationMillis: Pointer("1000234898"),
	}
	windows = append(windows, window)
	return windows
}
