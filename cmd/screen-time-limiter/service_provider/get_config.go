package service_provider

import (
	"os"
	"screen-time-limiter/internal/config"
)

func (p *Provider) MustGetConfig() *config.Config {
	if p.cfg == nil {
		cfg, err := config.Load(os.Getenv("CONFIG_FILE"))
		if err != nil {
			panic(err)
		}

		p.cfg = cfg
	}

	return p.cfg
}
