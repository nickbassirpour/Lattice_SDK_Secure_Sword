package entities

type Position struct {
	// WGS84 geodetic latitude in decimal degrees.
	LatitudeDegrees float64 `json:"latitudeDegrees"`

	// WGS84 longitude in decimal degrees.
	LongitudeDegrees float64 `json:"longitudeDegrees"`

	// Altitude as height above ellipsoid (WGS84) in meters.
	// float64 Value wrapper is used to distinguish optional from default 0.
	AltitudeHaeMeters float64 `json:"altitudeHaeMeters"`

	// Altitude as AGL (Above Ground Level) if the upstream data source has this value set.
	// This value represents the entity's height above the terrain. This is typically measured
	// with a radar altimeter or by using a terrain tile set lookup. If the value is not set
	// from the upstream, this value is not set.
	AltitudeAglMeters float64 `json:"altitudeAglMeters"`

	// Altitude as ASF (Above Sea Floor) if the upstream data source has this value set.
	// If the value is not set from the upstream, this value is not set.
	AltitudeAsfMeters float64 `json:"altitudeAsfMeters"`

	// The depth of the entity from the surface of the water through sensor measurements
	// based on differential pressure between the interior and exterior of the vessel.
	// If the value is not set from the upstream, this value is not set.
	PressureDepthMeters float64 `json:"pressureDepthMeters"`
}
