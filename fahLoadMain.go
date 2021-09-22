package main

import (
	"fahLoadTest/api/fahTest"
	"fahLoadTest/utils/config"
	"fahLoadTest/utils/logger"
	"fahLoadTest/utils/oracle"
	"go.uber.org/zap"
	"log"
)

var Config config.Cfg
var Log *zap.Logger

func init() {
	//Читаем переменные окружения
	err := config.InitConfig(&Config)
	if err != nil {
		log.Fatalf("Не удалось прочитать переменные окружения: %v", err)
	}
	//запускаем логер
	Log, err = logger.LogInit(Config)
	if err != nil {
		log.Fatalf("Не удалось запустить ZAP LOGGER: %v", err)
	}
}

func main() {
	Log.Info("Начали !!!!")
	db, err  := oracle.InitConn(Log, Config.FAH_CONN_STRING)
	if err != nil {
		Log.Fatal("Не удалось подключиться к БД",zap.String("ТНС",Config.FAH_CONN_STRING),zap.Error(err))
	}
	fahTest.Run(db,Log,Config.PARALLEL_CNT,Config.FIBER_CNT)
	//db.Close()
}
