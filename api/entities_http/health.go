package entities

import "time"

// General health of the entity as reported by the entity.
type Health struct {
	// Status indicating whether the entity is able to communicate with Entity Manager.
	ConnectionStatus ConnectionStatus `json:"connectionStatus"`

	// Top-level health status; typically a roll-up of individual component healths.
	HealthStatus HealthStatus `json:"healthStatus"`

	// Health of individual components running on this Entity.
	Components []Component `json:"components"`

	// The update time for the top-level health information. If this timestamp is unset,
	// the data is assumed to be most recent
	UpdateTime time.Time `json:"updateTime"`

	// Active alerts indicate a critical change in system state sent by the asset that must
	// be made known to an operator or consumer of the common operating picture. Alerts are
	// different from ComponentHealth messages--an active alert does not necessarily indicate
	// a component is in an unhealthy state. For example, an asset may trigger an active alert
	// based on fuel levels running low. Alerts should be removed from this list when their
	// conditions are cleared. In other words, only active alerts should be reported here.
	ActiveAlerts []ActiveAlert `json:"activeAlerts"`
}

type ConnectionStatus int

const (
	CONNECTION_STATUS_INVALID = iota
	CONNECTION_STATUS_ONLINE
	CONNECTION_STATUS_OFFLINE
)

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

type HealthStatus int

const (
	HEALTH_STATUS_INVALID = iota
	HEALTH_STATUS_HEALTHY
	HEALTH_STATUS_WARN
	HEALTH_STATUS_FAIL
	HEALTH_STATUS_OFFLINE
	HEALTH_STATUS_NOT_READY
)

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
