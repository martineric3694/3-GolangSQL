package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	empId  int
	name   string
	salary int
}

func connectDB() (*sql.DB, error) {
	// template connect DB
	// db, error := sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/golang_sql")

	// jika tanpa password
	db, error := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")

	if error != nil {
		fmt.Println(error)
		return nil, error
	}

	fmt.Println("Connection Success")
	return db, nil
}

func insertEmployee() {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	nama := "Susi"
	salary := 20000

	_, err = db.Exec("INSERT INTO Employee (name, salary) VALUES (?,?)", nama, salary)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("INSERT SUCCESS")
}

func allEmployee() (hasil []Employee) {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM Employee")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	var result []Employee
	for rows.Next() {
		emp := Employee{}
		var err = rows.Scan(&emp.empId, &emp.name, &emp.salary)

		if err != nil {
			fmt.Print(err)
			return
		}

		result = append(result, emp)
	}
	hasil = result
	return hasil
}

func getOneEmp(id int) (hasil Employee) {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	var result Employee
	// Menggunakan QueryRow hanya single use
	err = db.QueryRow("SELECT * FROM Employee WHERE EmployeeID = ?", id).Scan(&result.empId, &result.name, &result.salary)

	// Menggunakan Prepare
	// stmt, err := db.Prepare("SELECT * FROM Employee WHERE EmployeeID = ?")
	// stmt.QueryRow(id).Scan(&result.empId, &result.name, &result.salary)

	if err != nil {
		fmt.Println(err)
		return
	}
	hasil = result
	return hasil
}

func (e Employee) updateEmployee(id int) {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Exec("UPDATE Employee SET name = ?, salary = ? WHERE employeeID=?", e.name, e.salary, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("UPDATE SUCCESS")
}

func deleteEmployee(id int) {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM Employee WHERE EmployeeID = ?", id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("DELETE SUCCESS")
}

func main() {
	// insertEmployee()

	hasil := allEmployee()

	for _, val := range hasil {
		fmt.Println("Employee ID : ", val.empId)
		fmt.Println("Name : ", val.name)
		fmt.Println("Salary : ", val.salary)
		fmt.Println("---------------------------")
	}

	// hasil := getOneEmp(1)
	// fmt.Printf("Nama : %s (%d) with salary %d", hasil.name, hasil.empId, hasil.salary)

	// emp := Employee{
	// 	empId:  0,
	// 	name:   "Adi",
	// 	salary: 15000,
	// }
	// emp.updateEmployee(1)

	// deleteEmployee(1)
}
