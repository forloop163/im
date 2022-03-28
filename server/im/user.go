package im

import (
	"net"
	"strings"
)

type User struct {
	Name       string
	Addr       string
	ReviceChan chan string
	Conn       net.Conn
	Server     *Sev
}

func NewUser(conn net.Conn, server *Sev) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:       userAddr,
		Addr:       userAddr,
		ReviceChan: make(chan string),
		Conn:       conn,
		Server:     server,
	}

	go user.ListenMessage()
	return user
}

func (u *User) Online() {
	u.Server.OnlineMapLock.Lock()
	u.Server.OnlineMap[u.Name] = u
	u.Server.OnlineMapLock.Unlock()

	u.Server.BroadCast(u, "已上线")
}

func (u *User) Offline() {
	u.Server.OnlineMapLock.Lock()
	delete(u.Server.OnlineMap, u.Name)
	u.Server.OnlineMapLock.Unlock()
	u.Server.BroadCast(u, "已下线")
}

func (u *User) SendMessage(msg string) {
	u.Conn.Write([]byte(msg))
}

func (u *User) DoMessage(msg string) {
	if msg == "who" {
		u.Server.OnlineMapLock.Lock()
		for _, user := range u.Server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线... \n"
			u.SendMessage(onlineMsg)
		}
		u.Server.OnlineMapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|"  {
		newName := strings.Split(msg, "|")[1]

		_, ok := u.Server.OnlineMap[newName]
		if ok {
			u.SendMessage("当前用户名被使用")
		} else {
			u.Server.OnlineMapLock.Lock()
			delete(u.Server.OnlineMap, u.Name)
			u.Server.OnlineMap[newName] = u
			u.Server.OnlineMapLock.Unlock()
			u.Name = newName
			u.SendMessage("用户名已修改成功")
		}
	} else if len(msg) > 3 && msg[:3] == "to|" {
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			u.SendMessage("消息格式不正确")
			return
		}

		remoteUser, ok := u.Server.OnlineMap[remoteName]

		if !ok {
			u.SendMessage("当前用户名不存在")
			return
		}
		content := strings.Split(msg, "|")[2]
		if content == "" {
			u.SendMessage("无消息内容,请重发\n")
			return
		}
		remoteUser.SendMessage(u.Name + "对您说： " + content)
	} else {
		u.Server.BroadCast(u, msg)
	}
}

func (u *User) ListenMessage() {
	for {
		msg := <-u.ReviceChan
		u.Conn.Write([]byte(msg + "\n"))
	}
}
