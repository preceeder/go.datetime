package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestDateTime_DiffYears(t1 *testing.T) {
	birthday, _ := Parse("1996-02-09", time.DateOnly)
	now := Now()
	fmt.Println(birthday)
	fmt.Println(birthday.Age())
	fmt.Println(birthday.Constellation())
	fmt.Println(birthday.DiffSeconds(), now.DiffSeconds(birthday), now.DiffSeconds())
	fmt.Println(birthday.DiffMinutes(), now.DiffMinutes(birthday), now.DiffMinutes())
	fmt.Println(birthday.DiffHours(), now.DiffHours(birthday), now.DiffHours())
	fmt.Println(birthday.DiffDays(), now.DiffDays(birthday), now.DiffDays())
	fmt.Println(birthday.DiffWeeks(), now.DiffWeeks(birthday), now.DiffWeeks())
	fmt.Println(birthday.DiffYears(), now.DiffYears(birthday), now.DiffYears())
	fmt.Println(birthday.DiffMouths(), now.DiffMouths(birthday), now.DiffMouths())

	fmt.Println(TimestampToDateTime(birthday.Timestamp()))
}
