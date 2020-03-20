package once

import (
	"time"
)

type Once func() error

func (fn Once) Every(interval time.Duration) error {
	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for range ticker.C {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}
