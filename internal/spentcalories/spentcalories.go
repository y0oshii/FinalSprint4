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
	train := strings.Split(data, ",")
	if len(train) != 3 {
		return 0, "", 0, fmt.Errorf(data)
	}

	steps, err := strconv.Atoi(train[0])
	if err != nil {
		return 0, "", 0, err
	}

	duration, err := time.ParseDuration(train[2])
	if err != nil {
		return 0, "", 0, err
	}

	return steps, train[1], duration, nil
}

func distance(steps int, height float64) float64 {
	length := height * lenStep
	distanceKm := (length * float64(steps)) / mInKm

	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}

	dist := distance(steps, height)
	hour := duration.Hours()
	speed := dist / hour

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, train, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var dur, dist, speed, cal float64
	switch train {
	case "Ходьба":
		dur = float64(duration.Hours())
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		cal, err = WalkingSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}

	case "Бег":
		dur = float64(duration.Hours())
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		cal, err = RunningSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}

	default:
		log.Println("неизвестный тип тренировки")
	}

	result := fmt.Sprintf("Тип тренировки: %s \nДлительность: %.2f ч. \nДистанция: %.2f км. \nСкорость: %.2f км/ч \nСожгли калорий: %.2f", train, dur, dist, speed, cal)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	speed := meanSpeed(steps, height, duration)
	if speed <= 7 || speed > 45 {
		return 0, fmt.Errorf("Неккоректная скорость %2.f", speed)
	}

	calories := (weight * speed * duration.Minutes()) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	speed := meanSpeed(steps, height, duration)
	if speed <= 0 || speed > 7 {
		return 0, fmt.Errorf("Неккоректная скорость %2.f", speed)
	}

	calories := ((weight * speed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient

	return calories, nil
}
