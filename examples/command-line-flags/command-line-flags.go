// [_Флаги командной строки_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// являются распространенным способом задания параметров
// для программ командной строки. Например, в `wc -l` - `-l`
// является флагом командной строки.

package main

// Go предоставляет пакет `flag`, поддерживающий базовый
// парсинг флагов командной строки. Мы будем использовать
// этот пакет для реализации нашего примера программы
// командной строки.
import (
	"flag"
	"fmt"
)

func main() {

	// Основные объявления флагов доступны для строковых,
	// целочисленных и логических параметров. Здесь мы
	// объявляем строковой флаг `word` со значением по
	// умолчанию `"foo"` и кратким описанием. Функция
	// `flag.String` возвращает строковый указатель (не
	// строковое значение); мы увидим, как использовать
	// этот указатель ниже.
	wordPtr := flag.String("word", "foo", "a string")

	// Объявляем флаги `numb` и `fork`, используя тот же
	// подход, что и выше.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// Также возможно вызвать метод, который использует
	// существующую переменную, объявленную в другом месте
	// программы. Обратите внимание, что в данном случае
	// необходимо передать указатель.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Как только все флаги объявлены, вызовите `flag.Parse()`,
	// чтобы выполнить парсинг командной строки.
	flag.Parse()

	// Здесь мы просто выведем результат парсинга и все
	// введеные аргументы. Обратите внимание, что нам
	// нужно разыменовать указатели, например, с
	// помощью `*wordPtr`, чтобы получить фактические
	// значения.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
