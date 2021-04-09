package storage

import (
	"fmt"
	"github.com/Elias2389/api-rest/model"
)

// Memory .
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory .
func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

// Create .
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

// Update
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	m.Persons[ID] = *person

	return nil
}

// Delete
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	delete(m.Persons, ID)

	return nil
}

// GetByID
func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	return person, nil
}

// GetAll
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}

