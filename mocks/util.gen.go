// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"io/fs"
	"sync"
	"time"
)

// Ensure, that FSMock does implement FS.
// If this is not the case, regenerate this file with moq.
var _ FS = &FSMock{}

// FSMock is a mock implementation of FS.
//
//	func TestSomethingThatUsesFS(t *testing.T) {
//
//		// make and configure a mocked FS
//		mockedFS := &FSMock{
//			GlobFunc: func(pattern string) ([]string, error) {
//				panic("mock out the Glob method")
//			},
//			OpenFunc: func(name string) (fs.File, error) {
//				panic("mock out the Open method")
//			},
//			ReadFileFunc: func(name string) ([]byte, error) {
//				panic("mock out the ReadFile method")
//			},
//			StatFunc: func(name string) (fs.FileInfo, error) {
//				panic("mock out the Stat method")
//			},
//			SubFunc: func(dir string) (fs.FS, error) {
//				panic("mock out the Sub method")
//			},
//		}
//
//		// use mockedFS in code that requires FS
//		// and then make assertions.
//
//	}
type FSMock struct {
	// GlobFunc mocks the Glob method.
	GlobFunc func(pattern string) ([]string, error)

	// OpenFunc mocks the Open method.
	OpenFunc func(name string) (fs.File, error)

	// ReadFileFunc mocks the ReadFile method.
	ReadFileFunc func(name string) ([]byte, error)

	// StatFunc mocks the Stat method.
	StatFunc func(name string) (fs.FileInfo, error)

	// SubFunc mocks the Sub method.
	SubFunc func(dir string) (fs.FS, error)

	// calls tracks calls to the methods.
	calls struct {
		// Glob holds details about calls to the Glob method.
		Glob []struct {
			// Pattern is the pattern argument value.
			Pattern string
		}
		// Open holds details about calls to the Open method.
		Open []struct {
			// Name is the name argument value.
			Name string
		}
		// ReadFile holds details about calls to the ReadFile method.
		ReadFile []struct {
			// Name is the name argument value.
			Name string
		}
		// Stat holds details about calls to the Stat method.
		Stat []struct {
			// Name is the name argument value.
			Name string
		}
		// Sub holds details about calls to the Sub method.
		Sub []struct {
			// Dir is the dir argument value.
			Dir string
		}
	}
	lockGlob     sync.RWMutex
	lockOpen     sync.RWMutex
	lockReadFile sync.RWMutex
	lockStat     sync.RWMutex
	lockSub      sync.RWMutex
}

// Glob calls GlobFunc.
func (mock *FSMock) Glob(pattern string) ([]string, error) {
	if mock.GlobFunc == nil {
		panic("FSMock.GlobFunc: method is nil but FS.Glob was just called")
	}
	callInfo := struct {
		Pattern string
	}{
		Pattern: pattern,
	}
	mock.lockGlob.Lock()
	mock.calls.Glob = append(mock.calls.Glob, callInfo)
	mock.lockGlob.Unlock()
	return mock.GlobFunc(pattern)
}

// GlobCalls gets all the calls that were made to Glob.
// Check the length with:
//
//	len(mockedFS.GlobCalls())
func (mock *FSMock) GlobCalls() []struct {
	Pattern string
} {
	var calls []struct {
		Pattern string
	}
	mock.lockGlob.RLock()
	calls = mock.calls.Glob
	mock.lockGlob.RUnlock()
	return calls
}

