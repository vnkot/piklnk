package link

import (
	"errors"
)

type LinkServiceDeps struct {
	LinkRepository *LinkRepository
}

type LinkService struct {
	LinkRepository *LinkRepository
}

func NewLinkService(deps *LinkServiceDeps) *LinkService {
	return &LinkService{
		LinkRepository: deps.LinkRepository,
	}
}

func (service *LinkService) Create(url string, userID *uint) (*Link, error) {
	link := NewLink(url, userID)

	const maxAttempts = 10
	for attempt := 0; attempt < maxAttempts; attempt++ {
		existedLink, err := service.LinkRepository.GetByHash(link.Hash)

		if err != nil || existedLink == nil {
			break
		}

		link.GenerateHash()

		if attempt == maxAttempts-1 {
			return nil, errors.New("cannot generate unique hash")
		}
	}

	return service.LinkRepository.Create(link)

}
