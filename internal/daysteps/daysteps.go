package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	activity := strings.Split(data, ",")
	if len(activity) != 2 {
		return 0, 0, fmt.Errorf("Неправильные параметры %s", data)
	}

	steps, err := strconv.Atoi(activity[0])
	if err != nil {
		return 0, 0, err
	}

	dur, err := time.ParseDuration(activity[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, dur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, dur, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 {
		return ""
	}

	length := float64(steps) * stepLength
	km := length / 1000

	calories, err := spentcalories.WalkingSpentCalories(steps, height, weight, dur)

	if err != nil {
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d. \nДистанция составила %.2f км. \nВы сожгли %.2f ккал.", steps, km, calories)

	return result
}
