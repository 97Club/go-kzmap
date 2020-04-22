package main

import (
	"flag"
	"fmt"
	"kzmap/pkg/kreedz"
	"os"
)

var (
	mapSaveDir string
	downloadMap, downloadDemoList, unRar, reformat bool
	target int64
)

func init()  {
	flag.StringVar(&mapSaveDir, "p", "", "保存路径")
	flag.Int64Var(&target, "t", 0, " 1-XJ 2-CC 3-WS")
	flag.BoolVar(&downloadDemoList, "d", false, "是否下载Demo列表文件")
	flag.BoolVar(&downloadMap, "m", false, "是否下载地图包")
	flag.BoolVar(&unRar, "u", false, "是否加压地图包")
	flag.BoolVar(&reformat, "f", false, "是否整理地图包文件")
}

func main() {
	flag.Parse()

	if mapSaveDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	if target < 1 || target > 3 {
		flag.Usage()
		os.Exit(1)
	}

	if !downloadDemoList && !downloadMap && !unRar && !reformat {
		flag.Usage()
		os.Exit(1)
	}

	record := &kreedz.WorldRecord{
		Organization: kreedz.Organization(target),
		MapSaveDir:   mapSaveDir,
	}

	if downloadDemoList {
		err := record.DownloadDemoFile()
		if err != nil {
			panic(err)
		}
	}

	if downloadMap {
		record.DownloadMapFile()
	}

	if unRar {
		err := record.UnRarFiles()
		if err != nil {
			panic(err)
		}
	}

	if reformat {
		record.Reformat()
	}
	fmt.Println("done")
}
