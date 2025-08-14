package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// Инициализация генератора случайных чисел (обязательно!)
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/rand", Number)
	http.ListenAndServe(":9091", nil)
}

func Number(w http.ResponseWriter, r *http.Request) {
	randomnumber := rand.Intn(7) // Генерируем число от 0 до 6

	// Преобразуем число в строку и отправляем клиенту
	w.Write([]byte(strconv.Itoa(randomnumber)))

	// Дополнительно выводим в консоль (необязательно)
	fmt.Println(randomnumber)

}
