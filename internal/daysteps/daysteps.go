package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/AnnaVlas05/fictional-train/internal/personaldata"
	"github.com/AnnaVlas05/fictional-train/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("incorrect data format: two elements are expected")
	}

	//
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("error when parsing the number of steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("the number of steps must be greater than 0")
	}

	//
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("error when parsing the duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("the duration must be greater than 0")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distanceKm := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distanceKm, calories,
	)
	return result, nil
}
