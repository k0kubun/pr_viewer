package controllers

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"pr_viewer/app/models"
)

func InitDB() {
	db, err := sql.Open("sqlite3", "./db/app.db")
	if err != nil {
		panic(err.Error())
	}
	models.DbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	models.DbMap.AddTable(models.User{}).SetKeys(true, "Id")
	models.DbMap.AddTable(models.Repository{}).SetKeys(true, "Id")
	models.DbMap.AddTable(models.PullRequest{}).SetKeys(true, "Id")

	models.DbMap.DropTables()
	models.DbMap.CreateTables()
}

type GorpController struct {
	*revel.Controller
	Transaction *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := models.DbMap.Begin()
	if err != nil {
		panic(err)
	}
	c.Transaction = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Transaction == nil {
		return nil
	}
	err := c.Transaction.Commit()
	if err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Transaction = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Transaction == nil {
		return nil
	}
	err := c.Transaction.Rollback()
	if err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Transaction = nil
	return nil
}
