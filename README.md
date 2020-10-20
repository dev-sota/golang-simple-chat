# golang-simple-chat

Visit http://localhost:8080

Connecting and communicating with a WebSocket server

```
wscat -c ws://localhost:8080/chat
```

POST request

```
curl --location --request POST 'http://localhost:8080/send' \
--form 'userName=yourname' \
--form 'message=hi!'
```
