package oracle

import (
	"context"
	"database/sql"
	_ "github.com/godror/godror"
	"go.uber.org/zap"
	"time"
)

//https://blogs.oracle.com/developers/how-to-connect-a-go-program-to-oracle-database-using-goracle
//https://oracle.github.io/odpi/doc/installation.html#id3
//ln -s /opt/oracle/instantclient_19_8/libclntsh.dylib /usr/local/lib/

func InitConn(Log *zap.Logger, dsn string) (connect *sql.DB, err error) {
	db, err := sql.Open("godror", dsn)
	if err != nil {
		Log.Error("Не удалось создать подключение к БД", zap.String("dsn:", dsn), zap.Error(err))
		return nil, err
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(5 * time.Minute)

	var ctx context.Context
	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		Log.Error("Не удалось проверить подключение к БД", zap.String("dsn:", dsn), zap.Error(err))
		return nil, err
	} else {
		Log.Info("Подключение к БД - OK", zap.String("dsn:", dsn))
	}

	return db, nil
}
