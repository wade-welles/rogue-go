package rogue

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/unix"
)

const (
	None   Perms = 0x0
	Read   Perms = 0x1
	Write  Perms = 0x2
	Exec   Perms = 0x4
	Priv   Perms = 0x8
	Shared Perms = 0x10
)

type Perms uint8

type Map struct {
	Start  uintptr // Beginning memory address.
	End    uintptr // Ending memory address.
	Perms  Perms   // Memory protection bitmask.
	Offset uintptr // Offset where mapping begins.
	Maj    uint64  // Major device number.
	Min    uint64  // Minor device number.
	Inode  uint64  // If mapped from a file, the file's inode.
	Path   string
	Type   Type // Type of the region. Depends on the Path member.
}

type Mapping []Map

func (m Mapping) Len() int           { return len(m) }
func (m Mapping) Less(i, j int) bool { return m[i].Start < m[j].Start }
func (m Mapping) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Map) String() string {
	return fmt.Sprintf("%0.8x-%0.8x %s %0.8x %d:%d %d %s",
		m.Start, m.End, m.Perms, m.Offset,
		m.Maj, m.Min, m.Inode, m.Path,
	)
}

func (m Map) IsPrivate() bool { return m.Perms&Priv != 0 }

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

func FindByPattern(region *Region, pattern string, reload bool) {}
func GetRegion(name string, index int, filter uint8) *Region    { return &Region{} }
func ReadMemory(address uintptr, size int) []byte               { return []byte{} }
func WriteMemory(address uintptr, bytes []byte) bool            { return false }
func Hex2Bin(pattern string, bytes []byte, byteMask uint8) []byte {
	return []byte{}
}

type Region struct {
	Path     string
	Filename string
	Mode     uint8
	Start    uintptr
	End      uintptr
}

func NewMemoryRegion(path string, start, end uintptr) *Region {
	return &Region{}
}

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

type GameProcess struct {
	BinaryName    string
	ProcDirectory string
	Process       *os.Process
	PID           int
	EngineSO      string
	ClientSO      string
	IsRunning     bool
	Maps          []byte // /proc/PID/maps
}

func (self *GameProcess) ProcPrefix(sectionName string) string {
	return fmt.Sprintf("/proc/%v/%s", self.PID, sectionName)
}

func (self *GameProcess) MapsPath() string { return self.ProcPrefix("maps") }
func (self *GameProcess) ExePath() string  { return self.ProcPrefix("exe") }

func FindGameProcess() *GameProcess {
	gameProcess := &GameProcess{
		BinaryName: "csgo_linux64",
		EngineSO:   "engine_client.so",
		ClientSO:   "client_panorama_client.so",
	}

	cmdlineFiles, _ := filepath.Glob("/proc/*/cmdline")
	for _, file := range cmdlineFiles {
		cmdlineData, _ := ioutil.ReadFile(file)
		if len(cmdlineData) != 0 {
			cmdline := string(cmdlineData)
			directory, binaryName := filepath.Split(cmdline)
			if strings.Contains(binaryName, gameProcess.BinaryName) {
				gameProcess.ProcDirectory = directory
				pathSegments := strings.Split(file, "/")
				gameProcess.PID, _ = strconv.Atoi(pathSegments[2])
				Info("Successfully found ", White("Counter-strike:GO"), LightGray(" process with the PID: "), Yellow(strconv.Itoa(gameProcess.PID)))
			}
		}
	}
	if gameProcess.PID != 0 {

		gameProcess.ParseMaps()

		return gameProcess
	} else {
		Warning("Failed to locate a running instance of ", White("Counter-strike:GO"), LightGray(". Please start the game..."))
		Info("Waiting [", White("15 seconds"), LightGray("] and then searching again, use [")+Yellow("CTRL+C")+LightGray("] to cancel."))
		time.Sleep(15 * time.Second)
		FindGameProcess()
		return nil
	}
}

// ParseMaps parses the process' procfs mapping into a useable data structure.
func (self *GameProcess) ParseMaps() (maps Mapping, err error) {
	file, err := os.Open(self.MapsPath())
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

		maj, min := splitOn(parts[3], ':')
		m.Maj = parseUint(maj)
		m.Min = parseUint(min)

		m.Inode = parseUint(parts[4])
		m.Path = string(parts[len(parts)-1])
		m.Type = self.ParseType(m.Path)
		maps = append(maps, m)
	}
	return maps, s.Err()
}

func (self *GameProcess) Find(pc uintptr) (m Map, ok bool) {
	maps, err := self.ParseMaps()
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

func (self *GameProcess) ParseType(s string) Type {
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
		default:
			return Unknown
		}
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
	err := unix.Stat(self.ExePath(), &stat)
	if err != nil {
		return Data
	}

	if stat.Ino == ino {
		return Exe
	}
	return Unknown
}

