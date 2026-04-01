package controllers

import (
	"services-sipala/config"
	"services-sipala/utilities"
)

type controller struct {
	cfg *config.Config
	utilities utilities.IUtility
}

func New(
	cfg *config.Config,
	utilities utilities.IUtility,
) IMonitoringController {
	return &controller{
		cfg: cfg,
		utilities: utilities,
	}
}