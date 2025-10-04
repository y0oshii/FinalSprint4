package daysteps

import (
	"fmt"
	"log"
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
		return 0, 0, fmt.Errorf("неправильные параметры %s", data)
	}

	steps, err := strconv.Atoi(activity[0])
	if err != nil || steps <= 0 {
		return 0, 0, fmt.Errorf("некорректное число шагов %d", steps)
	}

	dur, err := time.ParseDuration(activity[1])
	if err != nil || dur <= 0 {
		return 0, 0, fmt.Errorf("некорректное время %s", dur)
	}

	return steps, dur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, dur, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 || dur <= 0 {
		log.Println(err)
		return ""
	}

	length := float64(steps) * stepLength
	km := length / 1000

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, dur)
	if err != nil {
		log.Println(err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, km, calories)

	return result
}
