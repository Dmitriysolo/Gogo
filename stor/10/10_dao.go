package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeDAO struct {
	c *mongo.Collection
}

func NewEmployeeDAO(ctx context.Context, client *mongo.Client) (*EmployeeDAO, error) {
	return &EmployeeDAO{
		c: client.Database("employee").Collection("employee"),
	}, nil
}

func (dao *EmployeeDAO) Insert(ctx context.Context, employee *Employee) error {
	_, err := dao.c.InsertOne(ctx, employee)
	return err
}

func (dao *EmployeeDAO) FindByID(ctx context.Context, id int) (*Employee, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var employee Employee
	err := dao.c.FindOne(ctx, filter).Decode(&employee)
	switch {
	case err == nil:
		return &employee, nil
	case err == mongo.ErrNoDocuments:
		return nil, ErrNotFound
	default:
		return nil, err

	}
}

func (dao *EmployeeDAO) Update(ctx context.Context, e *Employee) error {
	filter := bson.D{{Key: "_id", Value: e.ID}}
	result, err := dao.c.ReplaceOne(ctx, filter, e)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return ErrNotFound
	}

	return nil

}

func (dao *EmployeeDAO) DeleteById(ctx context.Context, id int) error {
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := dao.c.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrNotFound
	}

	return nil

}
