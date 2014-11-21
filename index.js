var WebSocketServer = require("ws").Server
var http = require("http")
var fs = require("fs")
var path = require("path")
var express = require("express")
var app = express()
app.set("port", process.env.PORT || 5000)
app.set("watchDir", process.env.APP_SRC_DIR || "/app/")

app.use(express.static(__dirname + "/"))

var server = http.createServer(app)
server.listen(app.get("port"))

console.log("http server listening on %d", app.get("port"))

var wss = new WebSocketServer({server: server})
console.log("websocket server created")

wss.on("connection", function(ws) {
  fs.watchFile(path.join(app.get("watchDir"), ".hotrod-update"), function (curr, prev) {
    ws.send(JSON.stringify(new Date()), function() {  })
  });

  console.log("websocket connection open")

  ws.on("close", function() {
    console.log("websocket connection close")
    clearInterval(id)
  })
})
