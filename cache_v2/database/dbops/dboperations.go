package dbops

import (
	"fmt"
	"os/exec"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type employee struct {
	EmployeeID   *string
	EmployeeName *string
}

type Dbconnection struct {
	gormDB *gorm.DB
	dbuser string
	passwd string
	dbhost string
	dbport string
	dbname string
}

/*
const (
	dbuser = "root"
	passwd = ""
	dbhost = "127.0.0.1"
	dbport = "3306"
	dbname = "employeerecord"
)

func main() {
	gormDB, err := initDB(dbuser, passwd, dbhost, dbport, dbname)
	if err != nil {
		fmt.Printf("Error in creating connection: %v \n", err)
	}

	createRecord(gormDB, "2", "test2")
	createRecord(gormDB, "3", "test3")
	createRecord(gormDB, "4", "test4")
}
*/

func connectDb(dbUser, passwd, dbHost, dbPort string) (*gorm.DB, error) {

	conf := &mysql.Config{
		Net:                  "tcp",
		Addr:                 dbHost + ":" + dbPort,
		User:                 dbUser,
		Passwd:               passwd,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	return gorm.Open("mysql", conf.FormatDSN())
}

func (db *Dbconnection) InitDB() error {

	cmd := exec.Command("./database/start_db.sh")
	err := cmd.Run()

	if err != nil {
		return err
	}

	db.gormDB, err = connectDb(db.dbuser, db.passwd, db.dbhost, db.dbport)
	if err != nil {
		return err
	}

	db.gormDB = db.gormDB.Exec("CREATE DATABASE " + db.dbname)

	db.gormDB = db.gormDB.Exec("USE " + db.dbname)

	db.gormDB.AutoMigrate(&employee{})

	fmt.Printf("DB Initialize: %v\n", db.gormDB)
	return nil
}

func (db *Dbconnection) CreateRecord(key, value *string) bool {

	results := db.gormDB.Create(&employee{
		EmployeeID:   key,
		EmployeeName: value,
	})

	if results.Error != nil {
		fmt.Printf("Not able to insert data: %v\n", results.Error)
		return false
	}

	return true
}

func (db *Dbconnection) FetchOneRecord(key string) employee {
	var record employee
	db.gormDB.Where("employee_id = ?", key).First(&record)

	return record
}

func (db *Dbconnection) FetchMultipleRecord(key *string) []employee {
	var empName string
	empName = *key
	empName = "%" + empName + "%"

	var results []employee
	db.gormDB.Where("employee_name LIKE ?", empName).Find(&results)

	return results
}
