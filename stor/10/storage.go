package main

import (
	"context"
	"sync"
)

type Employee struct {
	ID     int    `bson:"_id"`
	Name   string `bson:"name"`
	Sex    string `bson:"sex"`
	Age    int    `bson:"age"`
	Salary int    `bson:"salary"`
}

type Storage interface {
	Insert(ctx context.Context, e *Employee) error
	Get(ctx context.Context, id int) (*Employee, error)
	Update(ctx context.Context, e *Employee) (*Employee, error)
	Delete(ctx context.Context, id int) error
}

type MemoryStorage struct {
	counter int
	sync.Mutex
	employeeDAO *EmployeeDAO
}

func NewMemoryStorage(employeeDAO *EmployeeDAO) *MemoryStorage {
	return &MemoryStorage{
		employeeDAO: employeeDAO,
		counter:     1,
	}
}

func (s *MemoryStorage) Insert(ctx context.Context, e *Employee) error {

	e.ID = s.counter
	err := s.employeeDAO.Insert(ctx, e)
	if err != nil {
		return err
	}

	s.counter++
	return err

}

func (s *MemoryStorage) Get(ctx context.Context, id int) (*Employee, error) {
	s.Lock()
	defer s.Unlock()

	e, err := s.employeeDAO.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (s *MemoryStorage) Update(ctx context.Context, e *Employee) (*Employee, error) {
	sEMP, err := s.employeeDAO.FindByID(ctx, e.ID)
	if err != nil {
		return nil, err
	}
	sEMP = e

	return sEMP, s.employeeDAO.Update(ctx, sEMP)

}

func (s *MemoryStorage) Delete(ctx context.Context, id int) error {

	return s.employeeDAO.DeleteById(ctx, id)

}
