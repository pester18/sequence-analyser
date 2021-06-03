package analysis

import (
	"fmt"
	"math"
)

func FrequencyAnalysis(sequence []uint8) (rate float32) {
	var sum float32

	length := len(sequence)

	for i := 0; i < length; i++ {

		sum += float32(sequence[i])

	}

	rate = sum / float32(length)

	return rate
}

func DifferentialAnalysis(sequence []uint8) (rate float32) {
	length := len(sequence)

	var diffSum float32

	for i := 1; i < length; i++ {

		current := sequence[i]

		previous := sequence[i-1]

		diff := current ^ previous

		diffSum += float32(diff)

	}

	rate = diffSum / float32(length-1)

	return rate
}

func LinearAnalysis(sequence []uint8) (complexity int) {
	length := len(sequence)

	BD := make([]uint8, 1)
	CD := make([]uint8, length)

	CD[0] = 1
	BD[0] = 1

	L := 0
	x := 1

	for n := 0; n < length; n++ {
		d := sequence[n]

		for j := 1; j <= L; j++ {

			d ^= CD[j] * sequence[n-j]

		}

		if d == 1 {

			if 2*L > n {

				for i, bd := range BD {

					CD[x+i] ^= bd

				}

			} else {

				copyCD := make([]uint8, L+1)

				copy(copyCD, CD)

				L = n - L + 1

				for i, bd := range BD {

					CD[x+i] ^= bd

				}

				BD = copyCD

				x = 0
			}

		}

		x += 1
	}

	complexity = L

	return complexity
}

func NonlinearAnalysis(sequence []uint8) int {
	length := len(sequence)

	var complexity int

	minRegisterSize := int(math.Floor(math.Log2(float64(length)) / 2))

	maxRegisterSize := int(math.Ceil(math.Log2(float64(length)) * 2.5))

	sequenceBytes := make([]byte, length)

	for i := 0; i < length; i++ {

		sequenceBytes[i] = '0' + sequence[i]

	}

	for registerSize := minRegisterSize; registerSize <= maxRegisterSize; registerSize++ {
		mistakeDone := false

		fnTable := make(map[string]uint8)

		for i := 0; i < length-registerSize; i++ {

			key := string(sequenceBytes[i : i+registerSize])

			value := sequence[i+registerSize]

			if _, ok := fnTable[key]; !ok {

				fnTable[key] = value

			} else if fnTable[key] != value {

				mistakeDone = true

				break

			}

		}

		if !mistakeDone {

			return registerSize

		}

	}

	return complexity
}

func RankAnalysis(
	generatedValues []uint8,
	minWindowSize int,
	maxWindowSize int,
	accuracy float64,
) int {
	fmt.Println("Rank analysis: ")

	length := len(generatedValues)

	generatedValuesBytes := make([]byte, length)

	for i := 0; i < length; i++ {

		generatedValuesBytes[i] = '0' + generatedValues[i]

	}

	for windowSize := minWindowSize; windowSize < maxWindowSize; windowSize++ {

		combinations := make(map[string]int)

		combinationsCount := length - windowSize + 1

		for i := 0; i < combinationsCount; i++ {

			value := string(generatedValuesBytes[i : i+windowSize])

			if _, ok := combinations[value]; !ok {

				combinations[value] = 1

			} else {

				combinations[value]++

			}
		}

		averageCount := float64(combinationsCount) / math.Pow(2.0, float64(windowSize))

		for combination, count := range combinations {

			deviation := math.Abs((float64(count) - averageCount) / float64(length))

			if math.Abs(deviation) <= accuracy {

				fmt.Printf(
					"Window size: %v; Combination: %v; Percent: %v;\n",
					windowSize,
					combination,
					deviation*100,
				)

				return windowSize

			}

		}

	}
	return 0
}
