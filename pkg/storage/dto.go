package storage

import (
	"encoding/json"
	"log"

	"github.com/pester18/sequence-analyser/pkg/generator"
)

type GeneratorType int

const (
	_ GeneratorType = iota
	LinearGeneratorType
	NonlinearGeneratorType
)

type GeneratorDTO struct {
	Type          GeneratorType `json:"type"`
	GeneratorBody []byte        `json:"generator"`
}

type NonlinearGeneratorDTO struct {
	State   []byte           `json:"state"`
	FnTable map[string]uint8 `json:"fn_table"`
}

type LinearGeneratorDTO struct {
	State   []uint8 `json:"state"`
	Polynom []uint8 `json:"polynom"`
}

// Serializing

func SerializeGenerator(gen generator.Generator) (body []byte, err error) {
	var genBody []byte
	var genType GeneratorType

	switch g := gen.(type) {
	case *generator.LinearGenerator:
		genType = LinearGeneratorType
		genBody, err = serializeLinearGenerator(g)

	case *generator.NonlinearGenerator:
		genType = NonlinearGeneratorType
		genBody, err = serializeNonlinearGenerator(g)
	}

	if err != nil {
		return nil, err
	}

	dto := GeneratorDTO{
		Type:          genType,
		GeneratorBody: genBody,
	}

	body, err = json.Marshal(dto)
	if err != nil {
		log.Println("Unable serialize generator, reason: ", err.Error())
		return nil, err
	}

	return body, nil
}

func serializeLinearGenerator(gen *generator.LinearGenerator) (body []byte, err error) {
	dto := LinearGeneratorDTO{
		State:   gen.State(),
		Polynom: gen.Polynom(),
	}

	body, err = json.Marshal(dto)
	if err != nil {
		log.Println("Unable serialize linear generator, reason: ", err.Error())
		return nil, err
	}

	return body, nil
}

func serializeNonlinearGenerator(gen *generator.NonlinearGenerator) (body []byte, err error) {
	dto := NonlinearGeneratorDTO{
		State:   gen.State(),
		FnTable: gen.Table(),
	}

	body, err = json.Marshal(dto)
	if err != nil {
		log.Println("Unable serialize nonlinear generator, reason: ", err.Error())
		return nil, err
	}

	return body, nil
}

// Deserializing

func DeserializeGenerator(body []byte) (gen generator.Generator, err error) {
	var generatorDTO GeneratorDTO

	err = json.Unmarshal(body, &generatorDTO)
	if err != nil {
		log.Println("Unable deserialize generator, reason: ", err.Error())
		return nil, err
	}

	switch generatorDTO.Type {
	case LinearGeneratorType:
		gen, err = deserializeLinearGenerator(generatorDTO.GeneratorBody)
	case NonlinearGeneratorType:
		gen, err = deserializeNonlinearGenerator(generatorDTO.GeneratorBody)
	}

	if err != nil {
		return nil, err
	}

	return gen, nil
}

func deserializeLinearGenerator(body []byte) (gen *generator.LinearGenerator, err error) {
	var generatorDTO LinearGeneratorDTO

	err = json.Unmarshal(body, &generatorDTO)
	if err != nil {
		log.Println("Unable deserialize linear generator, reason: ", err.Error())
		return nil, err
	}

	gen = generator.NewLSFR(generatorDTO.State, generatorDTO.Polynom)

	return gen, nil
}

func deserializeNonlinearGenerator(body []byte) (gen *generator.NonlinearGenerator, err error) {
	var generatorDTO NonlinearGeneratorDTO

	err = json.Unmarshal(body, &generatorDTO)
	if err != nil {
		log.Println("Unable deserialize linear generator, reason: ", err.Error())
		return nil, err
	}

	gen = generator.NewNLSFR([]byte(generatorDTO.State), generatorDTO.FnTable)

	return gen, nil
}
