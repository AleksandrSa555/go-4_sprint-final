package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, fmt.Errorf("неверное количество элементов")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("колчиство шагов должно быть больше 0")
	}
	durationSteps, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, durationSteps, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
