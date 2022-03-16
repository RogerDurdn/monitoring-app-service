package main

import (
  "github.com/RogerDurdn/MonitoringApp/pkg/api"
  "github.com/RogerDurdn/MonitoringApp/pkg/lib"
)

func main(){
  //router := mux.NewRouter()

  //fs := http.FileServer(http.Dir("./static"))
  //router.HandleFunc("/api/thumbnail", api.ThumbnailHandler)
  //router.HandleFunc("/socket", api.ChatConnection)
  //router.PathPrefix("/").Handler(http.StripPrefix("/", fs))
  //log.Println("server on"); log.Panic(http.ListenAndServe(":9090", router))
  lib.InitCron()
  api.Serve()
}