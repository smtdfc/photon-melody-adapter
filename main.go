package photonMelodyAdapter

import (
	"fmt"
	"github.com/olahol/melody"
	"github.com/smtdfc/photon/pkg/base"
	"net/http"
)

type MelodyAdapter struct {
	Instance *melody.Melody
	handlers map[string]photon.SocketEventHandler
	name     string
}

func Init() (*MelodyAdapter, *melody.Melody) {
	m := melody.New()
	return &MelodyAdapter{
		Instance: m,
		handlers: make(map[string]photon.SocketEventHandler),
		name:     "MelodyAdapter",
	}, m
}

func (m *MelodyAdapter) GetName() string {
	return m.name
}

func (m *MelodyAdapter) Init() error {
	m.Instance.HandleConnect(func(s *melody.Session) {
		fmt.Println("[MelodyAdapter] Client connected:", s.Request.RemoteAddr)
	})

	m.Instance.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("[MelodyAdapter] Client disconnected:", s.Request.RemoteAddr)
	})

	m.Instance.HandleMessage(func(s *melody.Session, msg []byte) {
		if h, ok := m.handlers["message"]; ok {
			if err := h(msg); err != nil {
				fmt.Println("[MelodyAdapter] Handler error:", err)
			}
		} else {
			s.Write(msg)
		}
	})

	return nil
}

func (m *MelodyAdapter) Start() error {
	return nil
}

func (m *MelodyAdapter) Listen(port string) error {
	return nil
}

func (m *MelodyAdapter) On(event string, handler photon.SocketEventHandler) {
	m.handlers[event] = handler
}

func (m *MelodyAdapter) Emit(event string, data []byte) error {
	if event != "message" {
		return fmt.Errorf("MelodyAdapter chỉ hỗ trợ event 'message'")
	}
	return m.Instance.Broadcast(data)
}

func (m *MelodyAdapter) Stop() error {
	return m.Instance.Close()
}

func (m *MelodyAdapter) HTTPHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Instance.HandleRequest(w, r)
	}
}
