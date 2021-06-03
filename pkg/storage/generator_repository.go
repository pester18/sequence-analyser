package storage

import (
	"io/ioutil"
	"log"

	"github.com/pester18/sequence-analyser/pkg/generator"
)

type GeneratorRepository struct {
	filepath string
}

func NewGeneratorRepository(filepath string) *GeneratorRepository {
	repo := GeneratorRepository{
		filepath: filepath,
	}

	return &repo
}

func (repo *GeneratorRepository) SaveGenerator(gen generator.Generator) error {
	data, err := SerializeGenerator(gen)
	if err != nil {
		return SerializationErr
	}

	err = ioutil.WriteFile(repo.filepath, data, 0664)
	if err != nil {
		log.Println("Unable to save generator, reason: ", err.Error())
		return WriteErr
	}

	return nil
}

func (repo *GeneratorRepository) LoadGenerator() (generator.Generator, error) {
	data, err := ioutil.ReadFile(repo.filepath)
	if err != nil {
		log.Println("Unable to load generator, reason: ", err.Error())
		return nil, ReadErr
	}

	gen, err := DeserializeGenerator(data)
	if err != nil {
		return nil, SerializationErr
	}

	return gen, nil
}
