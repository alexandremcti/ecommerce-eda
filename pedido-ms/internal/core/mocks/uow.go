package mocks

import (
	"context"
	"pedido-ms/shared/uow"
)

type UnitOfWorkMock struct {
	repositories map[uow.RepositoryName]uow.Repository
}

func NewUnitOfWorkMock() *UnitOfWorkMock {
	return &UnitOfWorkMock{
		repositories: make(map[uow.RepositoryName]uow.Repository),
	}
}

func (u *UnitOfWorkMock) Register(name uow.RepositoryName, repository uow.Repository) error {
	if _, ok := u.repositories[name]; ok {
		return uow.ErrRepositoryAlreadyRegistered
	}

	u.repositories[name] = repository
	return nil
}

func (u *UnitOfWorkMock) Remove(name uow.RepositoryName) error {
	if _, ok := u.repositories[name]; !ok {
		return uow.ErrRepositoryNotRegistered
	}

	delete(u.repositories, name)
	return nil
}

func (u *UnitOfWorkMock) Has(name uow.RepositoryName) bool {
	_, ok := u.repositories[name]
	return ok
}

func (u *UnitOfWorkMock) Clear() {
	u.repositories = make(map[uow.RepositoryName]uow.Repository)
}

func (u *UnitOfWorkMock) Do(ctx *context.Context, fn func(ctx *context.Context, tx uow.TX) error) error {
	return nil
}
