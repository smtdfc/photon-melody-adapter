# photon-melody-adapter

**Photon Melody Adapter** is a plug-and-play WebSocket adapter for the [Photon](https://github.com/smtdfc/photon) framework, based on the excellent [Melody](https://github.com/olahol/melody) WebSocket library.

This adapter allows you to integrate real-time WebSocket support into your Photon applications using Melody’s event-driven API.

---

## ✨ Features

* 🎧 Built on top of Melody — lightweight, elegant WebSocket handling
* 🔌 Seamlessly integrates with Photon’s adapter system
* 🔁 Simple event-based interface (`On`, `Emit`, etc.)
* 🧹 Fully decoupled and modular

---

## 📦 Installation

```bash
go get github.com/smtdfc/photon-melody-adapter
```

---

## 🚀 Usage Example

### 1. Add to your Photon app

```go
package app

import (
  "github.com/smtdfc/photon"
  "github.com/smtdfc/photon-melody-adapter"
)

func Init() *photon.App {
  socketAdapter,_ := photonMelodyAdapter.Init(),
  
  app := photon.NewApp()
  app.Adapter.UseSocketAdapter(socketAdapter)
  
  // Init modules
  InitModule(app)
  
  return app
}

```

Inside `internal/your-module/routes.go`:
```go

func (m *HelloModule) InitRoute(){
  
  socketController := photon.InitSocketController(
    m.App,
    m.Module,
  )
  
  socketController.On("hello",func(client *photon.SocketSession, msg *photon.SocketEventMessage){
    fmt.Println("Client send")
    fmt.Println(msg)
  })
  
  httpController := photon.InitHttpController(
    m.App,
    m.Module,
  )
  
  httpController.RouteSocket("/socket") // Register route for socket use 
}

```

### 2. Connect via browser

```js
const socket = new WebSocket("ws://localhost:3000/ws");
socket.onmessage = (event) => console.log("Message from server:", event.data);
socket.send("Hello!");
```



## 🤝 Contributing

Pull requests, suggestions, and issues are welcome!
Feel free to open a discussion or PR anytime.

---

## 📜 License

This adapter is licensed under the MIT License.
It integrates the [Melody WebSocket library](https://github.com/olahol/melody),
which is distributed under the BSD-2-Clause license.

Please refer to `THIRD_PARTY_LICENSES/melody-LICENSE` for full license text.

---

## 🙏 Acknowledgments

* [Melody](https://github.com/olahol/melody) — Minimalist WebSocket library for Go
* [Photon](https://github.com/smtdfc/photon) — Modular backend framework for Go
