package main

type Dbconf struct {
	Type     string
	Host     string
	Port     uint16
	User     string
	Pass     string
	Database string
}

type Status struct {
	Installed uint8
}
