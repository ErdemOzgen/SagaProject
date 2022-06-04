package cmd

import (
	"errors"
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

func ParseFlag() {
	if Environment == "development" {
		flag.Int64Var(&entropy, "entropy", 0, "random number used as entropy seed")
		simulationName = "simulation"
	}
	flag.StringVar(&simulationName, "sn", simulationName, "simulation name (SN)")
	flag.IntVar(&movements, "m", AlienMovement, "Movements count to be made by the alien (M)")
	flag.IntVar(&alienNumber, "c", DefaultNumberOfAliens, "Count of aliens to be created (C)")
	flag.StringVar(&worldFile, "w", DefaultWorldFile, "Input file containing the world map (W)")
	flag.StringVar(&alienNameFile, "n", DefaultAlienNameFile, "File containing alien names (N)")
	flag.Parse()
}

// checkFlags validates input flags
func CheckFlags() error {
	if alienNumber <= 0 {
		log.Warn("alien number less than or equal to zero parameter given")
		return errors.New("alien number must be greater than 0")

	}
	if alienNumber > DefaultAlienNameFileLen {
		//utils.GenerateAlienNamesAndWriteToFile(alienNumber)
		// TODO: change default allien name file to a generated file
		log.Warn("alien number greater than default file length")
	}
	if movements <= 0 {
		log.Warn("movements less than or equal to zero parameter given")
		return errors.New("movements must be greater than 0")
	}
	if len(worldFile) == 0 {
		log.Warn("world file parameter given as empty string")
		return errors.New("can not find world file")
	}
	return nil
}

func ControlCheckedFlags() error {
	if err := CheckFlags(); err != nil {
		log.Info("Error while checking flags: %s\n", err)
		flag.Usage()
		//log.Fatal("Error while checking flags closing simulator")
		os.Exit(1)
	}
	return nil
}
