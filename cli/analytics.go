package cli

import (
	"fmt"

	"github.com/bitrise-io/bitrise-plugins-analytics/analytics"
	"github.com/bitrise-io/bitrise-plugins-analytics/payload"
	log "github.com/bitrise-io/go-utils/log"
)

func sendAnalytics(source payload.Source) error {
	payload, err := source.Payload()
	if err != nil {
		return fmt.Errorf("failed to read payload: %s", err)
	}

	log.Infof("")
	log.Infof("Submitting anonymized usage information...")
	log.Infof("For more information visit:")
	log.Infof("https://github.com/bitrise-io/bitrise-plugins-analytics/blob/master/README.md")

	return analytics.SendAnonymizedAnalytics(payload)
}
