package trimmedmean_test

import (
	"math"
	"testing"
	"trimmedmean"
)

// closeEnough returns true when a and b differ by less than tol.
func closeEnough(a, b, tol float64) bool {
	return math.Abs(a-b) < tol
}

const eps = 1e-9

// ---------------------------------------------------------------------------
// Happy-path tests
// ---------------------------------------------------------------------------

func TestSymmetricTrimming_Float64(t *testing.T) {
	// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], trim 10% each end → drop 1 from each
	// Remaining: [2, 3, 4, 5, 6, 7, 8, 9] → mean = 5.5
	data := []float64{3, 1, 7, 5, 9, 2, 8, 4, 6, 10}
	got, err := trimmedmean.TrimmedMean(data, 0.1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 5.5
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

func TestSymmetricTrimming_Int(t *testing.T) {
	// Same data as above but with int slice.
	data := []int{3, 1, 7, 5, 9, 2, 8, 4, 6, 10}
	got, err := trimmedmean.TrimmedMean(data, 0.1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 5.5
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

func TestAsymmetricTrimming(t *testing.T) {
	// [1..10], trim 20% lo (drop 2), 10% hi (drop 1)
	// Remaining: [3, 4, 5, 6, 7, 8, 9] → mean = 6.0
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got, err := trimmedmean.TrimmedMean(data, 0.2, 0.1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 6.0
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

func TestAsymmetricTrimming_HeavyHigh(t *testing.T) {
	// [1..10], trim 0% lo, 30% hi (drop 3)
	// Remaining: [1, 2, 3, 4, 5, 6, 7] → mean = 4.0
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got, err := trimmedmean.TrimmedMean(data, 0.0, 0.3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 4.0
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

func TestZeroTrimming(t *testing.T) {
	// trim=0 → plain arithmetic mean
	data := []float64{1, 2, 3, 4, 5}
	got, err := trimmedmean.TrimmedMean(data, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 3.0
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

func TestZeroTrimming_Asymmetric(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	got, err := trimmedmean.TrimmedMean(data, 0.0, 0.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 3.0
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

func TestSingleElement(t *testing.T) {
	data := []float64{42.0}
	got, err := trimmedmean.TrimmedMean(data, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !closeEnough(got, 42.0, eps) {
		t.Errorf("got %.10f, want 42.0", got)
	}
}

func TestNegativeValues(t *testing.T) {
	// [-5, -4, -3, -2, -1, 0, 1, 2, 3, 4], trim 20% each → drop 2 each end
	// Remaining: [-3, -2, -1, 0, 1, 2] → mean = -0.5
	data := []float64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4}
	got, err := trimmedmean.TrimmedMean(data, 0.2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := -0.5
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

func TestFloat32Slice(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got, err := trimmedmean.TrimmedMean(data, 0.1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 5.5
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

func TestInt32Slice(t *testing.T) {
	data := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got, err := trimmedmean.TrimmedMean(data, 0.1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 5.5
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

// Verify the original slice is not mutated.
func TestOriginalSliceUnchanged(t *testing.T) {
	data := []float64{5, 3, 1, 4, 2}
	original := make([]float64, len(data))
	copy(original, data)

	_, err := trimmedmean.TrimmedMean(data, 0.2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for i := range data {
		if data[i] != original[i] {
			t.Errorf("original slice modified at index %d: got %v, want %v",
				i, data[i], original[i])
		}
	}
}

// Fractional trimming: with 7 elements and trim=0.1, floor(7*0.1)=0 elements
// dropped from each end → plain mean.
func TestFractionalTrimDropsZero(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7}
	got, err := trimmedmean.TrimmedMean(data, 0.1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 4.0
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

// Verify behaviour matches R's mean(x, trim=0.25) on a known dataset.
// R: mean(c(2,4,6,8,10,12,14,16,18,20), trim=0.25)
// → sorted: [2,4,6,8,10,12,14,16,18,20], drop 2 each → [6,8,10,12,14,16] → mean=11
func TestMatchesR(t *testing.T) {
	data := []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	got, err := trimmedmean.TrimmedMean(data, 0.25)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 11.0
	if !closeEnough(got, want, eps) {
		t.Errorf("got %.10f, want %.10f", got, want)
	}
}

// ---------------------------------------------------------------------------
// Error-path tests
// ---------------------------------------------------------------------------

func TestErrorEmptySlice(t *testing.T) {
	_, err := trimmedmean.TrimmedMean([]float64{}, 0.1)
	if err == nil {
		t.Fatal("expected error for empty slice, got nil")
	}
}

func TestErrorNoTrimArgs(t *testing.T) {
	_, err := trimmedmean.TrimmedMean([]float64{1, 2, 3})
	if err == nil {
		t.Fatal("expected error for zero trimming arguments, got nil")
	}
}

func TestErrorTooManyTrimArgs(t *testing.T) {
	_, err := trimmedmean.TrimmedMean([]float64{1, 2, 3}, 0.1, 0.1, 0.1)
	if err == nil {
		t.Fatal("expected error for three trimming arguments, got nil")
	}
}

func TestErrorNegativeProportion(t *testing.T) {
	_, err := trimmedmean.TrimmedMean([]float64{1, 2, 3}, -0.1)
	if err == nil {
		t.Fatal("expected error for negative proportion, got nil")
	}
}

func TestErrorProportionEqualToOne(t *testing.T) {
	_, err := trimmedmean.TrimmedMean([]float64{1, 2, 3}, 1.0)
	if err == nil {
		t.Fatal("expected error for proportion == 1.0, got nil")
	}
}

func TestErrorProportionGreaterThanOne(t *testing.T) {
	_, err := trimmedmean.TrimmedMean([]float64{1, 2, 3}, 1.5)
	if err == nil {
		t.Fatal("expected error for proportion > 1.0, got nil")
	}
}

func TestErrorCombinedProportionExceedsOne(t *testing.T) {
	_, err := trimmedmean.TrimmedMean([]float64{1, 2, 3}, 0.6, 0.5)
	if err == nil {
		t.Fatal("expected error for lo+hi >= 1.0, got nil")
	}
}

func TestErrorNaNProportion(t *testing.T) {
	_, err := trimmedmean.TrimmedMean([]float64{1, 2, 3}, math.NaN())
	if err == nil {
		t.Fatal("expected error for NaN proportion, got nil")
	}
}
