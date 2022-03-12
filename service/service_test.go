package service

import (
	"testing"
	"time"

	"github.com/arisnotargon/t_write_test/models/request"
)

func TestNewService(t *testing.T) {
	srv := NewService()

	if srv == nil {
		t.Error("create service failed")
	}
}

func TestUpdateSoldNumList(t *testing.T) {
	srv := NewService()

	if srv == nil {
		t.Error("create service failed in TestUpdateSoldNumList")
		return
	}

	if srv.SoldNumList == nil {
		t.Error("SoldNumList is nil")
		return
	}

	if srv.SoldNumList.Len() != 1 {
		t.Error("SoldNumList len error")
		return
	}
	srv.updateSoldNumList()

	if srv.SoldNumList.Len() != 2 {
		t.Error("SoldNumList len error after updated")
		return
	}

}
func TestSetAge(t *testing.T) {
	srv := NewService()

	if srv == nil {
		t.Error("create service failed in TestUpdateSoldNumList")
		return
	}

	req := &request.SetAge{
		Age: 5,
	}

	res, code := srv.SetAge(req)
	if code != 200 {
		t.Error("SetAge code error")
		return
	}

	if !res.Ok {
		t.Error("SetAge res.Ok error")
		return
	}

	if *(srv.age) != 5 {
		t.Error("srv.age not matched")
		return
	}
}

func TestGetAge(t *testing.T) {
	srv := NewService()

	if srv == nil {
		t.Error("create service failed in TestUpdateSoldNumList")
		return
	}

	res, code := srv.GetAge()
	if code != 200 {
		t.Error("GetAge code error")
		return
	}

	if res.Age != 2 {
		t.Error("GetAge Age init value error")
		return
	}

	req := &request.SetAge{
		Age: 5,
	}

	_, _ = srv.SetAge(req)

	res, code = srv.GetAge()
	if code != 200 {
		t.Error("2nd GetAge code error")
		return
	}
	if res.Age != 5 {
		t.Error("2nd GetAge Age value error")
		return
	}
}

func TestGetCar(t *testing.T) {
	srv := NewService()

	if srv == nil {
		t.Error("create service failed in TestUpdateSoldNumList")
		return
	}

	res, code := srv.GetCar()
	if code != 200 {
		t.Error("GetCar code error")
		return
	}

	if len(res.Car) != 19 {
		t.Error("GetCar res.Car length error")
		return
	}
}

func TestGetRate(t *testing.T) {
	srv := NewService()

	if srv == nil {
		t.Error("create service failed in TestUpdateSoldNumList")
		return
	}

	res, code := srv.GetRate()
	if code != 200 {
		t.Error("GetRate code error")
		return
	}
	if res.Rate != 0 {
		t.Error("GetRate res init val error")
		return
	}

	for i := 0; i < 3; i++ {
		// call GetCar for 3 times
		_, _ = srv.GetCar()
		time.Sleep(time.Second)
	}

	// wait a minute
	time.Sleep(time.Minute)
	res, code = srv.GetRate()
	if code != 200 {
		t.Error("2nd GetRate code error")
		return
	}
	if res.Rate != 3 {
		t.Error("2nd GetRate res init val error")
		return
	}

}

func TestGetBuffer(t *testing.T) {
	srv := NewService()

	if srv == nil {
		t.Error("create service failed in TestUpdateSoldNumList")
		return
	}

	res, code := srv.GetBuffer()

	if code != 200 {
		t.Error("GetBuffer code error")
		return
	}
	if res.Buffer != 0 {
		t.Error("GetBuffer res init val error")
		return
	}

	for i := 0; i < 3; i++ {
		// call GetCar for 3 times
		_, _ = srv.GetCar()
		time.Sleep(time.Second)
	}

	// wait a minute
	time.Sleep(time.Minute)

	// set age  to 3
	req := &request.SetAge{
		Age: 3,
	}
	_, _ = srv.SetAge(req)

	res, code = srv.GetBuffer()

	if code != 200 {
		t.Error("2nd GetBuffer code error")
		return
	}
	if res.Buffer != 3*3 {
		t.Error("2nd GetBuffer res init val error")
		return
	}

	//  set  age to 2
	req = &request.SetAge{
		Age: 2,
	}
	_, _ = srv.SetAge(req)
	res, code = srv.GetBuffer()

	if code != 200 {
		t.Error("3rd GetBuffer code error")
		return
	}
	if res.Buffer != 2*3 {
		t.Error("3nd GetBuffer res init val error")
		return
	}

}
