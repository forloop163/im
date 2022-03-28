package im

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Sev struct {
	Ip            string
	Port          uint
	OnlineMap     map[string]*User
	OnlineMapLock sync.Mutex
	MessageChan       chan string
}

func NewSev(ip string, port uint) *Sev {
	server := &Sev{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		MessageChan:   make(chan string),
	}

	return server
}

func (s *Sev) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg
	s.MessageChan <- sendMsg
}

func (s *Sev) ListenMessager() {
	for {
		msg := <- s.MessageChan
		s.OnlineMapLock.Lock()
		for _,cli := range s.OnlineMap {
			cli.ReviceChan <- msg
		}
		s.OnlineMapLock.Unlock()
	}
}

func (s *Sev) Handle(conn net.Conn) {
	user := NewUser(conn, s)
	user.Online()

	isLive := make(chan bool)
	go func() {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if n == 0 {
			user.Offline()
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("Conn Read err:", err)
			return
		}
		msg := string(buf[:n-1])
		user.SendMessage(msg)
		isLive <- true
	}()

	for {
		select {
		case <- isLive:
		case <- time.After(time.Second * 600):
			user.SendMessage("长时间不活跃，停止链接")
			close(user.ReviceChan)
			conn.Close()
			return
		}
	}
}
