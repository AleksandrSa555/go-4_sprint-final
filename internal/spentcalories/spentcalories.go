package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	sliceData := strings.Split(data, ",")
	if len(sliceData) != 3 {
		return 0, "", 0, fmt.Errorf("неверное количество элементов")
	}
	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	durationTraining, err := time.ParseDuration(sliceData[2])
	if err != nil {
		return 0, "", 0, err
	}
	if durationTraining <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	return steps, sliceData[1], durationTraining, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)

	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, training, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64
	dist := distance(steps, height)
	avgSpeed := meanSpeed(steps, height, duration)

	switch training {
	case "Ходьба":

		calories, err = WalkingSpentCalories(steps, weight, height, duration)

	case "Бег":

		calories, err = RunningSpentCalories(steps, weight, height, duration)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		log.Println(err)
		return "", err
	}

	res := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		training, duration.Hours(), dist, avgSpeed, calories)
	return res, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	avgSpeed := meanSpeed(steps, height, duration)

	calories := (weight * avgSpeed * duration.Minutes()) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	avgSpeed := meanSpeed(steps, height, duration)

	calories := (weight * avgSpeed * duration.Minutes()) / minInH

	return calories * walkingCaloriesCoefficient, nil
}
