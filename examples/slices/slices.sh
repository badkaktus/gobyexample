# Обратите внимание, несмотря на то что срезы
# являются отдельным типом данных, отображаются
# они так же как массивы командой `fmt.Println`.
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Посмотрите этот [отличный пост](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html),
# написаный командой Go, чтобы узнать больше о разработке
# и использовании срезов в Go.

# Теперь, когда мы рассмотрели массивы и срезы, мы
# посмотрим другую ключевую встроенную структуру
# данных Go: карты.
