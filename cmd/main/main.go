package main

import (
	"TestProject/internal/config"
	"TestProject/internal/modules"
	"TestProject/internal/server"
	"TestProject/internal/shared/bus"
)

func main() {

	srvr := server.NewServer()

	mainConfig, err := config.GetConfig(config.SetConfigPath())
	if err != nil {
		println(err.Error())
		return
	}

	busProvider, err := bus.NewBusProvider(mainConfig.ErrorDatabaseConfig, mainConfig.CommandBusConfig, mainConfig.QueryBusConfig)
	if err != nil {
		println(err.Error())
		return
	}

	initializer, err := modules.NewModuleInitializer(mainConfig.DatabaseConfig, busProvider.CommandBus(), busProvider.QueryBus())
	if err != nil {
		println(err.Error())
		return
	}

	initializer.AddControllers(srvr)
	err = srvr.Run(mainConfig.ServerConfig)

	if err != nil {
		println(err.Error())
		return
	}

}
