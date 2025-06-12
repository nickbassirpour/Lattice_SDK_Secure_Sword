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
	components := []*components.Component{}
	component := &components.Component{
		Id:         Pointer("199013"),
		Name:       Pointer("67 Radar"),
		Health:     Pointer(components.HealthStatus_HEALTH_STATUS_HEALTHY),
		Messages:   sampleMessages(),
		UpdateTime: timestamppb.New(time.Now()),
	}
	components = append(components, component)
	return components
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

}

/*

type Component struct {
	// Consistent internal ID for this component.
	Id string `json:"id"`

	// Display name for this component.
	Name string `json:"name"`

	// Health for this component.
	Health HealthStatus `json:"health"`

	// Human-readable describing the component state. These messages
	// should be understandable by end users.
	Messages []Message `json:"messages"`

	// The last update time for this specific component. If this timestamp
	// is unset, the data is assumed to be most recent
	UpdateTime time.Time `json:"updateTime"`
}
*/
