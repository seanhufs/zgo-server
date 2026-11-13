package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed web/**
var assets embed.FS

// http://localhost:50001/global_center/user_login?open_id=0&region_code=CN&version=1&time_stamp=1763452352&zone_id=31
// http://localhost:24002/login/getServerList?time_stamp=1453456546&&key=6841f08gfadfd
// http://localhost:50001/global_center/guess_bind?open_id=0&platform_id=0&time_stamp=1234&sign=1234
// 获取公告
func main() {
	fmt.Println("Hello, World!")
	//http.ListenAndServe(":8080", http.FileServer(http.Dir("./test")))
	mux := http.NewServeMux()
	fsroot, _ := fs.Sub(assets, "web")
	mux.Handle("/console/", http.FileServer(http.FS(fsroot)))
	mux.Handle("/js/", http.FileServer(http.FS(fsroot)))
	index := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
			r.URL.Path = "/"
			next.ServeHTTP(wr, r)
		})
	}
	mux.Handle("/", index(http.FileServer(http.FS(fsroot))))
	mux.Handle("/assets/", http.FileServer(http.FS(fsroot)))

	http.ListenAndServe(":8080", mux)
}
