package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
    var target_path string
    if len(os.Args) == 1 {
        target_path = "/"
    } else {
        target_path = os.Args[1]
    }

	err := filepath.Walk(target_path,
		func(path string, info os.FileInfo, err error) error {
			if err == nil {
				if !info.IsDir() {
					fi, err2 := os.Stat(path)
					if err2 != nil {
                        fmt.Println("Error:", err2)
					} else {
						mtime := fi.ModTime()
						stat := fi.Sys().(*syscall.Stat_t)
						atime := time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
						ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
						uid := stat.Uid
						gid := stat.Gid
                        // No easy way to find user/group name on Android
                        fmt.Println(path, "\t", info.Size(), "\t", info.Mode(), "\t", uid, "\t", gid, "\t", atime.Unix(), "\t", mtime.Unix(), "\t", ctime.Unix())
					}
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
