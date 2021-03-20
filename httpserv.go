package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Action struct {
	Action  string `json:"action"`
	ObjName string `json:"object"`
}

type Teacher struct {
	ID        string   `json:"id"`
	Salary    float64  `json:"salary"`
	Subject   string   `json:"subject"`
	Classroom string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

type Stuff struct {
	ID        string   `json:"id"`
	Salary    float64  `json:"salary"`
	Classroom string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

type Student struct {
	ID     string  `json:"id"`
	Class  string  `json:"class"`
	Person struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

var obj_sl []GeneralObject

func GetId(o GeneralObject) string {
	switch o.(type) {
	case *Teacher:
		t := o.(Teacher)
		return t.ID
	case *Stuff:
		s := o.(Stuff)
		return s.ID
	case *Student:
		st := o.(Student)
		return st.ID

	}
	return ""
}

/*func Print (w http.ResponseWriter, obj GeneralObject){
	switch obj.(type){
		case "Teacher":
			io.WriteString(w, obj.Data.ID)
			io.WriteString(w, string(obj.Data.Salary))
			io.WriteString(w, obj.Data.Subject)
			io.WriteString(w, obj.Data.Classroom)
			io.WriteString(w, obj.Person.Name)
			io.WriteString(w, obj.Person.Surname)
			io.WriteString(w, obj.Person.PersonalCode)
		case "Stuff":
			io.WriteString(w, obj.Data.ID)
			io.WriteString(w, string(obj.Data.Salary))
			io.WriteString(w, obj.Data.Classroom)
			io.WriteString(w, obj.Person.Name)
			io.WriteString(w, obj.Person.Surname)
			io.WriteString(w, obj.Person.PersonalCode)
		case "Student":
			io.WriteString(w, obj.Data.ID)
			io.WriteString(w, obj.Data.Class)
			io.WriteString(w, obj.Person.Name)
			io.WriteString(w, obj.Person.Surname)
			io.WriteString(w, obj.Person.PersonalCode)

			}
}*/

func (t Teacher) GetCreateAction() DefinedAction {
	return &CreateTeacher{}
}
func (t Teacher) GetUpdateAction() DefinedAction {
	return &UpdateTeacher{}
}
func (t Teacher) GetReadAction() DefinedAction {
	return &ReadTeacher{}
}
func (t Teacher) GetDeleteAction() DefinedAction {
	return &DeleteTeacher{}
}

func (s Stuff) GetCreateAction() DefinedAction {
	return &CreateStuff{}
}
func (s Stuff) GetUpdateAction() DefinedAction {
	return &UpdateStuff{}
}
func (s Stuff) GetReadAction() DefinedAction {
	return &ReadStuff{}
}
func (s Stuff) GetDeleteAction() DefinedAction {
	return &DeleteStuff{}
}

func (st Student) GetCreateAction() DefinedAction {
	return &CreateStudent{}
}
func (st Student) GetUpdateAction() DefinedAction {
	return &UpdateStudent{}
}
func (st Student) GetReadAction() DefinedAction {
	return &ReadStudent{}
}
func (st Student) GetDeleteAction() DefinedAction {
	return &DeleteStudent{}
}

type DefinedAction interface {
	GetFromJSON([]byte)
	Process()
}

type GeneralObject interface {
	GetCreateAction() DefinedAction
	GetUpdateAction() DefinedAction
	GetReadAction() DefinedAction
	GetDeleteAction() DefinedAction
}

type CreateTeacher struct {
	T Teacher `json:"data"`
}

func (action *CreateTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateTeacher) Process() {
	obj_sl = append(obj_sl, action.T)
}

type UpdateTeacher struct {
	T Teacher `json:"data"`
}

func (action *UpdateTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action UpdateTeacher) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.T.ID == GetId(obj_sl[i]) {
			obj_sl[i] = action.T
		}
	}
}

type ReadTeacher struct {
	T Teacher `json:"data"`
}

func (action *ReadTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadTeacher, w http.ResponseWriter) Process() {
	salary := fmt.Sprintf("%f", action.T.Salary)
	for i := 0; i < len(obj_sl); i++ {
		if action.T.ID == GetId(obj_sl[i]) {
			io.WriteString(w, action.T.ID)
			io.WriteString(w, salary)
			io.WriteString(w, action.T.Subject)
			io.WriteString(w, action.T.Classroom)
			io.WriteString(w, action.T.Person.Name)
			io.WriteString(w, action.T.Person.Surname)
			io.WriteString(w, action.T.Person.PersonalCode)
		}
	}
}

type DeleteTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *DeleteTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteTeacher) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if GetId(obj_sl[i]) == action.Data.ID {
			obj_sl[i] = nil
		}
	}
}

