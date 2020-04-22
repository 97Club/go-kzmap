package kreedz

import (
	"fmt"
	"github.com/gen2brain/go-unarr"
	"io/ioutil"
	"os"
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
			go func(fi os.FileInfo) {
				err := a.unRAR(path + "/" + fi.Name())
				if err != nil {
					logger.Println(fi.Name() + "   " + err.Error())
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

	_, err = r.Extract(CStrikeDir)
	if err != nil {
		return err
	}

	return nil
}