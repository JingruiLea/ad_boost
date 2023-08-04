package config

import (
	"github.com/oceanengine/ad_open_sdk_go/config"
)

var configuration *config.Configuration

func init() {
	configuration = config.NewConfiguration()
}
