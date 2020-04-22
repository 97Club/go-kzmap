package kreedz

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
	fmt.Println("start reformat [" + a.MapSaveDir + "]")
	dir, err := ioutil.ReadDir(a.MapSaveDir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	a.GetRecords()

	group := sync.WaitGroup{}
	for _, fi := range dir {
		if fi.IsDir() {
			//if a.isDefaultDir(fi.Name()) == false {
			if _, ok := a.mapNameCache[fi.Name()]; ok {
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
	//getwd, _ := os.Getwd()
	cmd := fmt.Sprintf(`cp -r %s/* %s && rm -rf %s`, src, to, src)
	//cmd := fmt.Sprintf(`cp -r %s/%s/* %s/%s'`, getwd, src, getwd, to)
	fmt.Println(cmd)
	command := exec.Command("/bin/bash", "-c", cmd)
	command.Wait()
	//output, _ := command.Output()
	//fmt.Println(output)
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
		if fi.IsDir() {
			if a.isDefaultDir(fi.Name()) == true && fi.Name() != "data" {
				a.ClearDir(path)
			}
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