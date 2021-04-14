package storage

import (
	"github.com/Elias2389/api-rest/model"
	"testing"
)

func TestCreate(t *testing.T) {
	data := []struct {
		name        string
		person      *model.Person
		wantedError error
	}{
		{"Empty person", nil, model.ErrPersonCanNotBeNil},
		{"AE", &model.Person{Name: "AE"}, nil},
		{"AE", &model.Person{Name: "TEST1"}, nil},
		{"AE", &model.Person{Name: "TEST2"}, nil},
	}

	m := NewMemory()
	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			err := m.Create(v.person)
			if err != v.wantedError {
				t.Errorf("Expected %v found %v", v.wantedError, err)
			}

		})
	}

	wantedCount := len(data) - 1

	if m.currentID != wantedCount {
		t.Errorf("Error, expected %d but foun %d", wantedCount, m.currentID)
	}
}

func TestCreate_empty_person(t *testing.T) {
	memory := NewMemory()
	err := memory.Create(nil)

	if err == nil {
		t.Error("Getting error instead nil")
	}

	if err != model.ErrPersonCanNotBeNil {
		t.Errorf("Expected error: %v, found %v", model.ErrPersonCanNotBeNil, err)
	}
}

func TestCreate_count_elements(t *testing.T) {
	memory := NewMemory()
	person := model.Person{Name: "AE"}
	err := memory.Create(&person)

	if err != nil {
		t.Errorf("Unexpected error but found %v", err)
	}

	if memory.currentID != 1 {
		t.Errorf("Expected 1 element but found %d", memory.currentID)
	}
}
