package factories

import (
	"fmt"
	"kumande/configs"
	"kumande/models"
	"math/rand"
	"time"
)

func UserTrackFactory() models.UserTrack {
	rand.Seed(time.Now().UnixNano())

	lat := -90 + rand.Float64()*180
	long := -180 + rand.Float64()*360
	trackSource := configs.TrackSources[rand.Intn(len(configs.TrackSources))]

	return models.UserTrack{
		TrackLat:    fmt.Sprintf("%f", lat),
		TrackLong:   fmt.Sprintf("%f", long),
		TrackSource: trackSource,
	}
}
