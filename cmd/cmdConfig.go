package cmd

const (
	// AlienMovement used if number of moves is not given
	AlienMovement int = 10000

	// DefaultNumberOfAliens used if number of Aliens is not otherwise specified
	DefaultNumberOfAliens int = 5

	// DefaultWorldFile used if World file is not otherwise specified
	DefaultWorldFile = "./files/example.txt"

	// DefaultalienNameFile used if intel file is not otherwise specified
	DefaultAlienNameFile    = "./files/aliensNames.txt"
	DefaultAlienNameFileLen = 1000

	// Environment used to determine if the program is running in development or production mode
	Environment = "development"
)
