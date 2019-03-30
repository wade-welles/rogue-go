package proc

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func NewProcess(pid int) Process {
	prefix := proc + strconv.Itoa(pid)
	return Process{
		prefix: prefix,
		maps:   prefix + "/maps",
		exe:    prefix + "/exe",
	}
}

// Process represents a running process.
type Process struct {
	prefix string
	maps   string
	exe    string
}

// ParseMaps parses the process' procfs mapping into a useable data structure.
func (p Process) ParseMaps() (maps Mapping, err error) {
	file, err := os.Open(p.maps)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	var m Map
	for s.Scan() {
		line := s.Bytes()

		parts := bytes.SplitN(line, []byte{' '}, 6)

		// 6 parts minimum, but no max since sometimes
		// there's a big space between inode and path.
		// Prior to 2.0 there was only 5, but I doubt anybody
		// has a kernel from ~2004 that runs Go.
		if len(parts) < 6 {
			return maps, errors.New("proc.ParseMaps: not enough parts (kernel < 2.0?)")
		}

		// Convert the address ranges from hex to uintptr.
		lo, hi := splitOn(parts[0], '-')
		m.Start = hexToUintptr(lo)
		m.End = hexToUintptr(hi)

		// Convert 'rwxp' to permissions bitmask.
		for _, c := range parts[1] {
			switch c {
			case 'r':
				m.Perms |= Read
			case 'w':
				m.Perms |= Write
			case 'x':
				m.Perms |= Exec
			case 'p':
				m.Perms |= Priv
			case 's':
				m.Perms |= Shared
			}
		}

		m.Offset = hexToUintptr(parts[2])

		// Split dev into Major:Minor parts.
		maj, min := splitOn(parts[3], ':')
		m.Maj = parseUint(maj)
		m.Min = parseUint(min)

		m.Inode = parseUint(parts[4])
		m.Path = string(parts[len(parts)-1])
		m.Type = p.ParseType(m.Path)
		maps = append(maps, m)
	}
	return maps, s.Err()
}

// Find searches through the process' mappings to the find the range that holds
// pc. It returns the Map and a boolean indicating whether the Map was found.
func (p Process) Find(pc uintptr) (m Map, ok bool) {
	maps, err := ParseMaps()
	if err != nil {
		return m, false
	}
	for _, m := range maps {
		if pc >= m.Start && pc <= m.End {
			return m, true
		}
	}
	return m, false
}

// ParseType parses s into a Type.
func (p Process) ParseType(s string) Type {
	if s == "" {
		return Unknown
	}

	// See if it's a special value.
	if s[0] == '[' {
		switch s {
		case "[heap]":
			return Heap
		case "[stack]":
			return Stack
		case "[vsdo]":
			return VSDO
		case "[vsyscall]":
			return VSyscall
		case "[vvar]":
			return VVar
		}

		// Fish out stack with thread IDs like [stack:1234]
		if strings.HasPrefix(s, "[stack:") && strings.HasSuffix(s, "]") {
			return Stack
		}
	}

	// Probably is a path.
	// We can't use filepath.Ext here because if, for example, the path is:
	// "/usr/share/lib/libc.so.6" filepath.Ext will return ".6" which is a
	// false negative for a .so file.

	if strings.HasSuffix(s, ".so") || strings.LastIndex(s, ".so.") > 0 {
		return Lib
	}

	var stat unix.Stat_t
	if err := unix.Stat(s, &stat); err != nil {
		return Unknown
	}

	ino := stat.Ino

	err := unix.Stat(p.exe, &stat)
	if err != nil {
		return Data
	}

	if stat.Ino == ino {
		return Exe
	}
	return Unknown
}

// ExePath returns the path of the process (executable).
func (p Process) ExePath() (string, error) {
	return os.Readlink(p.exe)
}
