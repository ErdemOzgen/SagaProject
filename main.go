package main

import (
	"fmt"
	"sagaAlienInvasion/banner"
	"sagaAlienInvasion/cmd"
	"sagaAlienInvasion/utils"

	log "github.com/sirupsen/logrus"
)

func init() {
	fmt.Println("INIT MAIN")
	utils.SetLogLevel(cmd.Environment)
	log.Info("Running environment==>", cmd.Environment)
}
func main() {
	log.Info("Simulator has been started...")
	defer log.Info("Simulator has been finished...")
	banner.PrintBanner()
	//fmt.Println("SimulatorJSKJAKLDJLAJKLDJSL", banner.HashForSumBanner())
	cmd.Execute()

	//utils.GenerateAlienNamesAndWriteToFile(1000)
}
