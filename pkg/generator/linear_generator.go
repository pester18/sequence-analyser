package generator

type LinearGenerator struct {
	state   []uint8
	polynom []uint8
}

func NewLSFR(initVector []uint8, polynom []uint8) *LinearGenerator {
	g := LinearGenerator{
		state:   initVector,
		polynom: polynom,
	}

	return &g
}

func (g *LinearGenerator) Next() uint8 {
	var output uint8

	for i := 0; i < len(g.state); i++ {

		output ^= g.state[i] * g.polynom[i]

	}

	g.state = append(g.state[1:], output)

	return output
}

func (g *LinearGenerator) State() []uint8 {
	stateCopy := make([]uint8, len(g.state))

	copy(stateCopy, g.state)

	return stateCopy
}

func (g *LinearGenerator) Polynom() []uint8 {
	polynomCopy := make([]uint8, len(g.polynom))

	copy(polynomCopy, g.polynom)

	return polynomCopy
}
