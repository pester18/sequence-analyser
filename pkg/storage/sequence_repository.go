package storage

import (
	"io/ioutil"
	"log"
)

type SequenceRepository struct {
	filepath string
}

func NewSequenceRepository(filepath string) *SequenceRepository {
	repo := SequenceRepository{
		filepath: filepath,
	}

	return &repo
}

func (repo *SequenceRepository) SaveSequence(seq []uint8) error {
	seqBytes := make([]byte, len(seq))
	for i, d := range seq {
		seqBytes[i] = '0' + d
	}

	err := ioutil.WriteFile(repo.filepath, seq, 0664)
	if err != nil {
		log.Println("Unable to save sequence, reason: ", err.Error())
		return WriteErr
	}

	log.Println(string(seqBytes))

	return nil
}

func (repo *SequenceRepository) LoadSequence() ([]uint8, error) {
	seqBytes, err := ioutil.ReadFile(repo.filepath)
	if err != nil {
		log.Println("Unable to load sequence, reason: ", err.Error())
		return nil, ReadErr
	}

	log.Println(string(seqBytes))

	seq := make([]byte, len(seqBytes))
	for i, d := range seqBytes {
		seq[i] = d - '0'
	}

	return seq, nil
}
