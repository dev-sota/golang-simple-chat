# golang-simple-chat

Start the server

```
go run *.go
```

Visit http://localhost:8080/room/1

Connecting and communicating from console to WebSocket server

```
wscat -c ws://localhost:8080/ws/1
```