// Open calls OpenFunc.
func (mock *FSMock) Open(name string) (fs.File, error) {
	if mock.OpenFunc == nil {
		panic("FSMock.OpenFunc: method is nil but FS.Open was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockOpen.Lock()
	mock.calls.Open = append(mock.calls.Open, callInfo)
	mock.lockOpen.Unlock()
	return mock.OpenFunc(name)
}

// OpenCalls gets all the calls that were made to Open.
// Check the length with:
//
//	len(mockedFS.OpenCalls())
func (mock *FSMock) OpenCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockOpen.RLock()
	calls = mock.calls.Open
	mock.lockOpen.RUnlock()
	return calls
}

// ReadFile calls ReadFileFunc.
func (mock *FSMock) ReadFile(name string) ([]byte, error) {
	if mock.ReadFileFunc == nil {
		panic("FSMock.ReadFileFunc: method is nil but FS.ReadFile was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockReadFile.Lock()
	mock.calls.ReadFile = append(mock.calls.ReadFile, callInfo)
	mock.lockReadFile.Unlock()
	return mock.ReadFileFunc(name)
}

// ReadFileCalls gets all the calls that were made to ReadFile.
// Check the length with:
//
//	len(mockedFS.ReadFileCalls())
func (mock *FSMock) ReadFileCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockReadFile.RLock()
	calls = mock.calls.ReadFile
	mock.lockReadFile.RUnlock()
	return calls
}

// Stat calls StatFunc.
func (mock *FSMock) Stat(name string) (fs.FileInfo, error) {
	if mock.StatFunc == nil {
		panic("FSMock.StatFunc: method is nil but FS.Stat was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockStat.Lock()
	mock.calls.Stat = append(mock.calls.Stat, callInfo)
	mock.lockStat.Unlock()
	return mock.StatFunc(name)
}

// StatCalls gets all the calls that were made to Stat.
// Check the length with:
//
//	len(mockedFS.StatCalls())
func (mock *FSMock) StatCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockStat.RLock()
	calls = mock.calls.Stat
	mock.lockStat.RUnlock()
	return calls
}

// Sub calls SubFunc.
func (mock *FSMock) Sub(dir string) (fs.FS, error) {
	if mock.SubFunc == nil {
		panic("FSMock.SubFunc: method is nil but FS.Sub was just called")
	}
	callInfo := struct {
		Dir string
	}{
		Dir: dir,
	}
	mock.lockSub.Lock()
	mock.calls.Sub = append(mock.calls.Sub, callInfo)
	mock.lockSub.Unlock()
	return mock.SubFunc(dir)
}

// SubCalls gets all the calls that were made to Sub.
// Check the length with:
//
//	len(mockedFS.SubCalls())
func (mock *FSMock) SubCalls() []struct {
	Dir string
} {
	var calls []struct {
		Dir string
	}
	mock.lockSub.RLock()
	calls = mock.calls.Sub
	mock.lockSub.RUnlock()
	return calls
}

// Ensure, that FileMock does implement File.
// If this is not the case, regenerate this file with moq.
var _ File = &FileMock{}

// FileMock is a mock implementation of File.
//
//	func TestSomethingThatUsesFile(t *testing.T) {
//
//		// make and configure a mocked File
//		mockedFile := &FileMock{
//			CloseFunc: func() error {
//				panic("mock out the Close method")
//			},
//			ReadFunc: func(bytes []byte) (int, error) {
//				panic("mock out the Read method")
//			},
//			StatFunc: func() (fs.FileInfo, error) {
//				panic("mock out the Stat method")
//			},
//		}
//
//		// use mockedFile in code that requires File
//		// and then make assertions.
//
//	}
type FileMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// ReadFunc mocks the Read method.
	ReadFunc func(bytes []byte) (int, error)

	// StatFunc mocks the Stat method.
	StatFunc func() (fs.FileInfo, error)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Read holds details about calls to the Read method.
		Read []struct {
			// Bytes is the bytes argument value.
			Bytes []byte
		}
		// Stat holds details about calls to the Stat method.
		Stat []struct {
		}
	}
	lockClose sync.RWMutex
	lockRead  sync.RWMutex
	lockStat  sync.RWMutex
}

// Close calls CloseFunc.
func (mock *FileMock) Close() error {
	if mock.CloseFunc == nil {
		panic("FileMock.CloseFunc: method is nil but File.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//
//	len(mockedFile.CloseCalls())
func (mock *FileMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// Read calls ReadFunc.
func (mock *FileMock) Read(bytes []byte) (int, error) {
	if mock.ReadFunc == nil {
		panic("FileMock.ReadFunc: method is nil but File.Read was just called")
	}
	callInfo := struct {
		Bytes []byte
	}{
		Bytes: bytes,
	}
	mock.lockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	mock.lockRead.Unlock()
	return mock.ReadFunc(bytes)
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//
//	len(mockedFile.ReadCalls())
func (mock *FileMock) ReadCalls() []struct {
	Bytes []byte
} {
	var calls []struct {
		Bytes []byte
	}
	mock.lockRead.RLock()
	calls = mock.calls.Read
	mock.lockRead.RUnlock()
	return calls
}

// Stat calls StatFunc.
func (mock *FileMock) Stat() (fs.FileInfo, error) {
	if mock.StatFunc == nil {
		panic("FileMock.StatFunc: method is nil but File.Stat was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStat.Lock()
	mock.calls.Stat = append(mock.calls.Stat, callInfo)
	mock.lockStat.Unlock()
	return mock.StatFunc()
}

// StatCalls gets all the calls that were made to Stat.
// Check the length with:
//
//	len(mockedFile.StatCalls())
func (mock *FileMock) StatCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStat.RLock()
	calls = mock.calls.Stat
	mock.lockStat.RUnlock()
	return calls
}

// Ensure, that FileInfoMock does implement FileInfo.
// If this is not the case, regenerate this file with moq.
var _ FileInfo = &FileInfoMock{}

// FileInfoMock is a mock implementation of FileInfo.
//
//	func TestSomethingThatUsesFileInfo(t *testing.T) {
//
//		// make and configure a mocked FileInfo
//		mockedFileInfo := &FileInfoMock{
//			IsDirFunc: func() bool {
//				panic("mock out the IsDir method")
//			},
//			ModTimeFunc: func() time.Time {
//				panic("mock out the ModTime method")
//			},
//			ModeFunc: func() fs.FileMode {
//				panic("mock out the Mode method")
//			},
//			NameFunc: func() string {
//				panic("mock out the Name method")
//			},
//			SizeFunc: func() int64 {
//				panic("mock out the Size method")
//			},
//			SysFunc: func() any {
//				panic("mock out the Sys method")
//			},
//		}
//
//		// use mockedFileInfo in code that requires FileInfo
//		// and then make assertions.
//
//	}
type FileInfoMock struct {
	// IsDirFunc mocks the IsDir method.
	IsDirFunc func() bool

	// ModTimeFunc mocks the ModTime method.
	ModTimeFunc func() time.Time

	// ModeFunc mocks the Mode method.
	ModeFunc func() fs.FileMode

	// NameFunc mocks the Name method.
	NameFunc func() string

	// SizeFunc mocks the Size method.
	SizeFunc func() int64

	// SysFunc mocks the Sys method.
	SysFunc func() any

	// calls tracks calls to the methods.
	calls struct {
		// IsDir holds details about calls to the IsDir method.
		IsDir []struct {
		}
		// ModTime holds details about calls to the ModTime method.
		ModTime []struct {
		}
		// Mode holds details about calls to the Mode method.
		Mode []struct {
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
		// Size holds details about calls to the Size method.
		Size []struct {
		}
		// Sys holds details about calls to the Sys method.
		Sys []struct {
		}
	}
	lockIsDir   sync.RWMutex
	lockModTime sync.RWMutex
	lockMode    sync.RWMutex
	lockName    sync.RWMutex
	lockSize    sync.RWMutex
	lockSys     sync.RWMutex
}

// IsDir calls IsDirFunc.
func (mock *FileInfoMock) IsDir() bool {
	if mock.IsDirFunc == nil {
		panic("FileInfoMock.IsDirFunc: method is nil but FileInfo.IsDir was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsDir.Lock()
	mock.calls.IsDir = append(mock.calls.IsDir, callInfo)
	mock.lockIsDir.Unlock()
	return mock.IsDirFunc()
}

// IsDirCalls gets all the calls that were made to IsDir.
// Check the length with:
//
//	len(mockedFileInfo.IsDirCalls())
func (mock *FileInfoMock) IsDirCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsDir.RLock()
	calls = mock.calls.IsDir
	mock.lockIsDir.RUnlock()
	return calls
}

// ModTime calls ModTimeFunc.
func (mock *FileInfoMock) ModTime() time.Time {
	if mock.ModTimeFunc == nil {
		panic("FileInfoMock.ModTimeFunc: method is nil but FileInfo.ModTime was just called")
	}
	callInfo := struct {
	}{}
	mock.lockModTime.Lock()
	mock.calls.ModTime = append(mock.calls.ModTime, callInfo)
	mock.lockModTime.Unlock()
	return mock.ModTimeFunc()
}

// ModTimeCalls gets all the calls that were made to ModTime.
// Check the length with:
//
//	len(mockedFileInfo.ModTimeCalls())
func (mock *FileInfoMock) ModTimeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockModTime.RLock()
	calls = mock.calls.ModTime
	mock.lockModTime.RUnlock()
	return calls
}

// Mode calls ModeFunc.
func (mock *FileInfoMock) Mode() fs.FileMode {
	if mock.ModeFunc == nil {
		panic("FileInfoMock.ModeFunc: method is nil but FileInfo.Mode was just called")
	}
	callInfo := struct {
	}{}
	mock.lockMode.Lock()
	mock.calls.Mode = append(mock.calls.Mode, callInfo)
	mock.lockMode.Unlock()
	return mock.ModeFunc()
}

// ModeCalls gets all the calls that were made to Mode.
// Check the length with:
//
//	len(mockedFileInfo.ModeCalls())
func (mock *FileInfoMock) ModeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockMode.RLock()
	calls = mock.calls.Mode
	mock.lockMode.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *FileInfoMock) Name() string {
	if mock.NameFunc == nil {
		panic("FileInfoMock.NameFunc: method is nil but FileInfo.Name was just called")
	}
	callInfo := struct {
	}{}
	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//
//	len(mockedFileInfo.NameCalls())
func (mock *FileInfoMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}

// Size calls SizeFunc.
func (mock *FileInfoMock) Size() int64 {
	if mock.SizeFunc == nil {
		panic("FileInfoMock.SizeFunc: method is nil but FileInfo.Size was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSize.Lock()
	mock.calls.Size = append(mock.calls.Size, callInfo)
	mock.lockSize.Unlock()
	return mock.SizeFunc()
}

// SizeCalls gets all the calls that were made to Size.
// Check the length with:
//
//	len(mockedFileInfo.SizeCalls())
func (mock *FileInfoMock) SizeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSize.RLock()
	calls = mock.calls.Size
	mock.lockSize.RUnlock()
	return calls
}

// Sys calls SysFunc.
func (mock *FileInfoMock) Sys() any {
	if mock.SysFunc == nil {
		panic("FileInfoMock.SysFunc: method is nil but FileInfo.Sys was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSys.Lock()
	mock.calls.Sys = append(mock.calls.Sys, callInfo)
	mock.lockSys.Unlock()
	return mock.SysFunc()
}

// SysCalls gets all the calls that were made to Sys.
// Check the length with:
//
//	len(mockedFileInfo.SysCalls())
func (mock *FileInfoMock) SysCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSys.RLock()
	calls = mock.calls.Sys
	mock.lockSys.RUnlock()
	return calls
}
