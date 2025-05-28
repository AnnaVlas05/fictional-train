package trainings

import(
	"time"
	"github.com/AnnaVlas05/fictional-train/internal/personaldata"
	"github.com/AnnaVlas05/fictional-train/internal/spentenergy"
	"strconv"
	"strings"
	"fmt"
	"errors"
)

type Training struct {
	// TODO: добавить поля
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат строки: ожидалось 3 элемента")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return errors.New("ошибка парсинга количества шагов: " + err.Error())
	}

	if steps <= 0 {
		return errors.New("количество шагов должна быть больше нуля")
	}
	t.Steps = steps

	t.TrainingType = strings.TrimSpace(parts[1])

	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return errors.New("ошибка парсинга длительности: " + err.Error())
	}

	if duration <= 0 {
		return errors.New("длительность тренировки должна быть больше нуля")
	}
	t.Duration = duration

	return nil
}
func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if t.Duration <= 0 {
		return "", errors.New("длительность тренировки должна быть больше нуля")
	}

	distance := spentenergy.Distance(t.Steps, t.Height)

	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	durationHours := t.Duration.Hours()

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		durationHours,
		distance,
		avgSpeed,
		calories,
	)

	return result, nil
}
