package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"restfull-api-arcticles/helper"
	"strconv"
	"time"
)

func GetConnection(configuration Config) *sql.DB {
	var (
		TORIMASU = "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	)
	PORT, err := strconv.Atoi(configuration.Get("DB_PORT"))
	helper.PanicIfError(err)

	db, err := sql.Open("mysql", fmt.Sprintf(TORIMASU, configuration.Get("DB_USER"), configuration.Get("DB_PASSWORD"), configuration.Get("DB_HOST"), strconv.Itoa(PORT), configuration.Get("DB_NAME")))
	helper.PanicIfError(err)

	sqlOpenMax, err := strconv.Atoi(configuration.Get("SQL_POOL_MAX"))
	helper.PanicIfError(err)

	sqlOpenMin, err := strconv.Atoi(configuration.Get("SQL_POOL_MIN"))
	helper.PanicIfError(err)

	sqlMaxIdle, err := strconv.Atoi(configuration.Get("SQL_POOL_IDLE_TIME"))
	helper.PanicIfError(err)

	sqlMaxLife, err := strconv.Atoi(configuration.Get("SQL_POOL_LIFE_TIME"))
	helper.PanicIfError(err)

	// db pool
	db.SetMaxOpenConns(sqlOpenMax)
	db.SetMaxIdleConns(sqlOpenMin)
	db.SetConnMaxLifetime(time.Duration(sqlMaxLife) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(sqlMaxIdle) * time.Minute)

	log.Println("Successfully connect in port, ", PORT)

	return db
}
