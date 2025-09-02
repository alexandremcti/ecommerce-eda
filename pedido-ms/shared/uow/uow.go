package uow

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readconcern"
)

var (
	ErrRepositoryNotRegistered     = errors.New("repository not Registered")
	ErrRepositoryAlreadyRegistered = errors.New("repository already registered")
	ErrInvalidRepositoryType       = errors.New("invalid repository type")
)

type RepositoryName string
type Repository any
type RepositoryFactory func(c *mongo.Client) Repository

// Transaction interface
type TX interface {
	Get(name RepositoryName) (Repository, error)
}

// Unit of work interface
type UOW interface {
	Register(name RepositoryName, repository Repository) error
	Remove(name RepositoryName) error
	Has(name RepositoryName) bool
	Clear()
	Do(ctx *context.Context, fn func(ctx *context.Context, tx TX) error) error
}

type Transaction struct {
	tx           *mongo.Session
	repositories map[RepositoryName]Repository
}

func NewTransaction(tx *mongo.Session, repositories map[RepositoryName]Repository) *Transaction {
	return &Transaction{
		tx:           tx,
		repositories: repositories,
	}
}

// Return repository of type T if any found
func GetAs[T any](t TX, name RepositoryName) (T, error) {
	repository, err := t.Get(name)
	var res T
	if err != nil {
		return res, err
	}

	res, ok := repository.(T)
	if !ok {
		log.Println("[Uow GetAS] Tipo de repositório inválido: ", name)
		return res, ErrInvalidRepositoryType
	}

	return res, nil
}

func (t *Transaction) Get(name RepositoryName) (Repository, error) {
	if repository, ok := t.repositories[name]; ok {
		return repository, nil
	}

	return nil, ErrRepositoryNotRegistered
}

type UnitOfWork struct {
	db           *mongo.Client
	repositories map[RepositoryName]Repository
}

func NewUnitOfWork(db *mongo.Client) *UnitOfWork {
	return &UnitOfWork{
		db:           db,
		repositories: make(map[RepositoryName]Repository),
	}
}

func (u *UnitOfWork) Register(name RepositoryName, repository Repository) error {
	if _, ok := u.repositories[name]; ok {
		return ErrRepositoryAlreadyRegistered
	}

	u.repositories[name] = repository
	return nil
}

func (u *UnitOfWork) Remove(name RepositoryName) error {
	if _, ok := u.repositories[name]; !ok {
		return ErrRepositoryNotRegistered
	}

	delete(u.repositories, name)
	return nil
}

func (u *UnitOfWork) Has(name RepositoryName) bool {
	_, ok := u.repositories[name]
	return ok
}

func (u *UnitOfWork) Clear() {
	u.repositories = make(map[RepositoryName]Repository)
}

func (u *UnitOfWork) Do(ctx *context.Context, fn func(ctx *context.Context, tx TX) error) error {
	txnOption := options.Transaction().SetReadConcern(readconcern.Majority())
	sessOpts := options.Session().SetDefaultTransactionOptions(txnOption)
	session, err := u.db.StartSession(sessOpts)
	if err != nil {
		return err
	}
	defer session.EndSession(context.TODO())

	err = fn(ctx, NewTransaction(session, u.repositories))
	if err != nil {
		return err
	}

	return nil
}
