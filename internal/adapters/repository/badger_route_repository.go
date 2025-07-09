package repository

import (
	"errors"
	"github.com/dgraph-io/badger/v4"
	"traefik-admin-go/internal/domain"
)

// BadgerRouteRepository is a Badger implementation of the RouteRepository interface
type BadgerRouteRepository struct {
	db *badger.DB
}

// NewBadgerRouteRepository creates a new domain.RouteRepository backed by Badger
func NewBadgerRouteRepository(db *badger.DB) domain.RouteRepository {
	return &BadgerRouteRepository{
		db: db,
	}
}

// GetAll returns all routes
func (r *BadgerRouteRepository) GetAll() ([]*domain.RouteModel, error) {
	routes := make([]*domain.RouteModel, 0, 10)

	err := r.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(RoutePrefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				var route domain.RouteModel
				if err := deserialize(val, &route); err != nil {
					return err
				}
				routes = append(routes, &route)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return routes, err
}

// GetByID returns a route by ID
func (r *BadgerRouteRepository) GetByID(id string) (*domain.RouteModel, error) {
	var route domain.RouteModel

	err := r.db.View(func(txn *badger.Txn) error {
		key := makeKey(RoutePrefix, id)
		item, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return errors.New("route not found")
			}
			return err
		}

		return item.Value(func(val []byte) error {
			return deserialize(val, &route)
		})
	})

	if err != nil {
		return nil, err
	}
	return &route, nil
}

// Create creates a new route
func (r *BadgerRouteRepository) Create(route *domain.RouteModel) error {
	return r.db.Update(func(txn *badger.Txn) error {
		key := makeKey(RoutePrefix, route.ID)

		// Check if route already exists
		_, err := txn.Get(key)
		if err == nil {
			return errors.New("route already exists")
		} else if err != badger.ErrKeyNotFound {
			return err
		}

		data, err := serialize(route)
		if err != nil {
			return err
		}

		return txn.Set(key, data)
	})
}

// Update updates an existing route
func (r *BadgerRouteRepository) Update(route *domain.RouteModel) error {
	return r.db.Update(func(txn *badger.Txn) error {
		key := makeKey(RoutePrefix, route.ID)

		// Check if route exists
		_, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return errors.New("route not found")
			}
			return err
		}

		data, err := serialize(route)
		if err != nil {
			return err
		}

		return txn.Set(key, data)
	})
}

// Delete deletes a route by ID
func (r *BadgerRouteRepository) Delete(id string) error {
	return r.db.Update(func(txn *badger.Txn) error {
		key := makeKey(RoutePrefix, id)

		// Check if route exists
		_, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return errors.New("route not found")
			}
			return err
		}

		return txn.Delete(key)
	})
}
