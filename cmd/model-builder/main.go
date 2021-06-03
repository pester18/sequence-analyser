package main

import (
	"flag"
	"log"

	"github.com/pester18/sequence-analyser/pkg/analysis"
	"github.com/pester18/sequence-analyser/pkg/generator"
	"github.com/pester18/sequence-analyser/pkg/storage"
)

const (
	linearModelType      = "linear"
	nonlinearModelType   = "nonlinear"
	sequenceDefaultFile  = "./data/sequence.txt"
	generatorDefaultFile = "./data/generator.json"
	defaultFaultRate     = 0.05
)

func main() {
	sequenceFileFlag := flag.String("sequenceFile", sequenceDefaultFile, "File where pregenerated sequence is stored")
	generatorFileFlag := flag.String("generatorFile", generatorDefaultFile, "File where created generator will be stored")
	reproducingModelTypeFlag := flag.String("modelType", nonlinearModelType, "Type of sequence extrapolating model")
	faultRateFlag := flag.Float64("faultRate", defaultFaultRate, "Fault rate for nonlinear reproducing model")

	flag.Parse()

	sequenceFilepath := *sequenceFileFlag

	generatorFilepath := *generatorFileFlag

	modelType := *reproducingModelTypeFlag

	faultRate := *faultRateFlag

	sequenceRepo := storage.NewSequenceRepository(sequenceFilepath)

	generatorRepo := storage.NewGeneratorRepository(generatorFilepath)

	sequence, err := sequenceRepo.LoadSequence()
	if err != nil {
		log.Fatal(err)
	}

	sequenceExtrapolator := analysis.NewSequenceExtrapolator(sequence)

	var generator generator.Generator

	switch modelType {
	case nonlinearModelType:
		generator = sequenceExtrapolator.NonlinearModel(faultRate)
	case linearModelType:
		generator = sequenceExtrapolator.LinearModel()
	default:
		log.Fatalf("%s sequence reproducing model is not supported", modelType)
	}

	err = generatorRepo.SaveGenerator(generator)
	if err != nil {
		log.Fatal(err)
	}
}
