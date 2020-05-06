package snapshots

import "github.com/cosmos/cosmos-sdk/snapshots/types"

const (
	CurrentFormat = types.CurrentFormat
)

var (
	ErrChunkHashMismatch = types.ErrChunkHashMismatch
	ErrUnknownFormat     = types.ErrUnknownFormat
	SnapshotFromABCI     = types.SnapshotFromABCI
)

type (
	Snapshotter = types.Snapshotter
	Snapshot    = types.Snapshot
	Metadata    = types.Metadata
)
