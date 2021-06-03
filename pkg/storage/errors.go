package storage

import (
	"errors"
)

var SerializationErr = errors.New("Error during serialization")
var WriteErr = errors.New("Error during write")
var ReadErr = errors.New("Error during read")
