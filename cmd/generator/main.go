package main

import (
	"flag"
	"log"

	"github.com/pester18/sequence-analyser/pkg/storage"
)

const (
	sequenceDefaultFile   = "./data/sequence.txt"
	generatorDefaultFile  = "./data/generator.json"
	defaultSequenceLength = 1024
)

func main() {
	sequenceFileFlag := flag.String("sequenceFile", sequenceDefaultFile, "File where pregenerated sequence is stored")
	generatorFileFlag := flag.String("generatorFile", generatorDefaultFile, "File where created generator will be stored")
	sequenceLengthFlag := flag.Int("n", defaultSequenceLength, "Length of sequence to generate with given generator")

	flag.Parse()

	sequenceFilepath := *sequenceFileFlag

	generatorFilepath := *generatorFileFlag

	length := *sequenceLengthFlag

	sequenceRepo := storage.NewSequenceRepository(sequenceFilepath)

	generatorRepo := storage.NewGeneratorRepository(generatorFilepath)

	generator, err := generatorRepo.LoadGenerator()
	if err != nil {
		log.Fatal(err)
	}

	sequence := make([]uint8, length)

	for i := range sequence {
		sequence[i] = generator.Next()
	}

	err = sequenceRepo.SaveSequence(sequence)
	if err != nil {
		log.Fatal(err)
	}

	err = generatorRepo.SaveGenerator(generator)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sequence of length %d was generated and saved in %s successfully.\n", length, sequenceFilepath)
}
