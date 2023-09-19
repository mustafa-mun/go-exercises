package database

import (
	"encoding/binary"
	"fmt"

	"github.com/boltdb/bolt"
)

type Manager struct {
	DB *bolt.DB
}

func(s *Manager) CreateTasksBucket() error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

// CreateTask saves task to the db. The new task ID is set on u once the data is persisted.
func(s *Manager) CreateTask(task string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		// Retrieve the tasks bucket.
		b := tx.Bucket([]byte("Tasks"))

		id, _ := b.NextSequence()

		// Persist bytes to tasks bucket.
		return b.Put(itob(int(id)), []byte(task))
	})
}

// GetTask will print all tasks
func(s *Manager) GetTasks() error {
	return s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
	
		c := b.Cursor()
	
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%v: %s\n", k, v)
		}
	
		return nil
	})
}

func(s *Manager) DeleteTask(id int) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		err := b.Delete(itob(id))
		if err != nil {
			return err
		}
		return nil
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}