type CreateStuff struct {
	S Stuff `json:"data"`
}

func (action *CreateStuff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateStuff) Process() {
	obj_sl = append(obj_sl, action.S)
}

type UpdateStuff struct {
	S Stuff `json:"data"`
}

func (action *UpdateStuff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action UpdateStuff) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.S.ID == GetId(obj_sl[i]) {
			obj_sl[i] = action.S
		}
	}
}

type ReadStuff struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *ReadStuff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadStuff) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.Data.ID == GetId(obj_sl[i]) {
			io.WriteString(w, obj.Data.ID)
			io.WriteString(w, string(obj.Data.Salary))
			io.WriteString(w, obj.Data.Classroom)
			io.WriteString(w, obj.Person.Name)
			io.WriteString(w, obj.Person.Surname)
			io.WriteString(w, obj.Person.PersonalCode)
		}
	}
}

type DeleteStuff struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *DeleteStuff) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteStuff) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if GetId(obj_sl[i]) == action.Data.ID {
			obj_sl[i] = nil
		}
	}
}

type CreateStudent struct {
	St Student `json:"data"`
}

func (action *CreateStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateStudent) Process() {
	obj_sl = append(obj_sl, action.St)
}

type UpdateStudent struct {
	St Student `json:"data"`
}

func (action *UpdateStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action UpdateStudent) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.St.ID == GetId(obj_sl[i]) {
			obj_sl[i] = action.St
		}
	}
}

type ReadStudent struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *ReadStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadStudent) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if action.Data.ID == GetId(obj_sl[i]) {
			io.WriteString(w, obj.Data.ID)
			io.WriteString(w, obj.Data.Class)
			io.WriteString(w, obj.Person.Name)
			io.WriteString(w, obj.Person.Surname)
			io.WriteString(w, obj.Person.PersonalCode)
		}
	}
}

type DeleteStudent struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *DeleteStudent) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteStudent) Process() {
	for i := 0; i < len(obj_sl); i++ {
		if GetId(obj_sl[i]) == action.Data.ID {
			obj_sl[i] = nil
		}
	}
}

func Handler(w http.ResponseWriter, req *http.Request) {
	var act Action
	var obj GeneralObject
	switch act.ObjName {
	case "Teacher":
		obj = &Teacher{}
	case "Stuff":
		obj = &Stuff{}
	case "Student":
		obj = &Student{}

	}
	var toDo DefinedAction
	if req.Method == "GET" {
		for i := 0; i < len(obj_sl); i++{
			io.WriteString(w, GetId(obj_sl[i]))
			io.WriteString(w,"/n")
		}
	} else if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }
		
		fmt.Printf("%s\n", data)
		switch act.Action {
			case "create":
				toDo = obj.GetCreateAction()
				toDo.Process()
			case "read":
				toDo = obj.GetReadAction()
				toDo.Process()
		}
		//io.WriteString(w, "successful post")
	}else if req.Method == "PUT" {
		switch act.Action {
			case "update":
				toDo = obj.GetUpdateAction()
				toDo.Process()
		}
	}else {
		w.WriteHeader(405)
	}
	
}

func main() {
	http.HandleFunc("/", Handler)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
