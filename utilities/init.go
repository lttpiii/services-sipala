package utilities

import "services-sipala/config"

type Utility struct {
	cfg *config.Config
}

func New(cfg *config.Config) IUtility {
	return &Utility{
		cfg: cfg,
	}
}