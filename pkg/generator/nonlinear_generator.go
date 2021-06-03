package generator

import (
	"math/rand"
)

type NonlinearGenerator struct {
	state   []byte
	fnTable map[string]uint8
}

func NewNLSFR(initVector []byte, fnTable map[string]uint8) *NonlinearGenerator {
	g := NonlinearGenerator{
		state:   initVector,
		fnTable: fnTable,
	}

	return &g
}

func (g *NonlinearGenerator) Next() uint8 {
	input := string(g.state)

	output, ok := g.fnTable[input]

	if !ok {
		output = uint8(rand.Intn(2))

		g.fnTable[input] = output
	}

	outputChar := byte('0' + output)

	g.state = append(g.state[1:], outputChar)

	return output
}

func (g *NonlinearGenerator) State() []byte {
	stateCopy := make([]byte, len(g.state))

	copy(stateCopy, g.state)

	return stateCopy
}

func (g *NonlinearGenerator) Table() map[string]uint8 {
	table := make(map[string]uint8)

	for args, value := range g.fnTable {

		table[args] = value

	}

	return table
}
