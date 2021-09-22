package fahTest

import (
	"database/sql"
	"go.uber.org/zap"
	"sync"
)


func Run(db *sql.DB,Log *zap.Logger,cnt int,fiberCnt int) {

	Log.Info("Кол-во потоков",zap.Int("PARALLEL_CNT",cnt))
	Log.Info("Кол-во документов в потоке",zap.Int("FIBER_CNT",fiberCnt))
	wg := new(sync.WaitGroup)
	for i := 1; i <= cnt; i++ {
		wg.Add(1)
		go executeProc(db,Log,i,wg,fiberCnt)
		Log.Info("Запустили", zap.Int("поток", i))
	}
	wg.Wait()
}

func executeProc(db *sql.DB,Log *zap.Logger,n int,wg *sync.WaitGroup,fiberCnt int) {
	defer wg.Done()

	if _,err := db.Exec("begin  XXT_EB_LOAD_TEST_PKG.exec_100(:v,:p); end;",n,fiberCnt); err != nil {
		Log.Error("Откатили", zap.Int("№:", n), zap.Error(err))
	} else {
		Log.Info("Завершили", zap.Int("поток", n))
	}
}