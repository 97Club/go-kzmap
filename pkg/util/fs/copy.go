package fs

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/Unknwon/com"
)

const (
	IsDirectory = iota
	IsRegular
	IsSymlink
)

type sysFile struct {
	fType  int
	fName  string
	fLink  string
	fSize  int64
	fMtime time.Time
	fPerm  os.FileMode
}

type F struct {
	files []*sysFile
}

func (self *F) visit(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	var tp int
	if f.IsDir() {
		tp = IsDirectory
	} else if (f.Mode() & os.ModeSymlink) > 0 {
		tp = IsSymlink
	} else {
		tp = IsRegular
	}
	inoFile := &sysFile{
		fName:  path,
		fType:  tp,
		fPerm:  f.Mode(),
		fMtime: f.ModTime(),
		fSize:  f.Size(),
	}
	self.files = append(self.files, inoFile)
	return nil
}

func Copy(sourcedir, decdir string) {
	flag.Parse()

	source := F{
		files: make([]*sysFile, 0),
	}
	err := filepath.Walk(sourcedir, func(path string, f os.FileInfo, err error) error {
		return source.visit(path, f, err)
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	dec := F{
		files: make([]*sysFile, 0),
	}
	err = filepath.Walk(decdir, func(path string, f os.FileInfo, err error) error {
		return dec.visit(path, f, err)
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	for _, v := range source.files {

		if com.IsFile(v.fName) == true {

			tmp1 := strings.Split(v.fName, "\\")
			sourcename := tmp1[len(tmp1)-1]

			for _, r := range dec.files {
				if com.IsFile(r.fName) == true {

					tmp2 := strings.Split(r.fName, "\\")
					decname := tmp2[len(tmp2)-1]

					if sourcename != decname {
						com.Copy(v.fName, r.fName)
					}
				}
			}

		}
	}

}