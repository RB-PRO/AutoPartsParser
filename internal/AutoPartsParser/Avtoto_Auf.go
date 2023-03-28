package AutoPartsParser

import (
	"encoding/json"
	"os"
)

// Структура авторизации в серсисе Avtoto
type AvtotoAuf struct {
	UserId       int    `json:"user_id"`        // Уникальный идентификатор пользователя (номер клиента) (тип: целое)
	UserLogin    string `json:"user_login"`     // Логин пользователя (тип: строка)
	UserPassword string `json:"user_password "` // Пароль пользователя (тип: строка)
}

// Получение значение из файла
func DataFile(filename string) (AvtotoAuf, error) {
	// Прочитать файл
	fileBytes, ErrorReadFile := os.ReadFile(filename)
	if ErrorReadFile != nil {
		return AvtotoAuf{}, ErrorReadFile
	}
	// Распарсить
	var ReturnObject AvtotoAuf
	ErrorUnmarshal := json.Unmarshal(fileBytes, &ReturnObject)
	if ErrorUnmarshal != nil {
		return AvtotoAuf{}, ErrorUnmarshal
	}
	return ReturnObject, nil
}
