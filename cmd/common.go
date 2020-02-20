package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func GetEngineRoot()  (string,error){
	override := viper.GetString("rootDirOverride")
	if override != "" {
		fmt.Println("使用用户自定义引擎根路径")
		return override,nil
	}
	return DetectEngineRoot()
}

func DetectEngineRoot() (string ,error){
	baseDir := ""
	prefix := "UE_4."
	switch runtime.GOOS {
	case "windows":
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			if pair[0] == "PROGRAMFILES"{
				baseDir = pair[1] + "\\Epic Games\\"
				break
			}
		}
		versionDirs,err := filepath.Glob(baseDir+prefix+"*")
		if err != nil {
			panic("检测引擎目录出错")
		}

		//取出版本号
		versions := make([]int,0)
		for _,v:= range versionDirs  {
			num,_ := strconv.Atoi(strings.Replace(filepath.Base(v),prefix,"",-1))
			versions = append(versions, num )
		}
		sort.Ints(versions)
		maxVersion:= versions[len(versions)-1]
		return fmt.Sprintf("%s%s%d",baseDir,prefix,maxVersion),nil
	}
	return "",errors.New("无法检测最新安装的虚幻引擎的位置")
}