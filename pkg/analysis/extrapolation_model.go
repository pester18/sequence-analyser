package analysis

import (
	"math"

	"github.com/pester18/sequence-analyser/pkg/generator"
)

type SequenceExtrapolator struct {
	sequence []uint8
}

func NewSequenceExtrapolator(sequence []uint8) *SequenceExtrapolator {
	extrapolator := SequenceExtrapolator{
		sequence: sequence,
	}

	return &extrapolator
}

func (e *SequenceExtrapolator) LinearModel() *generator.LinearGenerator {
	length := len(e.sequence)

	auxiliaryCoefs := make([]uint8, 1)

	coefficients := make([]uint8, length)

	coefficients[0] = 1

	auxiliaryCoefs[0] = 1

	complexity := 0

	offset := 1

	for n := 0; n < length; n++ {
		d := e.sequence[n]

		for j := 1; j <= complexity; j++ {

			d ^= coefficients[j] * e.sequence[n-j]

		}

		if d == 1 {

			if 2*complexity > n {

				for i, bd := range auxiliaryCoefs {

					coefficients[offset+i] ^= bd

				}

			} else {

				copyCoefficients := make([]uint8, complexity+1)

				copy(copyCoefficients, coefficients)

				complexity = n - complexity + 1

				for i, coef := range auxiliaryCoefs {

					coefficients[offset+i] ^= coef

				}

				auxiliaryCoefs = copyCoefficients

				offset = 0
			}

		}

		offset += 1
	}

	initVector := make([]uint8, complexity)
	polynomCoeffs := make([]uint8, complexity)

	copy(initVector, e.sequence[length-complexity:])
	copy(polynomCoeffs, coefficients[1:complexity+1])

	return generator.NewLSFR(initVector, polynomCoeffs)
}

func (e *SequenceExtrapolator) NonlinearModel(faultRate float64) *generator.NonlinearGenerator {
	var complexity int
	var intermediateFnTable map[string][]int

	length := len(e.sequence)

	minRegisterSize := int(math.Floor(math.Log2(float64(length)) / 2))

	maxRegisterSize := int(math.Ceil(math.Log2(float64(length)) * 4))

	sequenceBytes := make([]byte, length)

	for i := 0; i < length; i++ {

		sequenceBytes[i] = '0' + e.sequence[i]

	}

	for registerSize := minRegisterSize; registerSize <= maxRegisterSize; registerSize++ {

		complexity = registerSize

		intermediateFnTable = make(map[string][]int)

		for i := 0; i < length-registerSize; i++ {

			key := string(sequenceBytes[i : i+registerSize])

			value := e.sequence[i+registerSize]

			if _, ok := intermediateFnTable[key]; !ok {

				intermediateFnTable[key] = []int{0, 0}

			}

			counter := intermediateFnTable[key]

			counter[int(value)]++

			intermediateFnTable[key] = counter

		}

		var mistakesNum float64

		for _, count := range intermediateFnTable {

			mistakes := math.Min(float64(count[0]), float64(count[1]))

			mistakesNum += mistakes

		}

		mistakesRatio := mistakesNum / float64(length-registerSize)

		if mistakesRatio <= faultRate {
			break
		}

	}

	fnTable := make(map[string]uint8)

	for args, count := range intermediateFnTable {

		if count[0] > count[1] {
			fnTable[args] = 0
		} else {
			fnTable[args] = 1
		}

	}

	currentState := sequenceBytes[length-complexity:]

	return generator.NewNLSFR(currentState, fnTable)
}
