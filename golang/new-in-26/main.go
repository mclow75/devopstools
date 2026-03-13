package main

import (
	"encoding/json"
	"fmt"
)

// Не требуется заведение переменной и указателя на неё
type User struct {
	Username   string
	RetryLimit int
	IsActive   bool
}

type UpdateSettingsParams struct {
	RetryLimit *int  `json:"retry_limit,omitempty"`
	IsActive   *bool `json:"is_active,omitempty"`
}

// Если настройки пользователя не равны конфигурации, то обновляем
func PatchUser(u *User, updates UpdateSettingsParams) {
	if updates.RetryLimit != nil {
		u.RetryLimit = *updates.RetryLimit
	}
	if updates.IsActive != nil {
		u.IsActive = *updates.IsActive
	}
}

func main() {
	// Новая инициализация
	// Создаем указаль на область памяти типа int32 или float64, инициализированная нулём
	p := new(float64)
	fmt.Printf("Значение при инициализации: %.2f\n", *p)
	// Разименовываем указатель и присваиваем значение переменной 42.22
	*p = 42.22
	fmt.Printf("Значение после присваивания значения: %.2f\n", *p)
	// Не требуется отдельная операция присваивания, можно сразу инициализировать
	d := new(5.76)
	fmt.Printf("Значение: %.2f, Адрес: %p\n", *d, d)

	dbUser := User{Username: "Alice", RetryLimit: 5, IsActive: false}
	// До 1.26 пришлось бы создавать переменную и указатель на неё
	// newRetryLimit := 10
	// newIsActive := true
	// PatchUser(&dbUser, UpdateSettingsParams{
	// 	RetryLimit: &newRetryLimit,
	// 	IsActive:   &newIsActive,
	// })
	PatchUser(&dbUser,
		UpdateSettingsParams{
			RetryLimit: new(20),
			IsActive:   new(true),
		})
	data, _ := json.Marshal(dbUser)
	fmt.Println(string(data))
}
