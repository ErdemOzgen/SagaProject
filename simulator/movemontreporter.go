package simulator

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type MovReport uint8

type MovError struct {
	reason MovReport
}

const (
	// MovAlienDead when Alien is Dead
	MovAlienDead MovReport = iota
	// MovAlienTrapped when Alien is Trapped
	MovAlienTrapped
	// MovWorldDestroyed when World is destroyed
	MovWorldDestroyed
	// MovMessage when no-op
	MovMessage = " cant move needs to invastigate %s\n"
)

func (err *MovError) Error() string {
	log.Info("Allien movement need to invastigate ")
	return fmt.Sprintf("Simulator movement error %d", err.reason)
}

// checkAlien returns NoOpError if Alien dead or trapped
func checkAlien(alien *Alien) *MovError {
	if alien.IsDead() {
		return &MovError{MovAlienDead}
	}
	if alien.IsTrapped() {
		return &MovError{MovAlienTrapped}
	}
	return nil
}
