package main

import (
    "fmt"
    "os"
    "log"
    "path/filepath"
    "syscall"
    "time"
)


func main() {
    err := filepath.Walk("/",
        func(path string, info os.FileInfo, err error) error {
            if err == nil {
                if (!info.IsDir()) {
                    fi, err2 := os.Stat(path)
                    if err2 != nil {
                        fmt.Println("oops")
                    } else {
                        mtime := fi.ModTime()
                        stat := fi.Sys().(*syscall.Stat_t)
                        atime := time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
                        ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))

                        fmt.Println(path, "\t", info.Size(), "\t", info.Mode(), "\t", atime.Unix(), "\t", mtime.Unix(), "\t", ctime.Unix())
                    }
                }
            }
            return nil
    })
    if err != nil {
        log.Println(err)
    }
}
