package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateGeoShape(entity *components.Entity, new_geoshape *components.GeoShape) error {
	if new_geoshape.Point != nil {
		if new_geoshape.Point.Position != nil {
			err := UpdatePosition(entity, new_geoshape.Point.Position)
			if err != nil {
				return err
			}
		}
	}
	if new_geoshape.Line != nil {
		if new_geoshape.Line.Position != nil {
			err := UpdatePosition(entity, new_geoshape.Line.Position)
			if err != nil {
				return err
			}
		}
	}
	if new_geoshape.Polygon != nil {
		if len(new_geoshape.Polygon.Rings) != 0 {
			entity.GeoShape.Polygon.Rings = new_geoshape.Polygon.Rings
		}
	}
	if new_geoshape.Ellipse != nil {
		entity.GeoShape.Ellipse = new_geoshape.Ellipse
	}
	if new_geoshape.Ellipsoid != nil {
		entity.GeoShape.Ellipsoid = new_geoshape.Ellipsoid
	}
	return nil
}
