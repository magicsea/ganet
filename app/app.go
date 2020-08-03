package app

import (
	. "github.com/magicsea/ganet/config"
	"github.com/magicsea/ganet/log"
	"github.com/magicsea/ganet/module"
	"github.com/magicsea/ganet/service"
	"github.com/magicsea/ganet/util"
	"os"
	"os/signal"
)

type MakeServiceFunc func() service.IService

var (
	serviceTypeMap map[string]MakeServiceFunc
	services       []service.IService
	modules        []module.IModule
)

func init() {
	serviceTypeMap = make(map[string]MakeServiceFunc)
}

func RegisterService(serviceType string, f MakeServiceFunc) {
	serviceTypeMap[serviceType] = f
}

func Run(conf *ServiceConfig, ms ...module.IModule) {
	SetGlobleConfig(conf)

	//init log
	if conf.LogConf.LogLevel != "" {
		err := log.NewLogGroup(conf.LogConf.LogLevel, conf.LogConf.LogPath, true, conf.LogConf.LogFlag)
		if err != nil {
			panic(err)
		}
		//log.Export(logger)
		defer log.Close()
	}

	defer util.PrintPanicStack()

	log.Info("log started.")
	modules = ms
	for _, m := range modules {
		if !m.OnInit() {
			log.Fatal("%v module.OnInit fail", m)
		}
	}
	for _, m := range modules {
		m.Run()
	}
	//cluster.InitCluster()
	//生成服务对象
	for _, sc := range conf.Services {
		makefunc := serviceTypeMap[sc.ServiceType]
		if makefunc != nil {
			ser := makefunc()
			log.Info("生成服务:%s", sc.ServiceName)
			ser.Init(sc.RemoteAddr, sc.ServiceName, sc.ServiceType)
			services = append(services, ser)
		} else {
			log.Fatal("未注册的服务类型:%v", sc)
		}
	}

	//init
	for _, ser := range services {
		log.Info("init服务:%s", ser.GetName())
		ser.OnInit()
	}

	//start
	for _, ser := range services {
		log.Info("start服务:%s", ser.GetName())
		service.StartService(ser)
	}

	//service.run
	log.Info("run services")
	//for _, ser := range services {
	//	log.Info("run服务:%s", ser.GetName())
	//	ser.GetPID().Tell(new(service.ServiceRun))
	//}
	//

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Info("closing down (signal: %v)", sig)
	OnDestory()
}

func OnDestory() {
	for _, ser := range services {
		log.Info("destory服务:%s", ser.GetName())
		ser.OnDestory()
	}
	for _, m := range modules {
		m.OnDestroy()
	}
}
