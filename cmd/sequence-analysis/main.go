package main

import (
	"flag"
	"log"

	"github.com/pester18/sequence-analyser/pkg/analysis"
	"github.com/pester18/sequence-analyser/pkg/storage"
)

const (
	sequenceDefaultFile = "./data/sequence.txt"
)

func main() {
	sequenceFileFlag := flag.String("sequenceFile", sequenceDefaultFile, "File where previously generated sequence is stored")

	flag.Parse()

	sequenceFilepath := *sequenceFileFlag

	sequenceRepo := storage.NewSequenceRepository(sequenceFilepath)

	sequence, err := sequenceRepo.LoadSequence()
	if err != nil {
		log.Fatal(err)
	}

	frequencyCriteria := analysis.FrequencyAnalysis(sequence)
	differentialCriteria := analysis.DifferentialAnalysis(sequence)
	linearComplexity := analysis.LinearAnalysis(sequence)
	nonlinearComplexity := analysis.NonlinearAnalysis(sequence)

	log.Println("Frequency criteria of sequence is: ", frequencyCriteria)
	log.Println("Differential criteria of sequence is: ", differentialCriteria)
	log.Println("Linear complexity of sequence is: ", linearComplexity)
	log.Println("Nonlinear complexity of sequence is: ", nonlinearComplexity)
}
