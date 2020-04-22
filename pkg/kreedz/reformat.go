package kreedz

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

var defaultDirs = []string {
	"addons",
	"gfx",
	"maps",
	"models",
	"sound",
	"sounds",
	"sprites",
}

var clearSuffix = []string {
	".jpg",
	".jpeg",
	".gif",
	".bmp",
	".log",
	".png",
	".txt",
	".url",
	".wc",
}

func (a *WorldRecord) isDefaultDir(s string) bool {
	for _, dir := range defaultDirs {
		if dir == s {
			return true
		}
	}
	return false
}

func (a *WorldRecord) Reformat()  {
	dir, err := ioutil.ReadDir(a.MapSaveDir)
	if err != nil {
		return
	}

	a.GetRecords()

	group := sync.WaitGroup{}
	for _, fi := range dir {
		if fi.IsDir() {
			//if a.isDefaultDir(fi.Name()) == false {
			if _, ok := a.mapNameCache[fi.Name()]; ok {
				//sd, _ := ioutil.ReadDir(CStrikeDir + fi.Name())
				//for _, fii := range sd {
				//	a.MoveDir(fi.Name() + "/" + fii.Name())
				//}
				err = a.MoveDir(a.MapSaveDir + "/" + fi.Name(), a.MapSaveDir)
				if err != nil {
					logger.Println(err.Error())
				}
			}
		} else if strings.Index(fi.Name(), ".bsp") > 0 {
			a.MoveBSP(fi.Name())
		}
	}
	group.Wait()

	a.ClearDir(a.MapSaveDir)
}

func (a *WorldRecord) MoveDir(src string, to string) error {
	copyDir(src + "/*", to)
	//command := exec.Command("cp", "-r", src + "/*", to)
	//err := command.Wait()
	//if err != nil {
	//	logger.Println(err.Error())
	//}
	//output, _ := command.Output()
	//fmt.Println(output)

	//fmt.Println(command.Stderr)
	//dir, err := ioutil.ReadDir(src)
	//if err != nil {
	//	return err
	//}
	//
	//for _, fi := range dir {
	//	s := src + "/" + fi.Name()
	//	t := to + "/" + fi.Name()
	//
	//	if fi.IsDir() {
	//		//fmt.Printf("move dir[%s] to [%s] \n", s, t)
	//		a.MoveDir(s, t)
	//	} else {
	//		d := path.Dir(t)
	//		_, err := os.Stat(d)
	//		if os.IsNotExist(err) {
	//			fmt.Println("create dir" + d)
	//			err := os.Mkdir(d, os.ModePerm)
	//			if err != nil {
	//				logger.Println(err.Error())
	//			}
	//		}
	//
	//		err = os.Rename(s, t)
	//		if err != nil {
	//			logger.Println(err.Error())
	//		}
	//	}
	//}
	return nil
}

func (a *WorldRecord) MoveBSP(filename string) {
	err := os.Rename(a.MapSaveDir + "/" + filename, a.MapSaveDir + "/maps/" + filename)
	if err != nil {
		logger.Println(err.Error())
	}
}

func (a *WorldRecord) ClearDir(dir string)  {
	d, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, fi := range d {
		path := dir + "/" + fi.Name()
		if fi.IsDir() && a.isDefaultDir(fi.Name()) == true {
			a.ClearDir(path)
		} else if a.isClearFile(fi.Name()) {
			err := os.Remove(path)
			if err != nil {
				logger.Println(err.Error())
			}
		}
	}
}

func (a *WorldRecord) isClearFile(filename string) bool {
	for _, suffix := range clearSuffix {
		if strings.Index(strings.ToLower(filename), suffix) > 0 {
			return true
		}
	}
	return false
}


func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		logger.Println("only support linux,windows,darwin, but os is " + runtime.GOOS)
		return s
	}
}

func copyDir(src string, dest string) {
	src = FormatPath(src)
	dest = FormatPath(dest)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	}
	if cmd == nil {
		return
	}

	err := cmd.Wait()
	if err != nil {
		logger.Println(err.Error())
		return
	}
}