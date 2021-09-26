package reqstash

import "time"

type Config struct {
	Storage struct {
		Capacity int           `envconfig:"capacity"`
		TTL      time.Duration `envconfig:"ttl"`
	} `envconfig:"storage"`
}
