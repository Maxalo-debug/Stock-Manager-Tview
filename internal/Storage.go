package internal

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

const (
	filePath = ".file.json"
)

type ApiResponse struct {
	Results []Stock `json:"results"`
}

type Stock struct {
	Name     string    `json:"name"`
	Symbol   string    `json:"symbol"`
	BoughtAt time.Time `json:"boughtAt"`
}
type ApiData struct {
	Results []Stock `json:"results"`
}

type Manager []Stock

func (m *Manager) Add(task string, symbol string) {
	todo := Stock{
		Name:     task,
		Symbol:   symbol,
		BoughtAt: time.Now(),
	}
	*m = append(*m, todo)
}

func (m *Manager) Delete(index int) error {
	if index < 0 || len(*m) <= index {
		return errors.New("Index out of range")
	}

	*m = append((*m)[:index-1], (*m)[index:]...)
	return nil
}

func (m *Manager) Load(filepath string) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(file, m)

	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) Save(filepath string) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
