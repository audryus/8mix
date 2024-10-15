package track

import (
	"context"
)

type ITrackRepo interface {
	Create(context.Context, *Track) (*Track, error)
	Find(context.Context, *Track) (*Track, error)
}
type TrackService struct {
	repo ITrackRepo
}

func NewTrackService(repo ITrackRepo) *TrackService {
	return &TrackService{
		repo: repo,
	}
}

func (s *TrackService) Save(ctx context.Context, track *Track) (*Track, error) {
	tr, err := s.repo.Find(ctx, track)

	if err != nil {
		return s.repo.Create(ctx, track)
	}
	return tr, nil
}
