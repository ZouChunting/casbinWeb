package utils

import (
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

/**
 *@Author  zct
 *@Date  2019-07-04
 *@Description
 */

var GlobalSessions *session.Manager

func init()  {
	sessionConfig := &session.ManagerConfig{
		CookieName:"go-session-id",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "127.0.0.1:6379,100",
	}
	var err error
	GlobalSessions, err = session.NewManager("redis",sessionConfig)
	if err!=nil{
		panic(err)
	}
	go GlobalSessions.GC()
}