package main

import (
	// 公共引入

	"github.com/micro/go-micro/util/log"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/gomsa/department-api/hander"
	departmentPB "github.com/gomsa/department-api/proto/department"

	"github.com/gomsa/user/client"
	m "github.com/gomsa/user/middleware"
)

func main() {
	// 设置权限
	h := m.Handler{
		Permissions: Conf.Permissions,
	}
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
		micro.WrapHandler(h.Wrapper), //验证权限
	)
	srv.Init()

	departmentPB.RegisterDepartmentsHandler(srv.Server(), &hander.Department{})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	// 同步权限
	if err := client.SyncPermission(Conf.Permissions); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
