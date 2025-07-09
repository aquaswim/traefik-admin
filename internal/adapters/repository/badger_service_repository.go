package repository

import (
	"errors"
	"github.com/dgraph-io/badger/v4"
	"traefik-admin-go/internal/domain"
)

// BadgerServiceRepository is a Badger implementation of the ServiceRepository interface
type BadgerServiceRepository struct {
	db *badger.DB
}

// NewBadgerServiceRepository creates a new domain.ServiceRepository backed by Badger
func NewBadgerServiceRepository(db *badger.DB) domain.ServiceRepository {
	return &BadgerServiceRepository{
		db: db,
	}
}

// GetAll returns all services
func (r *BadgerServiceRepository) GetAll() ([]*domain.ServiceModel, error) {
	services := make([]*domain.ServiceModel, 0, 10)

	err := r.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(ServicePrefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				var service domain.ServiceModel
				if err := deserialize(val, &service); err != nil {
					return err
				}
				services = append(services, &service)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return services, err
}

// GetByID returns a service by ID
func (r *BadgerServiceRepository) GetByID(id string) (*domain.ServiceModel, error) {
	var service domain.ServiceModel

	err := r.db.View(func(txn *badger.Txn) error {
		key := makeKey(ServicePrefix, id)
		item, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return errors.New("service not found")
			}
			return err
		}

		return item.Value(func(val []byte) error {
			return deserialize(val, &service)
		})
	})

	if err != nil {
		return nil, err
	}
	return &service, nil
}

// Create creates a new service
func (r *BadgerServiceRepository) Create(service *domain.ServiceModel) error {
	return r.db.Update(func(txn *badger.Txn) error {
		key := makeKey(ServicePrefix, service.ID)

		// Check if service already exists
		_, err := txn.Get(key)
		if err == nil {
			return errors.New("service already exists")
		} else if err != badger.ErrKeyNotFound {
			return err
		}

		data, err := serialize(service)
		if err != nil {
			return err
		}

		return txn.Set(key, data)
	})
}

// Update updates an existing service
func (r *BadgerServiceRepository) Update(service *domain.ServiceModel) error {
	return r.db.Update(func(txn *badger.Txn) error {
		key := makeKey(ServicePrefix, service.ID)

		// Check if service exists
		_, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return errors.New("service not found")
			}
			return err
		}

		data, err := serialize(service)
		if err != nil {
			return err
		}

		return txn.Set(key, data)
	})
}

// Delete deletes a service by ID
func (r *BadgerServiceRepository) Delete(id string) error {
	return r.db.Update(func(txn *badger.Txn) error {
		key := makeKey(ServicePrefix, id)

		// Check if service exists
		_, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return errors.New("service not found")
			}
			return err
		}

		return txn.Delete(key)
	})
}
