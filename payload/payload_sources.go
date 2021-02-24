package payload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/bitrise-io/bitrise/models"
)

var errNoInput = errors.New("nothing to read")

// EnvPayloadSource ...
type EnvPayloadSource struct {
	envValue string
}

// Payload ...
func (s EnvPayloadSource) Payload() (models.BuildRunResultsModel, error) {
	if s.envValue == "" {
		return models.BuildRunResultsModel{}, errNoInput
	}

	var payload models.BuildRunResultsModel
	if err := json.Unmarshal([]byte(s.envValue), &payload); err != nil {
		return models.BuildRunResultsModel{}, err
	}
	return payload, nil
}

// StdinPayloadSource ....
type StdinPayloadSource struct {
	reader io.Reader
}

// Payload ....
func (s StdinPayloadSource) Payload() (models.BuildRunResultsModel, error) {
	b, err := read(s.reader)
	if err != nil {
		return models.BuildRunResultsModel{}, err
	}
	if len(b) == 0 {
		return models.BuildRunResultsModel{}, errNoInput
	}

	var buildRunResults models.BuildRunResultsModel
	if err := json.Unmarshal(b, &buildRunResults); err != nil {
		return models.BuildRunResultsModel{}, fmt.Errorf("failed to parse plugin input (%s): %s", string(b), err)
	}
	return buildRunResults, nil
}

func read(r io.Reader) ([]byte, error) {
	var buff []byte
	for {
		chunk := make([]byte, 100)
		n, err := r.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if n == 0 {
			break
		}
		buff = append(buff, chunk[:n]...)
	}
	return buff, nil
}

// HasStat ...
type HasStat interface {
	Stat() (os.FileInfo, error)
}

// HasContent ...
func HasContent(f HasStat) (bool, error) {
	fi, err := f.Stat()
	if err != nil {
		return false, err
	}

	if fi.Mode()&os.ModeNamedPipe != 0 {
		return true, nil
	}

	return fi.Size() > 0, nil
}
