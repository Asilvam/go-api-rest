package models

import (
	"fmt"
	"go-api-rest/src/db"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gopkg.in/guregu/null.v3"
)

type Employees_emailnotification struct {
	Id                int         `json:"id"`
	Task_id           string      `json:"task_id"`
	Status            int         `json:"status"`
	Notification_type int         `json:"notification_type"`
	Moment            null.String `json:"moment"`
	Employees         string      `json:"employees"`
	Created           string      `json:"created"`
	Updated           string      `json:"updated"`
	From_employee_id  null.String `json:"from_employee_id"`
	To_employee_id    int         `json:"to_employee_id"`
}

func Get(ID string) (Employees_emailnotification, bool) {
	db := db.GetConnection()
	row := db.QueryRow("SELECT * FROM employees_emailnotification_2 WHERE id = $1", ID)

	var id int
	var task_id string
	var status int
	var notification_type int
	var moment null.String
	var employees string
	var created string
	var updated string
	var from_employee_id null.String
	var to_employee_id int

	err := row.Scan(
		&id,
		&task_id,
		&status,
		&notification_type,
		&moment,
		&employees,
		&created,
		&updated,
		&from_employee_id,
		&to_employee_id)
	if err != nil {
		return Employees_emailnotification{}, false
	}

	return Employees_emailnotification{
		id,
		task_id,
		status,
		notification_type,
		moment,
		employees,
		created,
		updated,
		from_employee_id,
		to_employee_id}, true
}

func GetAll() []Employees_emailnotification {
	db := db.GetConnection()
	rows, err := db.Query("SELECT * FROM employees_emailnotification_2 ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var todos []Employees_emailnotification
	for rows.Next() {
		t := Employees_emailnotification{}

		var id int
		var task_id string
		var status int
		var notification_type int
		var moment null.String
		var employees string
		var created string
		var updated string
		var from_employee_id null.String
		var to_employee_id int

		err := rows.Scan(
			&id,
			&task_id,
			&status,
			&notification_type,
			&moment,
			&employees,
			&created,
			&updated,
			&from_employee_id,
			&to_employee_id)
		if err != nil {
			log.Fatal(err)
		}

		t.Id = id
		t.Task_id = task_id
		t.Status = status
		t.Notification_type = notification_type
		t.Moment = moment
		t.Employees = employees
		t.Created = created
		t.Updated = updated
		t.From_employee_id = from_employee_id
		t.To_employee_id = to_employee_id

		todos = append(todos, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// Test Excel carga
	Get2Excel(todos)
	return todos
}

func Get2Excel(data []Employees_emailnotification) {
	f := excelize.NewFile()

	sheet1Name := "Hoja 1"
	f.SetSheetName(f.GetSheetName(1), sheet1Name)

	f.SetCellValue(sheet1Name, "A1", "id")
	f.SetCellValue(sheet1Name, "b1", "task_id")
	f.SetCellValue(sheet1Name, "c1", "status")
	f.SetCellValue(sheet1Name, "d1", "notification_type")
	f.SetCellValue(sheet1Name, "e1", "moment")
	f.SetCellValue(sheet1Name, "f1", "employees")
	f.SetCellValue(sheet1Name, "g1", "created")
	f.SetCellValue(sheet1Name, "h1", "updated")
	f.SetCellValue(sheet1Name, "i1", "from_employee_id")
	f.SetCellValue(sheet1Name, "j1", "to_employee_id")

	err := f.AutoFilter(sheet1Name, "A1", "J1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	for i, each := range data {
		f.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each.Id)
		f.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each.Task_id)
		f.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each.Status)
		f.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), each.Notification_type)
		f.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), each.Moment)
		f.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+2), each.Employees)
		f.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+2), each.Created)
		f.SetCellValue(sheet1Name, fmt.Sprintf("H%d", i+2), each.Updated)
		f.SetCellValue(sheet1Name, fmt.Sprintf("I%d", i+2), each.From_employee_id)
		f.SetCellValue(sheet1Name, fmt.Sprintf("J%d", i+2), each.To_employee_id)
	}

	if err := f.SaveAs("test_employees_emailnotification.xlsx"); err != nil {
		println(err.Error())
	}
}
