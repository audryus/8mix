package playlist

import (
	"context"
)

type IPlaylistRepo interface {
	Create(context.Context, *Playlist) (*Playlist, error)
	Find(ctx context.Context, playlist *Playlist) (*Playlist, error)
}
type IPlaylistWorkflow interface {
	Start(context.Context, *Playlist) error
}

type PlaylistService struct {
	repo IPlaylistRepo
	work IPlaylistWorkflow
}

func NewPlaylistService(repo IPlaylistRepo, work IPlaylistWorkflow) *PlaylistService {
	return &PlaylistService{
		repo: repo,
		work: work,
	}
}

func (s *PlaylistService) Save(ctx context.Context, pl *Playlist) (*Playlist, error) {
	list, err := s.repo.Find(ctx, pl)

	if err == nil {
		return list, err
	}

	newPl, err := s.repo.Create(ctx, pl)

	if err != nil {
		return nil, err
	}

	if err := s.work.Start(ctx, newPl); err != nil {
		return nil, err
	}

	return newPl, nil
}
