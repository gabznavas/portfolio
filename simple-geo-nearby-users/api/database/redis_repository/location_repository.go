package redisrepository

import (
	"api/database/models"
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisLocationRepository interface {
	PutLocation(ctx context.Context, location models.Location) (success bool, err error)
	GetLocationsByPosition(ctx context.Context, latitude, longitude float64, radiusKm *float64) (locations []*models.Location, err error)
}

type redisLocationRepositoryImpl struct {
	rdb *redis.Client
}

func NewRedisLocationRepository(rdb *redis.Client) RedisLocationRepository {
	return &redisLocationRepositoryImpl{rdb}
}

func (r *redisLocationRepositoryImpl) PutLocation(ctx context.Context, location models.Location) (success bool, err error) {
	success = true
	cmd := r.rdb.GeoAdd(ctx, "users", &redis.GeoLocation{
		Name:      location.Username,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	})
	err = cmd.Err()
	success = err == nil
	return
}

func (r *redisLocationRepositoryImpl) GetLocationsByPosition(
	ctx context.Context,
	latitude,
	longitude float64,
	radiusKm *float64,
) (locations []*models.Location, err error) {
	locations = []*models.Location{}
	radius := 50.0
	if radiusKm != nil {
		radius = float64(*radiusKm)
	}
	res := r.rdb.GeoSearchLocation(ctx, "users", &redis.GeoSearchLocationQuery{
		GeoSearchQuery: redis.GeoSearchQuery{
			Latitude:   latitude,
			Longitude:  longitude,
			Radius:     radius,
			RadiusUnit: "km",
		},
		WithCoord: true,
		WithDist:  true, //
	})
	err = res.Err()
	locationsData, err := res.Result()
	if err != nil {
		return locations, err
	}
	for _, locationData := range locationsData {
		loc := models.Location{}
		loc.Latitude = locationData.Latitude
		loc.Longitude = locationData.Longitude
		loc.Username = locationData.Name
		locations = append(locations, &loc)
	}
	return
}
