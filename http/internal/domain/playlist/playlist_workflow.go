package playlist

import (
	"context"

	"go.temporal.io/sdk/client"
)

type PlaylistWorkflow struct {
	cl client.Client
}

func NewPlaylisteWorkflow(cl client.Client) *PlaylistWorkflow {
	return &PlaylistWorkflow{
		cl: cl,
	}
}

func (w *PlaylistWorkflow) Start(ctx context.Context, playlist *Playlist) error {
	options := client.StartWorkflowOptions{
		ID:        playlist.ID.String(),
		TaskQueue: "PLAYLIST_QUEUE",
	}

	_, err := w.cl.ExecuteWorkflow(context.Background(), options, "PlaylistWorkflow", playlist)

	return err
}
