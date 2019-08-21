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

type Exception interface{}

type TryCatchBlock struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

func (tc TryCatchBlock) Do() {
	if tc.Finally != nil {
		defer tc.Finally()
	}
	if tc.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tc.Catch(r)
			}
		}()
	}
	tc.Try()
}
