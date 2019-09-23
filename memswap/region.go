package memswap

const (
	READ    uint8 = 1
	WRITE   uint8 = 2
	EXEC    uint8 = 4
	PRIVATE uint8 = 8
)

type Region struct {
	Path     string
	Filename string
	Mode     uint8
	Start    uintptr
	End      uintptr
}

func New(path string, start, end uintptr) *Region {
	return &Region{}
}
