package manager

import (
	"github.com/ue4-community/ue4cli/logic/manager/internal"
	"runtime"
)

type UnrealManager interface {
	ClearCachedData()
	GetEngineVersion(full bool) string
	GetEngineRoot()string
}

func Create() UnrealManager {
	if  runtime.GOOS == "windows"{
		return new(internal.UnrealManagerWindow)
	}

}

type UnrealManagerBase struct {
}

func (m UnrealManagerBase) GetEngineRoot() string {
	panic("implement me")
}

func (m UnrealManagerBase) GetEngineVersion(full bool) string {

}

func (m* UnrealManagerBase) ClearCachedData()  {

}

