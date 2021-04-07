package models

import (
	"go-api-rest/src/db"
	"log"

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
	return todos
}
