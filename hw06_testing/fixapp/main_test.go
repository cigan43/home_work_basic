package fixapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const FilePath = "data.json"

func TestReadJSON(t *testing.T) {

	employee, err := ReadJSON(FilePath)
	assert.NoError(t, err, "нет ошибок в JSON")
	assert.NotNil(t, employee, "Employees не должны быть nil")

	Employee := Employee{UserID: 11, Age: 30, Name: "George", DepartmentID: 2}
	assert.Equal(t, Employee, employee[0])
	//	}
}
