package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// Чтение данных из JSON файла
	jsonData, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	// Декодирование JSON данных в структуру DockerCompose
	err = json.Unmarshal(jsonData, &j.DockerCompose)
	if err != nil {
		return err
	}

	// Кодирование структуры DockerCompose в YAML
	yamlData, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return err
	}

	// Запись данных в YAML файл
	err = os.WriteFile(j.FileOutput, yamlData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlData, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	// Декодирование YAML данных в структуру DockerCompose
	err = yaml.Unmarshal(yamlData, &y.DockerCompose)
	if err != nil {
		return err
	}

	// Кодирование структуры DockerCompose в JSON
	jsonData, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return err
	}

	// Запись данных в JSON файл
	err = os.WriteFile(y.FileOutput, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
