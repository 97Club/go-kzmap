package kreedz

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (a *WorldRecord) downloadFile(remote string, local string) error {
	resp, err := http.Get(remote)
	if err != nil {
		handleHttpError(err)
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(local)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, resp.Body)
	return nil
}

func (a *WorldRecord) DownloadDemoFile() error {
	return a.downloadFile(organizations[a.Organization], localFile[a.Organization])
}

func (a *WorldRecord) DownloadMapFile() error {
	records, err := a.GetRecords()
	if err != nil {
		return err
	}

	for _, record := range records {
		// 跳过多路径的纪录，不重复下载地图
		if record == nil || record.Route != "" {
			continue
		}

		local := fmt.Sprintf(mapRarPaths[a.Organization], record.MapName)
		stat, err := os.Stat(local)
		if err == nil && stat.Size() > 0 {
			fmt.Println("downloading map " + record.MapName + " rar exist.")
			continue
		}

		// 建议使用单线程，避免被封禁ip
		stat, err = os.Stat(a.MapSaveDir + "/maps/" + record.MapName + ".bsp")
		if err == nil && stat.Size() > 0 {
			fmt.Println("downloading map " + record.MapName + " bsp exist.")
			continue
		}

		if os.IsNotExist(err) {
			local := fmt.Sprintf(mapRarPaths[a.Organization], record.MapName)
			remote := fmt.Sprintf(mapSite[a.Organization], record.MapName)
			fmt.Print("downloading map " + record.MapName)
			err = a.downloadFile(remote, local)
			if err != nil {
				logger.Println(err.Error())
				fmt.Println(" error")
			} else {
				fmt.Println(" done")
			}
		}
	}

	return nil
}
