// Go предлагает обширную поддержку для времени и
// продолжительности; вот несколько примеров.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Начнем с получения текущего времени
	now := time.Now()
	p(now)

	// Вы можете построить структуру `time`, указав год,
	// месяц, день и т.д. Время всегда связано с `местоположением`,
	// т.е. часовым поясом.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Вы можете извлечь различные компоненты значения времени.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Получения дня недели доступно через метод  `Weekday`.
	p(then.Weekday())

	// Эти методы сравниваются два раза, проверяя,
	// происходит ли первый случай до, после или
	// одновременно со вторым, соответственно.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Метод `Sub` возвращает `Duration`, интервал между
	// двумя временами.
	diff := now.Sub(then)
	p(diff)

	// Мы можем вычислить продолжительность.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Вы можете использовать `Add`, чтобы продвинуть
	// время на заданную продолжительность, или с `-`,
	// чтобы переместиться назад.
	p(then.Add(diff))
	p(then.Add(-diff))
}
