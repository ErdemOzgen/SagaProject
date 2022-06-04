package cmd

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"sagaAlienInvasion/banner"
	"sagaAlienInvasion/simulator"
	"sagaAlienInvasion/utils"
	"time"
)

var (
	entropy                                  int64
	movements, alienNumber                   int
	simulationName, worldFile, alienNameFile string
)

func init() {
	fmt.Println("INIT CMD")
	// flag.Int64Var(&entropy, "entropy", 0, "random number used as entropy seed")
	ParseFlag()

}
func Execute() {
	ControlCheckedFlags()
	world, in, err := simulator.ReadWorldMapFile(worldFile)
	if err != nil {
		fmt.Printf("Cant find map file \"%s\" with error: %s\n", worldFile, err)
		os.Exit(1)
	}
	// Building the invasion simulation
	r := buildRand()
	aliens := simulator.RandAliens(alienNumber, r)
	if alienNameFile != "" {
		if err := simulator.IdentifyAliens(aliens, alienNameFile); err != nil {
			fmt.Printf("Cant find Allien Name file\"%s\" with error: %s\n", alienNameFile, err)
			os.Exit(1)
		}
	}
	sim := simulator.NewSimulator(r, movements, world, aliens)
	// Start the simulation and print any errors
	if err := sim.Start(); err != nil {
		fmt.Printf(utils.PrintPretty("Sim Error: %s"), err)
		os.Exit(1)
	}
	// Finish Sim
	banner.PrintSimSumBanner()
	fmt.Printf(utils.PrintPretty("Sim Completed Showing Cities left"))
	fmt.Print(in.FilterDestroyed(world))

}

// buildRand build a pseudorandom numbers generator from input flags
// for development env set entropy to 0
func buildRand() *rand.Rand {
	var seed int64
	var source rand.Source
	if entropy != 0 {
		seed = entropy
		source = rand.NewSource(entropy)

	} else if len(simulationName) > 0 {
		hash := sha256.Sum256([]byte(simulationName))
		seed = int64(binary.BigEndian.Uint64(hash[:8]))
		source = rand.NewSource(seed)

	} else {
		seed = time.Now().UnixNano()
		source = rand.NewSource(seed)
	}
	return rand.New(source)
}
