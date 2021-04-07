package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"go-api-rest/src/models"

	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
)

type Data struct {
	Success bool                                 `json:"success"`
	Data    []models.Employees_emailnotification `json:"data"`
	Errors  []string                             `json:"errors"`
}

func GetTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var data Data
	var todo models.Employees_emailnotification
	var success bool
	todo, success = models.Get(id)
	if !success {
		data.Success = false
		data.Errors = append(data.Errors, "employees_emailnotification not found")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	data.Success = true
	data.Data = append(data.Data, todo)
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetTodos(w http.ResponseWriter, req *http.Request) {
	var todos []models.Employees_emailnotification = models.GetAll()
	GenerateExcel(todos)

	data, err := ioutil.ReadFile("./test_employees_emailnotification.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+"test_employees_emailnotification.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	http.ServeContent(w, req, "test_employees_emailnotification.xlxs", time.Now(), bytes.NewReader(data))

	// var data = Data{true, todos, make([]string, 0)}
	// json, _ := json.Marshal(data)

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(json)
}

func GenerateExcel(data []models.Employees_emailnotification) {
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

	if err := f.SaveAs("./test_employees_emailnotification.xlsx"); err != nil {
		println(err.Error())
	}

}
