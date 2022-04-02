package mascot_test

import (
	"testing"

	"github.com/djmin43/learning-go/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "austin leee" {
		t.Fatal("wrong mascot")
	}
}