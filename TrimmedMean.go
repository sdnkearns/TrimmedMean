package trimmedmean

import (
	"fmt"
	"math"
	"sort"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func TrimmedMean[T Number](src []T, trimming ...float64) (float64, error) {
	n := len(src)
	if n == 0 {
		return 0, fmt.Errorf("TrimmedMean: slice must not be empty")
	}

	var lo, hi float64
	switch len(trimming) {
	case 1:
		lo, hi = trimming[0], trimming[0]
	case 2:
		lo, hi = trimming[0], trimming[1]
	default:
		return 0, fmt.Errorf(
			"TrimmedMean: expected 1 or 2 trimming proportions, got %d", len(trimming))
	}

	if err := validateProportion("low", lo); err != nil {
		return 0, err
	}
	if err := validateProportion("high", hi); err != nil {
		return 0, err
	}
	if lo+hi >= 1.0 {
		return 0, fmt.Errorf(
			"TrimmedMean: combined trimming proportions (%.6g + %.6g = %.6g) must be < 1",
			lo, hi, lo+hi)
	}

	data := make([]float64, n)
	for i, v := range src {
		data[i] = float64(v)
	}
	sort.Float64s(data)

	lowTrim := int(math.Floor(float64(n) * lo))
	highTrim := int(math.Floor(float64(n) * hi))

	trimmed := data[lowTrim : n-highTrim]
	if len(trimmed) == 0 {
		return 0, fmt.Errorf("TrimmedMean: slice empty after trimming")
	}

	var sum float64
	for _, v := range trimmed {
		sum += v
	}
	mean := sum / float64(len(trimmed))

	if math.IsNaN(mean) || math.IsInf(mean, 0) {
		return 0, fmt.Errorf(
			"TrimmedMean: computed mean is not finite (%v); check for NaN/Inf in src", mean)
	}
	return mean, nil
}

func validateProportion(label string, p float64) error {
	if math.IsNaN(p) || p < 0 || p >= 1 {
		return fmt.Errorf(
			"trimmean: %s trimming proportion must be in [0, 1), got %v", label, p)
	}
	return nil
}
