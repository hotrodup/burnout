package main

import (
  "github.com/gorilla/websocket"
  "net/http"
  "log"
  fp "path/filepath"
  "gopkg.in/fsnotify.v1"
  "fmt"
)

const (
  UPDATE_FILE = ".hotrod-update"
  BASE_DIR = "/app/"
)

var (
  upgrader  = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
  }
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, server running.")
}

func main() {
  http.HandleFunc("/", baseHandler)
  http.HandleFunc("/ws", wsHandler)
  http.ListenAndServe(":8585", nil)
}

func reader(ws *websocket.Conn) {
    for {
        if _, _, err := ws.NextReader(); err != nil {
            ws.Close()
            break
        }
    }
}

func pingOn(filename, message string, ws *websocket.Conn) *fsnotify.Watcher {
  watcher, err := fsnotify.NewWatcher()
  if err != nil {
    log.Fatal(err)
  }
  go func() {
    for {
      select {
        case _ = <-watcher.Events:
          if err := ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
            return
          }
        case err := <-watcher.Errors:
          log.Println("error:", err)
      }
    }
  }()
  watcher.Add(filename)
  return watcher
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
  ws, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    if _, ok := err.(websocket.HandshakeError); !ok {
      log.Println(err)
    }
    return
  }
  watcher := pingOn(fp.Join(BASE_DIR, UPDATE_FILE), "UPDATE", ws)
  defer watcher.Close()
  reader(ws)
}