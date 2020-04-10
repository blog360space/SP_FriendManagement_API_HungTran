package repositories

import "testing"

func Test_Order_GetAll(t *testing.T) {
    rs := Order_GetAll()
    if rs != true {
        t.Errorf("Order_GetAll() = %T; want true", rs)
    }
}
