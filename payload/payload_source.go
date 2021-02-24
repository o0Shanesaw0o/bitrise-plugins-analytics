package payload

import (
	"os"

	"github.com/bitrise-io/bitrise-plugins-analytics/configs"
	"github.com/bitrise-io/bitrise/models"
)

// Source ...
type Source interface {
	Payload() (models.BuildRunResultsModel, error)
}

// SourceType ...
type SourceType int

// SourceTypes ...
const (
	StdinSource SourceType = iota
	EnvSource
)

// SourceFactory ....
func SourceFactory(t SourceType) Source {
	if t == StdinSource {
		return StdinPayloadSource{os.Stdin}
	}
	return EnvPayloadSource{os.Getenv(configs.PluginConfigPayloadKey)}
}
