package dbtools

import (
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"kangqing2008/tools/errors"
)

const(
	USERNAME 	= "root"
	PASSWORD 	= "powerdata"
	DBNAME		= "stocks"
	PROTOCOL	= "tcp"
	HOST		= "localhost"
	PORT		= 3306
	DRIVER		= "mysql"
)

func OpenDatabase()*sql.DB{
	db,err := sql.Open(DRIVER,USERNAME + ":" + PASSWORD + "@" + PROTOCOL + "(" + HOST + ":" + strconv.Itoa(PORT) + ")/" + DBNAME + "?charset=utf8&parseTime=true")
	errors.PanicIfError(err)
	err = db.Ping()
	errors.PanicIfError(err)
	return db
}

//只查询一行数据
func QueryRow(sql,parser func(*sql.Row,interface{},error),args ...interface{})(interface{},error){
	db := OpenDatabase()
	defer db.Close()
	row := db.QueryRow(sql,args)
	item,err := parser(row)
	if err != nil{
		return nil,err
	}
	return item,nil
}

//查询出一个整数值
func QueryInt(sql string,args ...interface{})int64{
	db := OpenDatabase()
	defer db.Close()
	row := db.QueryRow(sql,args)
	var result int64 = -1
	err := row.Scan(&result)
	errors.PanicIfError(err)
	return result
}

//查询当个字符串
func QueryString(sql string,args ...interface{})string{
	db := OpenDatabase()
	defer db.Close()
	row := db.QueryRow(sql,args)
	var result string = ""
	err := row.Scan(&result)
	errors.PanicIfError(err)
	return result
}

//查询单个浮点值
func QueryFloat(sql string,args ...interface{})float64{
	db := OpenDatabase()
	defer db.Close()
	row := db.QueryRow(sql,args)
	var result float64 = 0
	err := row.Scan(&result)
	errors.PanicIfError(err)
	return result
}

