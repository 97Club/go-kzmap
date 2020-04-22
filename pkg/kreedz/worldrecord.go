package kreedz

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	MapRarPath = "./data/kreedz/maprar"
)

var names = []string{
	"ts",
	"xj",
	"cc",
	"ws",
}

var organizations = []string{
	"http://kztop:8080/debug.txt",
	"https://xtreme-jumps.eu/demos.txt",
	"https://cosy-climbing.net/demoz.txt",
	"http://world-surf.com/demos.txt",
}

var localFile = []string{
	"./data/kreedz/debug.txt",
	"./data/kreedz/xj.txt",
	"./data/kreedz/cc.txt",
	"./data/kreedz/ws.txt",
}

var mapSite = []string{
	"http://kztop:8080/maps/%s.rar",
	"http://files.xtreme-jumps.eu/maps/%s.rar",
	"https://cosy-climbing.net/files/maps/%s.rar",
	"http://kztop:8080/maps/%s.rar",
}

var mapRarPaths = []string{
	MapRarPath + "/ts/%s.rar",
	MapRarPath + "/xj/%s.rar",
	MapRarPath + "/cc/%s.rar",
	MapRarPath + "/ws/%s.rar",
}

func init() {
	err := os.MkdirAll("./data/kreedz", 0755)
	err = os.MkdirAll(MapRarPath, 0755)
	err = os.MkdirAll(MapRarPath + "/ts", 0755)
	err = os.MkdirAll(MapRarPath + "/xj", 0755)
	err = os.MkdirAll(MapRarPath + "/cc", 0755)
	err = os.MkdirAll(MapRarPath + "/ws", 0755)
	if err != nil {
		panic(err)
	}
}

func (a *WorldRecord) GetRecords() ([]*RecordInfo, error) {
	newFile, err := os.Open(localFile[a.Organization])
	if err != nil {
		return nil, err
	}
	defer newFile.Close()


	a.mapNameCache = make(map[string]*RecordInfo)
	var records []*RecordInfo
	br := bufio.NewReader(newFile)
	//_, _, _ = br.ReadLine()
	for {
		record, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		recordInfo := a.unserialize(string(record))
		records = append(records, recordInfo)
	}

	return records, nil
}

func (a *WorldRecord) unserialize(record string) *RecordInfo {
	split := strings.Split(record, " ")
	if len(split) < 4 {
		return nil
	}

	time, err := strconv.ParseFloat(split[1], 10)
	if err != nil {
		panic(err)
	}

	mapName := split[0]
	_, route := a.getRoute(mapName)
	recordInfo := &RecordInfo{
		MapName: mapName,
		Holder:  split[2],
		Region:  split[3],
		Time:    time,
		Route:   route,
	}
	a.mapNameCache[mapName] = recordInfo
	return recordInfo
}

func (a *WorldRecord) getRoute(mapName string) (bool, string) {
	r := regexp.MustCompile(`([^\[]+)`)
	//r := regexp.MustCompile(`([^\[]+)\[([^\]]+)\]`)
	subMatch := r.FindAllString(mapName, -1)
	if len(subMatch) > 1 {
		replace := strings.Replace(subMatch[1], "]", "", -1)
		return true, replace
	}
	return false, ""
}
