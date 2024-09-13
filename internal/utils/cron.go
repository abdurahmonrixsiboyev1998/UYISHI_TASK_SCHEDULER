package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

type CronParser struct {
	cronSchedule cron.Schedule
}

func NewCronParser(cronExpression string) (*CronParser, error) {
	schedule, err := cron.ParseStandard(cronExpression)
	if err != nil {
		return nil, err
	}

	return &CronParser{cronSchedule: schedule}, nil
}

func (cp *CronParser) Next(from time.Time) time.Time {
	return cp.cronSchedule.Next(from)
}

func ValidateCronExpression(cronExpression string) error {
	fields := strings.Fields(cronExpression)
	if len(fields) != 5 {
		return errors.New("invalid cron expression: must contain 5 fields")
	}
	return nil
}

func Example() {
	cronExpr := "0 10 * * *" 
	if err := ValidateCronExpression(cronExpr); err != nil {
		fmt.Println("Error:", err)
		return
	}

	parser, err := NewCronParser(cronExpr)
	if err != nil {
		fmt.Println("Error creating CronParser:", err)
		return
	}

	nextRun := parser.Next(time.Now())
	fmt.Println("Next run:", nextRun)
}
