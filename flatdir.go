package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/qwenode/gogo/convert"
	"github.com/qwenode/gogo/file"
)

func main() {
	// catch all panic
	defer func() {
		err := recover()
		if err != nil {
			//debug.PrintStack()
			log.Println("System Error:", err)
		}
	}()
	sourceDir, _ := filepath.Abs("./")
	toDir := filepath.Join(sourceDir, "_flat")
	log.Println(sourceDir, toDir)
	if !file.Exist(sourceDir) || !file.IsDirectory(sourceDir) {
		log.Fatalln("source directory does not exist")
	}
	if !file.Exist(toDir) || !file.IsDirectory(toDir) {
		os.MkdirAll(toDir, os.ModePerm)
	}
	if sourceDir == toDir {
		log.Fatalln("source directory can't same as flat directory")
	}
	moved := 0
	failed := 0
	_ = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(path, toDir) {
			return nil
		}

		if err != nil || info.IsDir() {
			return nil
		}
		log.Println(path, sourceDir, filepath.Base(filepath.Dir(path)), filepath.Base(sourceDir))
		if filepath.Base(filepath.Dir(path)) == filepath.Base(sourceDir) {
			return nil
		}
		//log.Println(path, info.IsDir(), info.Name(), err)

		toFile := toDir + "/" + info.Name()
		if file.Exist(toFile) {
			toFile = fmt.Sprintf("%s/%s_%s", toDir, convert.ToString(time.Now().Unix()), info.Name())
		}
		log.Println("Moving File:", path, " To:", toFile)
		err = os.Rename(path, toFile)
		if err != nil {
			failed++
			log.Println("Move failed.")
			return nil
		}
		moved++
		log.Println("Success Moved.")
		return nil
	})
	log.Println(fmt.Sprintf("Moved:%d,Failed:%d", moved, failed))
}
