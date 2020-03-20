package once

import (
	"time"
)

type Func func() error

func (fn Func) Every(interval time.Duration) error {
	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for range ticker.C {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}
