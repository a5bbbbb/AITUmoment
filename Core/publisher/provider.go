package publisher

import (
	"aitu-moment/logger"
	"sync"
)

var (
	providerInstance *Provider
	providerOnce     sync.Once
)

type Provider struct {
	EmailPublisher *EmailPublisher
}

func InitProvider() *Provider {
	providerOnce.Do(func() {
		logger.GetLogger().Info("Publisher: Initializing provider...")
		emailPub, err := NewEmailPublisher()
		if err != nil {
			logger.GetLogger().Error("Publisher: Failed to initialize EmailPublisher")
			panic(err)
		}

		providerInstance = &Provider{
			EmailPublisher: emailPub,
		}
	})

	return providerInstance
}

func GetProvider() *Provider {
	if providerInstance == nil {
		logger.GetLogger().Error("Publisher: Provider accessed before initialization")
		panic("publisher provider not initialized")
	}
	return providerInstance
}
