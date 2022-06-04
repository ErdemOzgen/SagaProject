package banner

import (
	"testing"
)

func TestPrintBanner(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"Control ascii art hash"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var want uint32 = 754652247
			if got := HashForASCII(AsciiArt); got != want {
				t.Errorf("PrintBanner() = %v, want %v", got, want)
			}
			t.Log("PrintBanner() hash is correct")
		})
	}
}

func TestPrintSimSumBanner(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"Control Simulation sum banner hash"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var want uint32 = 364400229
			if got := HashForASCII(SimulationSumArt); got != want {
				t.Errorf("PrintSimSumBanner() = %v, want %v", got, want)
			}
		})
	}
}
