package staff

var overPaidLimit = 75000
var underPaidLimit = 30000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee
	for _, v := range e.AllStaff {
		if v.Salary > overPaidLimit {
			overpaid = append(overpaid, v)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee
	for _, v := range e.AllStaff {
		if v.Salary < underPaidLimit {
			underpaid = append(underpaid, v)
		}
	}
	return underpaid
}
