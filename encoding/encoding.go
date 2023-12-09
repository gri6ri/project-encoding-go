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
	// ниже реализуйте метод
	//var dockerCompose models.DockerCompose
	// Прочитать данные из JSON файла
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	// Распаковать JSON данные
	err = json.Unmarshal(jsonFile, &j.DockerCompose)
	if err != nil {
		return err
	}

	// Преобразовать данные в YAML формат
	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return err
	}

	// Записать YAML данные в файл
	err = os.WriteFile(j.FileOutput, yamlData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	//var dockerCompose models.DockerCompose

	// Открыть YAML файл для чтения
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	// Распаковать YAML данные
	err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
	if err != nil {
		return err
	}

	// Преобразовать данные в JSON формат
	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		return err
	}

	// Записать JSON данные в файл
	err = os.WriteFile(y.FileOutput, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
