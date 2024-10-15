package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/audryus/8mix/http/internal/domain/playlist"
	"github.com/audryus/8mix/http/internal/domain/track"
	"github.com/audryus/8mix/http/internal/domain/user"
	"github.com/audryus/8mix/http/internal/usecase"
	"github.com/audryus/8mix/http/pkg/logger"
	"github.com/audryus/8mix/http/pkg/mongo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

var _ playlist.IPlaylistWorkflow = (*MockPlayslistWorkflow)(nil)

type MockPlayslistWorkflow struct {
}

func (w *MockPlayslistWorkflow) Start(ctx context.Context, pl *playlist.Playlist) error {
	return nil
}

func NewMockPlayslistWorkflow() *MockPlayslistWorkflow {
	return &MockPlayslistWorkflow{}
}

func TestCreate(t *testing.T) {

	t.Run("create new playlist with tracks", func(t *testing.T) {
		app := fxtest.New(t, fx.Provide(
			logger.New,
			mongo.New,
			fx.Annotate(track.NewTrackRepo, fx.As(new(track.ITrackRepo))),
			fx.Annotate(track.NewTrackService, fx.As(new(usecase.ITrackService))),
			fx.Annotate(playlist.NewPlaylistRepo, fx.As(new(playlist.IPlaylistRepo))),
			fx.Annotate(NewMockPlayslistWorkflow, fx.As(new(playlist.IPlaylistWorkflow))),
			fx.Annotate(playlist.NewPlaylistService, fx.As(new(usecase.IPlaylistService))),
			usecase.NewPlaylistUC),
			fx.Invoke(func(u *usecase.PlaylistUC, logger *logger.Log) {
				urls := []string{"url_1", "url_2", "url_3", "url_2"}
				playlist, err := u.Create(user.User{
					ID:    "id",
					Email: "a@b.c",
				}, urls, logger)

				fmt.Printf("ID: %s\n\n", playlist.ID)
				assert.Nil(t, err, "error creating")
				assert.NotNil(t, playlist, "playlist nil")

				assert.Equal(t, 3, len(playlist.Tracks))
				assert.NotEqual(t, "", playlist.ID, "id should exists after insert")
			}))

		defer app.RequireStop()
		app.RequireStart()
	})

}
