package encoding

import (
	"encoding/json"
	"fmt"
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
	// Чтение файла с JSON
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Println("ошибка при чтении файла:", err)
		return err
	}

	// Десериализация JSON
	err = json.Unmarshal(jsonFile, &j.DockerCompose)
	if err != nil {
		fmt.Println("ошибка при десериализации JSON:", err)
		return err
	}

	// Сереализация в YAML
	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Println("ошибка при сереализации в YAML:", err)
		return err
	}

	// Создание файла yaml
	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Println("ошибка при создании файла YAML:", err)
		return err
	}

	defer yamlFile.Close()

	// Запись данных YAML в файл
	_, err = yamlFile.Write(yamlData)
	if err != nil {
		fmt.Println("ошибка при записи данных в YAML-файл:", err)
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Чтение данных из yaml-файла
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Println("ошибка при чтении данных из YAML-файла:", err)
		return err
	}

	// Десериализация YAML
	err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
	if err != nil {
		fmt.Println("ошибка при десериализзации YAML:", err)
		return err
	}

	// Сереализация в JSON
	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Println("ошибка при сереализации в JSON:", err)
		return err
	}

	// Создание файла JSON
	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Println("ошибка при создании JSON файла:", err)
		return err
	}

	defer jsonFile.Close()

	// Запись данных в файл-json
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("ошибка при записи данных JSON в файл:", err)
		return err
	}

	return nil
}
