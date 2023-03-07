package driver_gozero

import (
	_ "github.com/ansoda/dtmdriver-gozero"
	"github.com/dtm-labs/dtmdriver"
)

func init() {
	err := dtmdriver.Use("dtm-driver-gozero")
	if err != nil {
		panic(err)
	}
}
