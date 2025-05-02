package entities

// Visual details associated with the display of an entity in the client.
type VisualDetails struct {
	// The range rings to display around an entity.
	RangeRings RangeRings `json:"rangeRings"`
}

type RangeRings struct {
	// The minimum range ring distance, specified in meters.
	MinDistanceM float64 `json:"minDistanceM"`

	// The maximum range ring distance, specified in meters.
	MaxDistanceM float64 `json:"maxDistanceM"`

	// The count of range rings.
	RingCount uint32 `json:"ringCount"`

	// The color of range rings, specified in hex string.
	RingLineColor RingLineColor `json:"ringLineColor"`
}

type RingLineColor struct {
	// The amount of red in the color as a value in the interval [0, 1].
	Red float32 `json:"red"`

	// The amount of green in the color as a value in the interval [0, 1].
	Green float32 `json:"green"`

	// The amount of blue in the color as a value in the interval [0, 1].
	Blue float32 `json:"blue"`

	// The fraction of this color that should be applied to the pixel. That is, the final pixel
	// color is defined by the equation:
	// pixel color = alpha * (this color) + (1.0 - alpha) * (background color)
	// This means that a value of 1.0 corresponds to a solid color, whereas a value of 0.0
	// corresponds to a completely transparent color. This uses a wrapper message rather than a simple
	// float scalar so that it is possible to distinguish between a default value and the value
	// being unset. If omitted, this color object is rendered as a solid color (as if the alpha value
	// had been explicitly given a value of 1.0).
	Alpha float32 `json:"alpha"`
}
