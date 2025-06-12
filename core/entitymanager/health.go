package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateHealth(entity *components.Entity, newHealth *components.Health) error {
	if newHealth.ConnectionStatus != nil {
		entity.Health.ConnectionStatus = newHealth.ConnectionStatus
	}
	if newHealth.HealthStatus != nil {
		entity.Health.HealthStatus = newHealth.HealthStatus
	}
	if newHealth.Components != nil {
		existingHealthComponents := make(map[string]*components.Component)
		for _, component := range entity.Health.Components {
			existingHealthComponents[*component.Id] = component
		}
		for _, newComponent := range newHealth.Components {
			if exists, found := existingHealthComponents[*newComponent.Id]; found {
				err := UpdateComponent(exists, newComponent)
				if err != nil {
					return err
				}
			}
		}
	}
	if newHealth.UpdateTime != nil {
		entity.Health.UpdateTime = newHealth.UpdateTime
	}
	if newHealth.ActiveAlerts != nil && len(newHealth.ActiveAlerts) > 0 {
		entity.Health.ActiveAlerts = append(entity.Health.ActiveAlerts, newHealth.ActiveAlerts...)
	}
	return nil
}

func UpdateComponent(existingComponent *components.Component, newComponent *components.Component) error {
	if newComponent.Name != nil {
		existingComponent.Name = newComponent.Name
	}
	if newComponent.Health != nil {
		existingComponent.Health = newComponent.Health
	}
	if newComponent.Messages != nil && len(newComponent.Messages) > 0 {
		existingComponent.Messages = append(existingComponent.Messages, newComponent.Messages...)
	}
	if newComponent.UpdateTime != nil {
		existingComponent.UpdateTime = newComponent.UpdateTime
	}
	return nil
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

type Message struct {
	// The status associated with this message.
	Status HealthStatus `json:"status"`

	// The human-readable content of the message.
	Message string `json:"message"`
}


type ActiveAlert struct {
	// Short, machine-readable code that describes this alert. This code is intended
	// to provide systems off-asset with a lookup key to retrieve more detailed
	// information about the alert.
	AlertCode string `json:"alertCode"`

	// Human-readable description of this alert. The description is intended for
	// display in the UI for human understanding and should not be used for machine
	// processing. If the description is fixed and the vehicle controller provides
	// no dynamic substitutions, then prefer lookup based on alert_code.
	Description string `json:"description"`

	// Alert level (Warning, Caution, or Advisory).
	Level AlertLevel `json:"level"`

	// Time at which this alert was activated.
	ActivatedTime time.Time `json:"activatedTime"`

	// Set of conditions which have activated this alert.
	ActiveConditions []ActiveCondition `json:"activeConditions"`
}

type AlertLevel int

const (
	ALERT_LEVEL_INVALID = iota
	ALERT_LEVEL_ADVISORY
	ALERT_LEVEL_CAUTION
	ALERT_LEVEL_WARNING
)

type ActiveCondition struct {
	// Short, machine-readable code that describes this condition. This code is
	// intended to provide systems off-asset with a lookup key to retrieve more
	// detailed information about the condition.
	ConditionCode string `json:"conditionCode"`

	// Human-readable description of this condition. The description is intended
	// for display in the UI for human understanding and should not be used for
	// machine processing. If the description is fixed and the vehicle controller
	// provides no dynamic substitutions, then prefer lookup based on condition_code.
	Description string `json:"description"`
}

*/
