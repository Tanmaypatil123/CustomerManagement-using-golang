package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type Customer struct {
	CustomerId int
	CustomerName string
	SSN string
}

func  GetConnection()(database *sql.DB)  {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := goDotEnvVariable("DATABASE_PASS")
	databaseName := "crm"
	database , error := sql.Open(databaseDriver,databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

func GetCustomer() []Customer{
	var database *sql.DB
	database = GetConnection()
	var error error
	var rows *sql.Rows
	rows , error = database.Query("SELECT  * FROM Customer ORDER BY Customerid DESC ")
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}

	var customers []Customer
	customers =[]Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId,&customerName,&ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers,customer)
	}
	defer database.Close()
	return customers
}

func InsertCustomer(customer Customer){
	var database *sql.DB
	database = GetConnection()
	var error error
	var insert *sql.Stmt
	insert , error = database.Prepare("INSERT INTO Customer (CustomerName , ssn) VALUES (?,?)")
	if error != nil{
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName,customer.SSN)
	defer database.Close()
}

func main() {
	var customers []Customer
	customers = GetCustomer()
	fmt.Println("Customers",customers)
}