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
	sliceData := strings.Split(data, ",")
	if len(sliceData) != 2 {
		return 0, 0, fmt.Errorf("неверное количество элементов")
	}
	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("колчиство шагов должно быть больше 0")
	}
	durationSteps, err := time.ParseDuration(sliceData[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, durationSteps, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, durationSteps, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distance := (float64(steps) * stepLength) / mInKm
	calories, err := WalkingSpentCalories(steps, weight, height, durationSteps)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)

}
