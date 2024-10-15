package usecase

import (
	"context"
	"time"

	"github.com/audryus/8mix/http/internal/domain/playlist"
	"github.com/audryus/8mix/http/internal/domain/track"
	"github.com/audryus/8mix/http/internal/domain/user"
	"github.com/audryus/8mix/http/pkg/logger"
)

type ITrackService interface {
	Save(context.Context, *track.Track) (*track.Track, error)
}

type IPlaylistService interface {
	Save(context.Context, *playlist.Playlist) (*playlist.Playlist, error)
}

type PlaylistUC struct {
	trackService    ITrackService
	playlistService IPlaylistService
}

func NewPlaylistUC(trackService ITrackService,
	playlistService IPlaylistService) *PlaylistUC {
	return &PlaylistUC{
		trackService:    trackService,
		playlistService: playlistService,
	}
}

func (u *PlaylistUC) Create(user user.User, urls []string, logger *logger.Log) (*playlist.Playlist, error) {
	ctx, timeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer timeout()

	set := make(map[string]bool)
	tracks := make([]string, 0)
	for _, url := range urls {
		if _, ok := set[url]; ok {
			continue
		}

		_, err := u.trackService.Save(ctx, &track.Track{
			Url:    url,
			Status: "received",
		})

		if err != nil {
			logger.Error("failed to create track %s: %s", url, err.Error())
		}

		tracks = append(tracks, url)
		set[url] = true
	}

	list := &playlist.Playlist{
		Tracks: tracks,
		Status: "pending",
		User:   user.ID,
	}

	playlist, err := u.playlistService.Save(ctx, list)
	if err != nil {
		logger.Error("failed to create playlist %+v: %s", list, err.Error())
		return nil, err
	}

	return playlist, nil
}
