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

func UpdateCustomer(customer Customer){
	var database *sql.DB
	database = GetConnection()
	var error error
	var update *sql.Stmt
	update , error = database.Prepare("UPDATE CUSTOMER SET CustomerName=?,SSN=? WHERE Customerid=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(customer.CustomerName,customer.SSN,customer.CustomerId)
	defer database.Close()
}
func deleteCustomer(customer Customer){
	var database *sql.DB
	database = GetConnection()
	var error error
	var delete *sql.Stmt
	delete,error = database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if error != nil {
		panic(error.Error())
	}
	delete.Exec(customer.CustomerId)
	defer database.Close()
}

func main() {
	var customers []Customer
	customers = GetCustomer()
	fmt.Println("Before inert",customers)
	var customer Customer
	customer.CustomerName = "Will Smith"
	customer.SSN = "2386343"
	InsertCustomer(customer)
	customers = GetCustomer()
	fmt.Println("After Insert",customers)
	// Update

	fmt.Println("Before Update",customers)
	customer.CustomerName = "MS dhoni"
	customer.SSN = "23233432"
	customer.CustomerId = 3
	UpdateCustomer(customer)
	customers = GetCustomer()
	fmt.Println("After update",customers)

	//DELETE

	fmt.Println("Before Delete",customers)
	customer.CustomerName = "Will Smith"
	customer.SSN = "2386343"
	customer.CustomerId = 4
	deleteCustomer(customer)
	customers = GetCustomer()
	fmt.Println("After Delete",customers)
}