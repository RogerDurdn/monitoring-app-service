package main

import "github.com/RogerDurdn/MonitoringApp/pkg/api"

func main(){
  //router := mux.NewRouter()

  //fs := http.FileServer(http.Dir("./static"))
  //router.HandleFunc("/api/thumbnail", api.ThumbnailHandler)
  //router.HandleFunc("/socket", api.ChatConnection)
  //router.PathPrefix("/").Handler(http.StripPrefix("/", fs))
  //log.Println("server on"); log.Panic(http.ListenAndServe(":9090", router))
  api.Serve()
}