//Pid_t Process::attach(const char* processName)
//{
//    // ** search processes for one with processName by inspecting each /proc/<uid>/exe path
//
//    // open proc directory
//    processID = 0;
//    DIR* procDirectory = opendir("/proc");
//    if (!procDirectory) {
//        return processID;
//    }
//
//    // parse directory files
//    char* error;
//    struct dirent* entry;
//    while ((entry = readdir(procDirectory))) {
//        // name of file is the pid
//        uint32_t id = strtol(entry->d_name, &error, 10);
//        if (id == 0 || *error != 0) {
//            continue;
//        }
//        char symbolicPath[128];
//        char truePath[FILENAME_MAX];
//
//        // resolve path of exe symlink
//        snprintf(symbolicPath, sizeof(symbolicPath), "/proc/%u/exe", id);
//        size_t pathLength = readlink(symbolicPath, truePath, sizeof(truePath));
//        if (pathLength > 0) {
//            truePath[pathLength] = 0;
//            // check if executable name matches supplied processName
//            const char* szFileName = basename(truePath);
//            if (strcmp(szFileName, processName) == 0) {
//                processID = id;
//                break;
//            }
//        }
//    }
//    closedir(procDirectory);
//    return processID;
//}
//
//
//Region* Process::getRegion(const char* regionName, int index, uint8_t filter)
//{
//    int counter = 0;
//    for (Region& region : regions) {
//        if (((filter & region.getMode()) == filter) && (strcmp(regionName, region.getFileName()) == 0)) {
//            if (counter == index) return &region;
//            else counter++;
//        }
//    }
//    return nullptr;
//}
//
//Size_t Process::convertHex2Bin(const char* pattern, uint8_t* hexBytes, uint8_t* patternMask)
//{
//    // hexBytes and patternMask must be initialized beforehand and must be atleast the size of the amount of hex bytes in the pattern
//    size_t byteCounter = 0;
//    size_t patternLength = strlen(pattern);
//    for (size_t i = 0; i < patternLength; ++i) {
//        if (isxdigit(pattern[i]) && isxdigit(pattern[i + 1])) {
//            sscanf(pattern + i, "%2x", reinterpret_cast<unsigned int*>(&hexBytes[byteCounter]));
//            patternMask[byteCounter] = true;
//            i++;
//        } else if (pattern[i] == '?') {
//            patternMask[byteCounter] = false;
//        } else if (pattern[i] == ' ') {
//            continue;
//        } else {
//            int errorPos = patternLength + 9;
//            fprintf(stderr, "findPattern: Invalid hex detected\nPattern: %s\n",
//                    pattern);
//            fprintf(stderr, "%*s\n", errorPos, "^");
//            return 0;
//        }
//        ++byteCounter;
//    }
//    // return amount of bytes parsed
//    return byteCounter;
//}
//
//Uintptr_t Process::findPattern(Region* region, const char* pattern, bool reload)
//{
//    if (region == nullptr) {
//        return 0;
//    }
//
//    uint8_t hexBytes[64], patternMask[64];
//    size_t  bytesParsed = convertHex2Bin(pattern, hexBytes, patternMask);
//    if (bytesParsed == 0) {
//        return 0;
//    }
//
//    uintptr_t regionSize = region->getSize();
//    if (readBuffer.size() < regionSize) {
//        readBuffer.resize(regionSize);
//    }
//
//    uintptr_t regionStart = region->getRegionStart();
//
//    ssize_t readSize = regionSize;
//
//    // only reload buffer if specified or new region being read
//    if (reload || lastRegionRead != region) {
//        readSize = readMemory(regionStart, readBuffer.data(), regionSize);
//        lastRegionRead = region;
//    }
//
//    // pattern scan
//    size_t patternIndex = 0;
//    for (ssize_t i = 0; i < readSize; ++i) {
//        if (readBuffer[i] == hexBytes[patternIndex] || !patternMask[patternIndex]) patternIndex++;
//        else patternIndex = 0;
//
//        if (patternIndex == bytesParsed) return regionStart + i - patternIndex + 1;
//    }
//
//    return 0;
//}
//
//Ssize_t Process::readMemory(uintptr_t address, void* result, size_t size)
//{
//    struct iovec local = {result, size};
//    struct iovec remote = {reinterpret_cast<void*>(address), size};
//    return process_vm_readv(processID, &local, 1, &remote, 1, 0);
//}
//
//Ssize_t Process::writeMemory(uintptr_t address, void* value, size_t size)
//{
//    struct iovec local = {value, size};
//    struct iovec remote = {reinterpret_cast<void*>(address), size};
//    return process_vm_writev(processID, &local, 1, &remote, 1, 0);
//}
//
//Template<typename T>
//T Process::read(uintptr_t address, size_t size)
//{
//    T result;
//    readMemory(address, &result, size);
//    return result;
//}
//
//Template<typename T>
//Bool Process::write(uintptr_t address, T value, size_t size)
//{
//    return writeMemory(address, &value, size) > 0;
//}
//
//}
// Process represents a running process.
