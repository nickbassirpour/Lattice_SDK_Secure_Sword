package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateSchedules(entity *components.Entity, newSchedules *components.Schedules) error {
	if len(newSchedules.Schedules) > 0 {
		existingSchedules := make(map[string]*components.Schedule)
		for _, schedule := range entity.Schedules.Schedules {
			existingSchedules[*schedule.ScheduleId] = schedule
		}

		for _, newSchedule := range newSchedules.Schedules {
			if exists, found := existingSchedules[*newSchedule.ScheduleId]; found {
				err := UpdateSchedule(exists, newSchedule)
				if err != nil {
					return nil
				}
			}
		}
	}
	return nil
}

func UpdateSchedule(existingSchedule *components.Schedule, newSchedule *components.Schedule) error {
	if newSchedule.ScheduleType != nil {
		existingSchedule.ScheduleType = newSchedule.ScheduleType
	}
	if newSchedule.Windows != nil && len(newSchedule.Windows) > 0 {
		existingSchedule.Windows = append(existingSchedule.Windows, newSchedule.Windows...)
	}
	return nil
}
