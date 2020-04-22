package main

import (
	"flag"
	"fmt"
	"kzmap/pkg/kreedz"
	"os"
)

func main()  {
	var downloadMap, downloadDemoList, UnRar, Reformat bool

	flag.BoolVar(&downloadDemoList, "d", false, "是否下载Demo列表文件")
	flag.BoolVar(&downloadMap, "m", false, "是否下载地图包")
	flag.BoolVar(&UnRar, "u", false, "是否加压地图包")
	flag.BoolVar(&Reformat, "f", false, "是否整理地图包文件")

	var target int64
	flag.Int64Var(&target, "t", 0, " 1-XJ 2-CC 3-WS")
	if target < 1 || target > 3 {
		flag.Usage()
		os.Exit(1)
	}

	if !downloadDemoList && !downloadMap && !UnRar && !Reformat {
		flag.Usage()
		os.Exit(1)
	}

	record := &kreedz.WorldRecord{
		Organization: kreedz.Organization(target),
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

	if UnRar {
		err := record.UnRarFiles()
		if err != nil {
			panic(err)
		}
	}

	if Reformat {
		record.Reformat()
	}
	fmt.Println("done")
}
