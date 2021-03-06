package handlers

import (
	"errors"
	"net/http"

	"github.com/sul-dlss-labs/taco/datautils"
	"github.com/sul-dlss-labs/taco/db"
	"github.com/sul-dlss-labs/taco/identifier"
	"github.com/sul-dlss-labs/taco/storage"
)

func handler(database db.Database, storage storage.Storage, identifierService identifier.Service) http.Handler {
	if database == nil {
		database = NewMockDatabase(nil)
	}
	if storage == nil {
		storage = NewMockStorage()
	}

	if identifierService == nil {
		identifierService = identifier.NewUUIDService()
	}

	return BuildAPI(database, storage, identifierService).Serve(nil)
}

type MockDatabase struct {
	record           *datautils.Resource
	CreatedResources []*datautils.Resource
	DeletedResources []string
}

func NewMockDatabase(record *datautils.Resource) db.Database {
	return &MockDatabase{
		CreatedResources: []*datautils.Resource{},
		DeletedResources: []string{},
		record:           record,
	}
}

func (d *MockDatabase) Insert(params *datautils.Resource) error {
	d.CreatedResources = append(d.CreatedResources, params)
	return nil
}

func (d *MockDatabase) RetrieveLatest(externalID string) (*datautils.Resource, error) {
	if d.record != nil {
		record := d.record
		d.record = nil
		return record, nil
	}
	return nil, &db.RecordNotFound{ID: &externalID}
}

func (d *MockDatabase) RetrieveVersion(externalID string, version *string) (*datautils.Resource, error) {
	return nil, errors.New("not implemented")
}

func (d *MockDatabase) Update(params *datautils.Resource) error {
	return nil
}

func (d *MockDatabase) DeleteByID(tacoIdentifier string) error {
	d.DeletedResources = append(d.DeletedResources, tacoIdentifier)
	return nil
}

type MockErrorDatabase struct {
	record *datautils.Resource
}

func NewMockErrorDatabase(record *datautils.Resource) db.Database {
	return &MockErrorDatabase{
		record: record,
	}
}

func (d *MockErrorDatabase) Insert(params *datautils.Resource) error {
	return errors.New("Broken")
}

func (d *MockErrorDatabase) Update(params *datautils.Resource) error {
	return nil
}

func (d *MockErrorDatabase) RetrieveLatest(externalID string) (*datautils.Resource, error) {
	return d.record, nil
}

func (d *MockErrorDatabase) DeleteByID(tacoIdentifier string) error {
	return errors.New("Broken")
}

func (d *MockErrorDatabase) RetrieveVersion(id string, version *string) (*datautils.Resource, error) {
	return nil, errors.New("not implemented")
}
