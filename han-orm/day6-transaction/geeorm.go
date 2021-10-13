package geeorm

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/session"
)

type Engine struct {
	db *sql.DB
	dialect dialect.Dialect
}


func NewEngine(driver, source string) (e *Engine, err error) {
	db,err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}

type TxFunc func(*session.Session) (interface{}, error)

func (engine *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p!= nil {
			// 补货到异常了  这里就回滚事务
			_ = s.Rollback()
		} else if err != nil {
			// 这里回滚  是开启市区失败了 其实这里可以不考虑的
			_ = s.Rollback()
		} else {
			err = s.Commit() // err is nil; if Commit returns error update err
		}
	}()
	return f(s)
}

