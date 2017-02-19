package cuberl

import (
	"math"
	"testing"

	"github.com/unixpickle/anyvec/anyvec32"
)

func TestCorrectedPred(t *testing.T) {
	in := anyvec32.MakeVectorData([]float32{
		0.163416, 0.072295, 0.197626, 0.843745, 0.141560, 0.629271, 0.383674, 0.285430, 0.078506, 0.288674, 0.264813, 0.426436, 0.132217, 0.235223, 0.804233, 0.440210, 0.037329, 0.659419,
		0.058389, 0.354383, 0.458381, 0.389175, 0.198247, 0.553172, 0.627449, 0.165429, 0.104431, 0.987851, 0.484646, 0.658479, 0.453316, 0.440728, 0.950924, 0.225095, 0.201714, 0.957615,
	})
	next := anyvec32.MakeVectorData([]float32{
		0.229171, 0.257058, 0.435237, 0.751883, 0.924023, 0.933591, 0.369723, 0.329040, 0.199867, 0.473958, 0.905251, 0.263628, 0.438647, 0.850818, 0.722620, 0.998309, 0.460241, 0.555185,
		0.645400, 0.365332, 0.933181, 0.283440, 0.597614, 0.941348, 0.152524, 0.282988, 0.513079, 0.061375, 0.078637, 0.298268, 0.362312, 0.750441, 0.206781, 0.359686, 0.119473, 0.948106,
	})
	chosen := []int{5, 16}
	actual := correctedPred(in, next, chosen, []float64{1, -2}).Data().([]float32)
	expected := []float32{
		0.163416, 0.072295, 0.197626, 0.843745, 0.141560, 1 + 0.99831, 0.383674, 0.285430, 0.078506, 0.288674, 0.264813, 0.426436, 0.132217, 0.235223, 0.804233, 0.440210, 0.037329, 0.659419,
		0.058389, 0.354383, 0.458381, 0.389175, 0.198247, 0.553172, 0.627449, 0.165429, 0.104431, 0.987851, 0.484646, 0.658479, 0.453316, 0.440728, 0.950924, 0.225095, -2 + 0.94811, 0.957615,
	}
	if len(actual) != len(expected) {
		t.Fatal("length mismatch")
	}
	for i, x := range expected {
		a := actual[i]
		if math.Abs(float64(a-x)) > 1e-4 {
			t.Errorf("idx %d: expected %v but got %v", i, x, a)
		}
	}
}
