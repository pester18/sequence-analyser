package main

import (
	"log"
	"math/rand"

	"github.com/pester18/sequence-analyser/pkg/storage"
)

const (
	sequenceDefaultFile   = "./data/sequence"
	defaultSequenceLength = 1024
)

func main() {

	sequenceFilepath := sequenceDefaultFile
	length := defaultSequenceLength

	//
	sequence := make([]uint8, defaultSequenceLength)

	for n := 0; n < length; n++ {
		d := rand.Intn(2)
		sequence[n] = uint8(d)
	}

	sequenceRepo := storage.NewSequenceRepository(sequenceFilepath)

	err := sequenceRepo.SaveSequence(sequence)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sequence of length %d was generated and saved in %s successfully.\n", length, sequenceFilepath)
}
