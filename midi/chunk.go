package midi

//go:generate stringer -type=ChunkType

type Chunk interface {
	Type() ChunkType
}

type ChunkType int

const (
	MTHD ChunkType = iota
	MTRK ChunkType = iota
)
