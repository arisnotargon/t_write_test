package service

import (
	"container/list"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/arisnotargon/t_write_test/models/request"
	"github.com/arisnotargon/t_write_test/models/response"
)

type Service struct {
	age         *int32 // initial value of X is 2
	carMutex    *sync.Mutex
	carNumCache int
	SoldNumList *list.List
}

func NewService() *Service {
	var age int32 = 2
	soldNumList := list.New()

	srv := &Service{age: &age, carMutex: &sync.Mutex{}, SoldNumList: soldNumList}
	srv.updateSoldNumList()
	go srv.updateSoldNumListEveryMinute()

	return srv
}

// updateSoldNumList private function run every minute
func (s *Service) updateSoldNumListEveryMinute() {
	ticker := time.NewTicker(time.Minute)
	for {
		<-ticker.C
		s.updateSoldNumList()
	}
}

func (s *Service) updateSoldNumList() {
	defer s.carMutex.Unlock()

	s.carMutex.Lock()

	// only keeps data for last 60 minutes
	if s.SoldNumList.Len() == 60 {
		// delete the last one
		e := s.SoldNumList.Back()
		s.SoldNumList.Remove(e)
	}

	s.SoldNumList.PushFront(s.carNumCache)

	// reset carNumCache
	s.carNumCache = 0
}

func (s *Service) SetAge(p *request.SetAge) (*response.OkResponse, int) {
	atomic.StoreInt32(s.age, int32(p.Age))

	return &response.OkResponse{Ok: true}, 200
}

func (s *Service) GetAge() (*response.AgeResponse, int) {
	return &response.AgeResponse{Age: *s.age}, 200
}

func (s *Service) GetCar() (*response.CarResponse, int) {
	// release lock after return
	defer s.carMutex.Unlock()

	// get lock
	s.carMutex.Lock()
	s.carNumCache++

	nowUnixNano := time.Now().UnixNano()

	return &response.CarResponse{Car: fmt.Sprintf("%d", nowUnixNano)}, 200
}

func (s *Service) getRate() int {
	count := 0
	e := s.SoldNumList.Front()
	for {
		soldNum, ok := e.Value.(int)
		if !ok {
			fmt.Printf("not ok value=>[%+v]", e.Value)
		}

		count += soldNum
		if e.Next() == nil {
			break
		} else {
			e = e.Next()
		}

	}

	return count
}

func (s *Service) GetRate() (*response.RateResponse, int) {
	rate := s.getRate()

	return &response.RateResponse{Rate: rate}, 200
}

func (s *Service) GetBuffer() (*response.BufferResponse, int) {
	rate := s.getRate()

	return &response.BufferResponse{Buffer: rate * int(*s.age)}, 200
}
