package kreedz

import (
	"fmt"
	"github.com/gen2brain/go-unarr"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func (a *WorldRecord) UnRarFiles() error {
	name := names[a.Organization]
	path := fmt.Sprintf("%s/%s", MapRarPath, name)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	group := sync.WaitGroup{}
	for _, fi := range dir {
		if !fi.IsDir() {
			group.Add(1)

			mapName := strings.ReplaceAll(fi.Name(), ".rar", "")
			stat, err := os.Stat(a.MapSaveDir + "/maps/" + mapName + ".bsp")
			if err == nil && stat.Size() > 0 {
				fmt.Println("unzip " + mapName + " bsp exist.")
				continue
			}

			go func(fi os.FileInfo) {
				fmt.Println("unzip " + fi.Name())
				err := a.unRAR(path + "/" + fi.Name())
				if err != nil {
					logger.Println(fi.Name() + "   " + err.Error())
					fmt.Println("unzip " + fi.Name() + " error: " + err.Error())
				} else {
					fmt.Println("unzip " + fi.Name() + " done")
				}
				group.Done()
			}(fi)
		}
	}
	group.Wait()

	return nil
}

func (a *WorldRecord) unRAR(file string) error {
	r, err := unarr.NewArchive(file)
	if err != nil {
		return err
	}
	defer r.Close()

	_, err = r.Extract(a.MapSaveDir)
	if err != nil {
		return err
	}

	return nil
}