package proc

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

const proc = "/proc/"

var (
	prefix         = proc + strconv.Itoa(os.Getpid())
	defaultProcess = Process{
		prefix: prefix,
		maps:   prefix + "/maps",
		exe:    prefix + "/exe",
	}
)

// mmaped region protections and flags.
const (
	None   Perms = 0x0
	Read   Perms = 0x1
	Write  Perms = 0x2
	Exec   Perms = 0x4
	Priv   Perms = 0x8
	Shared Perms = 0x10
)

// Map is a mapped memory region, found in /proc/$$/maps
// See: mmap(2)
type Map struct {
	Start  uintptr // Beginning memory address.
	End    uintptr // Ending memory address.
	Perms  Perms   // Memory protection bitmask.
	Offset uintptr // Offset where mapping begins.
	Maj    uint64  // Major device number.
	Min    uint64  // Minor device number.
	Inode  uint64  // If mapped from a file, the file's inode.

	// If mapped from a file, the file's path. Special values
	// include [stack], [heap], and [vsdo]. See related methods.
	Path string

	Type Type // Type of the region. Depends on the Path member.
}

func (m Map) String() string {
	return fmt.Sprintf("%0.8x-%0.8x %s %0.8x %d:%d %d %s",
		m.Start, m.End, m.Perms, m.Offset,
		m.Maj, m.Min, m.Inode, m.Path,
	)
}

func (m Map) IsPrivate() bool { return m.Perms&Priv != 0 }

// ErrVersion indicates the mapping does not have a thread ID.
// (Usually means the linux version is too old.)
var ErrVersion = errors.New("thread ID needs linux >= 3.4")

// ErrNotStack is returned if the mapping is not a stack.
var ErrNotStack = errors.New("mapping is not a stack")

// ThreadID returns the thread (mapping) ID that corresponds
// to the /proc/$$/task/[id] path. It returns an error if the
// mapping is either not a stack or does not have a thread id.
func (m Map) ThreadID() (int, error) {
	if m.Type&Stack == 0 {
		return 0, ErrNotStack
	}
	i := strings.IndexByte(m.Path, ':')
	if i < 0 {
		return 0, ErrVersion
	}
	return strconv.Atoi(m.Path[i+1 : len(m.Path)-1])
}

// ParseMaps parses /proc/$$/maps into a useable data structure.
func ParseMaps() (maps Mapping, err error) {
	return defaultProcess.ParseMaps()
}

// Find searches through /proc/$$/maps to the find the range that holds
// pc. It returns the Map and a boolean indicating whether the Map was found.
func Find(pc uintptr) (m Map, ok bool) {
	return defaultProcess.Find(pc)
}

// Mprotect calls mprotect(2) on the mmapped region.
func (m Map) Mprotect(prot Perms) (err error) {
	_, _, e1 := unix.Syscall(
		unix.SYS_MPROTECT,
		uintptr(m.Start),
		uintptr(m.End-m.Start),
		uintptr(prot),
	)
	if e1 != 0 {
		return e1
	}
	return
}

// Perms are mmap(2)'s memory prot bitmask.
type Perms uint8

func (p Perms) String() string {
	b := [4]byte{'-', '-', '-', 's'}
	if p&None == 0 {
		if p&Read != 0 {
			b[0] = 'r'
		}
		if p&Write != 0 {
			b[1] = 'w'
		}
		if p&Exec != 0 {
			b[2] = 'x'
		}
	}
	if p&Priv != 0 {
		b[3] = 'p'
	}
	return string(b[:])
}

// Type indicates the type of mmaped region.
type Type uint8

const (
	Unknown Type = iota
	Data
	Exe
	Heap
	Lib
	Stack
	VSDO
	VSyscall
	VVar
)

// ParseType parses s into a Type.
func ParseType(s string) Type {
	return defaultProcess.ParseType(s)
}

func (t Type) String() string {
	if int(t) < len(typeStrings) {
		return typeStrings[t]
	}
	return typeStrings[0] // unknown
}

var typeStrings = [...]string{
	Unknown:  "unknown",
	Data:     "data",
	Exe:      "exe",
	Heap:     "[heap]",
	Lib:      "lib",
	Stack:    "[stack]",
	VSDO:     "[vsdo]",
	VSyscall: "[vsyscall]",
	VVar:     "[vvar]",
}
