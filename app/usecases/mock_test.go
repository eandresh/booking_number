package usecases

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type mongoServiceMock struct {
	mock.Mock
}

func (mmock *mongoServiceMock) Insert(ctx context.Context, obj interface{}) error {
	args := mmock.Called(ctx, obj)

	return args.Error(0)
}

func (mmock *mongoServiceMock) FindAll(ctx context.Context, filter interface{}) ([]bson.M, error) {
	args := mmock.Called(ctx, filter)

	return args.Get(0).([]bson.M), args.Error(1)
}
