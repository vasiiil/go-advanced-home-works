package stat

import (
	"api-project/pkg/event"
	"log"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (service *StatService) AddClick() {
	for msg := range service.EventBus.Subscribe() {
		if msg.Type != event.EventLinkVisited {
			continue
		}
		linkId, ok := msg.Data.(uint)
		if !ok {
			log.Fatalln("Bad EventLinkVisited Data: ", msg.Data)
			continue
		}
		service.StatRepository.AddClick(linkId)
	}
}
