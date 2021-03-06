package noggo

import (
	. "github.com/nogio/noggo/base"
)

var (
	//编号
	//每一个noggo实例应当设置一个唯一Id，好在分布式环境中区分
	Id	string

	//常量模块
	Const *constModule
	//Map模块
	Mapping	*mappingModule


	//全局配置
	Config configConfig

	//日志模块
	Logger *loggerModule

	//路由器模块
	Router *routerModule
	//会话模块
	Session *sessionModule

	//触发器模块
	Trigger *triggerModule

	//计划模块
	Plan *planModule
)


func init() {
	Config = readJsonConfig()

	Const = &constModule{
		mimes: Map{}, regulars: Map{}, states: Map{},
	}

	Mapping = &mappingModule{
		types: map[string]Map{}, cryptos: map[string]Map{},
	}


	//日志模块
	Logger = &loggerModule{
		drivers: map[string]LoggerDriver{},
	}

	//路由器模块
	Router = &routerModule{
		drivers: map[string]RouterDriver{},
	}
	//会话模块
	Session = &sessionModule{
		drivers: map[string]SessionDriver{},
	}





	//触发器模块
	Trigger = &triggerModule{
		contexts: map[string]TriggerContext{},
		routes: map[string]Map{}, routeNames:[]string{}, routeUris: map[string]string{},
	}

	//计划模块
	Plan = &planModule{
		drivers: map[string]PlanDriver{}, contexts: map[string]PlanContext{},
		routes: map[string]Map{}, routeNames:[]string{}, routeUris: map[string]string{}, routeTimes:map[string][]string{},
	}
}