package repositories

import (
	"testing"
)

func Test_Employee_GetAll (t *testing.T) {
	rs, err := Employee_GetAll()
    if len(rs) != 0 {
        t.Errorf("Employee_GetAll() = %d; want 0", 0)
    }
}
