package sampledata

import (
	"time"

	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SampleHealth() *components.Health {
	return &components.Health{
		ConnectionStatus: Pointer(components.ConnectionStatus_CONNECTION_STATUS_OFFLINE),
		HealthStatus:     Pointer(components.HealthStatus_HEALTH_STATUS_HEALTHY),
		Components:       sampleComponents(),
		UpdateTime:       timestamppb.New(time.Now()),
		ActiveAlerts:     sampleActiveAlerts(),
	}
}

func sampleComponents() []*components.Component {
	componentSlice := []*components.Component{}
	component := &components.Component{
		Id:         Pointer("199013"),
		Name:       Pointer("67 Radar"),
		Health:     Pointer(components.HealthStatus_HEALTH_STATUS_HEALTHY),
		Messages:   sampleMessages(),
		UpdateTime: timestamppb.New(time.Now()),
	}
	componentSlice = append(componentSlice, component)
	return componentSlice
}

func sampleMessages() []*components.Message {
	messages := []*components.Message{}
	message := &components.Message{
		Status:  Pointer(components.HealthStatus_HEALTH_STATUS_HEALTHY),
		Message: Pointer("This system is fine."),
	}
	messages = append(messages, message)
	return messages
}

func sampleActiveAlerts() []*components.ActiveAlert {
	activeAlerts := []*components.ActiveAlert{}
	activeAlert := &components.ActiveAlert{
		AlertCode:        Pointer("100-1079-88923"),
		Description:      Pointer("sample active alert description"),
		Level:            Pointer(components.AlertLevel_ALERT_LEVEL_ADVISORY),
		ActivatedTime:    timestamppb.New(time.Now()),
		ActiveConditions: sampleActiveConditions(),
	}
	activeAlerts = append(activeAlerts, activeAlert)
	return activeAlerts
}

func sampleActiveConditions() []*components.ActiveCondition {
	activeConditions := []*components.ActiveCondition{}
	activeCondition := &components.ActiveCondition{
		ConditionCode: Pointer("1002-23954"),
		Description:   Pointer("sample active condition"),
	}
	activeConditions = append(activeConditions, activeCondition)
	return activeConditions
}
