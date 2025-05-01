package entities

// A component that describes the shape of a geo-entity.
type GeoShape struct {
	Point     Point     `json:"point"`
	Line      Line      `json:"line"`
	Polygon   Polygon   `json:"polygon"`
	Ellipse   Ellipse   `json:"ellipse"`
	Ellipsoid Ellipsoid `json:"ellipsoid"`
}

// A point shaped geo-entity.
type Point struct {
	Position Position `json:"position"`
}

// A line shaped geo-entity.
type Line struct {
	Position Position `json:"position"`
}

// A polygon shaped geo-entity.
type Polygon struct {
	// An array of LinearRings where the first item is the
	// exterior ring and subsequent items are interior rings.
	Rings []Ring `json:"rings"`

	// An extension hint that this polygon is a rectangle. When true this implies several things:
	// exactly 1 linear ring with 5 points (starting corner, 3 other corners and start again)
	// each point has the same altitude corresponding with the plane of the rectangle
	// each point has the same height (either all present and equal, or all not present)
	IsRectanble bool `json:"isRectangle"`
}

type Ring struct {
	Positions []LinearRing `json:"positions"`
}

type LinearRing struct {
	Position Position `json:"position"`
	HeightM  float32  `json:"heightM"`
}

// An ellipse shaped geo-entity. For a circle, the major and minor axis would be the same values.
// This shape is NOT Geo-JSON compatible.
type Ellipse struct {
	// Defines the distance from the center point of the ellipse to the furthest distance on the perimeter in meters.
	SemiMajorAxisM float64 `json:"semiMajorAxisM"`

	// Defines the distance from the center point of the ellipse to the shortest distance on the perimeter in meters.
	SemiMinorAxisM float64 `json:"semiMinorAxisM"`

	// The orientation of the semi-major relative to true north in degrees from clockwise: 0-180 due to
	// symmetry across the semi-minor axis.
	OrientationD float64 `json:"orientationD"`

	// Optional height above entity position to extrude in meters. A non-zero value creates an elliptic cylinder
	HeightM float64 `json:"heightM"`
}

type Ellipsoid struct {
	// Defines the distance from the center point to the surface along the forward axis
	ForwardAxisM float64 `json:"forwardAxisM"`

	// Defines the distance from the center point to the surface along the side axis
	SideAxisM float64 `json:"sideAxisM"`

	// Defines the distance from the center point to the surface along the up axis
	UpAxisM float64 `json:"upAxisM"`
}
