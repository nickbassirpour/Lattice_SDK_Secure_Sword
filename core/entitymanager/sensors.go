package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateSensors(existingSensors []*components.Sensor, newSensors []*components.Sensor) ([]*components.Sensor, error) {
	existing_sensors := make(map[string]*components.Sensor)
	for _, sensor := range existingSensors {
		existing_sensors[*sensor.SensorId] = sensor
	}

	for _, new_sensor := range newSensors {
		if exists, found := existing_sensors[*new_sensor.SensorId]; found {
			err := UpdateSensor(exists, new_sensor)
			if err != nil {
				return nil, err
			}
		} else {
			existingSensors = append(existingSensors, new_sensor)
		}
	}
	return existingSensors, nil
}

func UpdateSensor(existing_sensor *components.Sensor, new_sensor *components.Sensor) error {
	if new_sensor.OperationalState != nil {
		existing_sensor.OperationalState = new_sensor.OperationalState
	}
	if new_sensor.SensorType != nil {
		existing_sensor.SensorType = new_sensor.SensorType
	}
	if new_sensor.SensorDescription != nil {
		existing_sensor.SensorDescription = new_sensor.SensorDescription
	}
	if new_sensor.LastDetectionTimestamp != nil {
		existing_sensor.LastDetectionTimestamp = new_sensor.LastDetectionTimestamp
	}
	if new_sensor.RfConfiguration != nil {
		err := UpdateRFConfiguration(existing_sensor.RfConfiguration, new_sensor.RfConfiguration)
		if err != nil {
			return err
		}
	}
	if len(new_sensor.FieldsOfView) > 0 {
		err := UpdateFieldsOfView(existing_sensor.FieldsOfView, new_sensor.FieldsOfView)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateRFConfiguration(existing_rf_configuration *components.RFConfiguration, new_rf_configuration *components.RFConfiguration) error {
	if len(new_rf_configuration.FrequencyRangeHz) > 0 {
		existing_rf_configuration.FrequencyRangeHz = append(existing_rf_configuration.FrequencyRangeHz, new_rf_configuration.FrequencyRangeHz...)
	}
	if len(new_rf_configuration.BandwidthRangeHz) > 0 {
		existing_rf_configuration.BandwidthRangeHz = append(existing_rf_configuration.BandwidthRangeHz, new_rf_configuration.BandwidthRangeHz...)
	}
	return nil
}

func UpdateFieldsOfView(existing_fovs []*components.FieldOfView, new_fovs []*components.FieldOfView) error {
	existing_fovs_map := make(map[int32]*components.FieldOfView)
	for _, existing_fov := range existing_fovs {
		existing_fovs_map[*existing_fov.FovId] = existing_fov
	}

	for _, new_fov := range new_fovs {
		if exists, found := existing_fovs_map[*new_fov.FovId]; found {
			err := UpdateFieldOfView(exists, new_fov)
			if err != nil {
				return err
			}
		} else {
			existing_fovs = append(existing_fovs, new_fov)
		}
	}
	return nil
}

func UpdateFieldOfView(existing_fov *components.FieldOfView, new_fov *components.FieldOfView) error {
	if new_fov.MountId != nil {
		existing_fov.MountId = new_fov.MountId
	}
	if new_fov.ProjectFrustum != nil {
		if new_fov.ProjectFrustum.UpperLeft != nil {
			existing_fov.ProjectFrustum.UpperLeft = new_fov.ProjectFrustum.UpperLeft
		}
		if new_fov.ProjectFrustum.UpperRight != nil {
			existing_fov.ProjectFrustum.UpperRight = new_fov.ProjectFrustum.UpperRight
		}
		if new_fov.ProjectFrustum.BottomLeft != nil {
			existing_fov.ProjectFrustum.BottomLeft = new_fov.ProjectFrustum.BottomLeft
		}
		if new_fov.ProjectFrustum.BottomRight != nil {
			existing_fov.ProjectFrustum.BottomRight = new_fov.ProjectFrustum.BottomRight
		}
	}
	if new_fov.ProjectedCenterRay != nil {
		err := UpdatePosition(existing_fov.ProjectedCenterRay, new_fov.ProjectedCenterRay)
		if err != nil {
			return err
		}
	}
	if new_fov.CenterRayPose != nil {
		err := UpdateRelativePose(existing_fov.CenterRayPose, new_fov.CenterRayPose)
		if err != nil {
			return err
		}
	}
	if new_fov.HorizontalFov != nil {
		existing_fov.HorizontalFov = new_fov.HorizontalFov
	}
	if new_fov.VerticalFov != nil {
		existing_fov.VerticalFov = new_fov.VerticalFov
	}
	if new_fov.Range != nil {
		existing_fov.Range = new_fov.Range
	}
	if new_fov.Mode != nil {
		existing_fov.Mode = new_fov.Mode
	}
	return nil
}
