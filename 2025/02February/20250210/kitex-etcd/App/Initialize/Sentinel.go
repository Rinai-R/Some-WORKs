package Initialize

import (
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Logger"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func InitSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		Logger.Logger.Panic("sentinel error")
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "kitex-etcd",
			Threshold:              50,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})
}
