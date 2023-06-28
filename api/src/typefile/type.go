package typefile

type Employee struct{
    Id   int    "json: id"
    Name string "json: name"
}

type Employees []Employee
