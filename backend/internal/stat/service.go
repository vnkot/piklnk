package stat

import "github.com/vnkot/piklnk/pkg/event"

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (repo *StatService) AddClick() {
	for e := range repo.EventBus.Subscribe() {
		if e.Type == event.EventLinkClick {
			repo.StatRepository.AddClick(e.Data.(uint))
		}
	}
}
