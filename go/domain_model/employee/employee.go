package employee

import "errors"

type Employee struct {
	EmployeeNumber *EmployeeNumber
	Name           *Name
	MailAddress    *MailAddress
	PhoneNumber    *PhoneNumber
}

// employeeNumberやName、MailAddress、PhoneNumberにてファクトリでの生成を強制することは難しいため、紳士協定とするしかない
// https://developers.cyberagent.co.jp/blog/archives/54440/
func NewEmployee(employeeNumber *EmployeeNumber, name *Name, mailAddress *MailAddress, phoneNumber *PhoneNumber) (*Employee, error) {
	if employeeNumber == nil {
		return nil, errors.New("employee number is nil")
	}
	if name == nil {
		return nil, errors.New("name is nil")
	}
	if mailAddress == nil {
		return nil, errors.New("mail address is nil")
	}
	if phoneNumber == nil {
		return nil, errors.New("phone number is nil")
	}
	return &Employee{
		EmployeeNumber: employeeNumber,
		Name:           name,
		MailAddress:    mailAddress,
		PhoneNumber:    phoneNumber,
	}, nil
}

func (e *Employee) ChangeName(name string) error {
	newName, err := NewName(name)
	if err != nil {
		return err
	}
	e.Name = newName
	return nil
}

func (e *Employee) ChangeMailAddress(mailAddress string) error {
	newMailAddress, err := NewMailAddress(mailAddress)
	if err != nil {
		return err
	}
	e.MailAddress = newMailAddress
	return nil
}

func (e *Employee) ChangePhoneNumber(phoneNumber string) error {
	newPhoneNumber, err := NewPhoneNumber(phoneNumber)
	if err != nil {
		return err
	}
	e.PhoneNumber = newPhoneNumber
	return nil
}

type Employees []*Employee
