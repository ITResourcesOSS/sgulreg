package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ITResourcesOSS/sgul"
	"github.com/ITResourcesOSS/sgulreg/internal/model"
	"github.com/boltdb/bolt"
)

var logger = sgul.GetLogger().Sugar()

// ServiceRepository defines the interface for the service info repository.
type ServiceRepository interface {
	Save(ctx context.Context, service *model.Service) error
	FindAllByServiceName(ctx context.Context, name string) ([]*model.Service, error)
	FindAll(ctx context.Context) ([]*model.Service, error)
}

type serviceRepository struct {
	db *bolt.DB
}

// NewServiceRepository returns a new instance of the BoltDB based service repository.
func NewServiceRepository(db *bolt.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (sr *serviceRepository) Save(ctx context.Context, service *model.Service) error {
	return sr.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("services"))

		// serviceName := []byte(service.Name)
		// for k, v := c.Seek(serviceName); k != nil && bytes.HasPrefix(k, serviceName); k, v = c.Next() {
		// 	fmt.Printf("key=%s, value=%s\n", k, v)
		// }

		buf, err := json.Marshal(service)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(string(buf))

		err = b.Put([]byte(service.InstanceID), buf)

		return err
	})
}

func (sr *serviceRepository) FindAllByServiceName(ctx context.Context, name string) ([]*model.Service, error) {
	var instances []*model.Service
	err := sr.db.View(func(tx *bolt.Tx) error {
		cursor := tx.Bucket([]byte("services")).Cursor()
		serviceName := []byte(name)
		logger.Debugf("Service Name: %s", serviceName)
		for k, v := cursor.Seek(serviceName); k != nil && bytes.HasPrefix(k, serviceName); k, v = cursor.Next() {
			logger.Debugf("key=%s, value=%s\n", k, v)
			var s *model.Service
			json.Unmarshal(v, &s)
			instances = append(instances, s)
		}
		return nil
	})

	return instances, err
}

func (sr *serviceRepository) FindAll(ctx context.Context) ([]*model.Service, error) {
	var instances []*model.Service
	err := sr.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("services"))
		b.ForEach(func(k, v []byte) error {
			var s *model.Service
			json.Unmarshal(v, &s)
			instances = append(instances, s)
			return nil
		})
		return nil
	})

	return instances, err
}
