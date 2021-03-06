# Когда мы запускаем программу, сообщение `"ping"` успешно
# передается из одной горутины в другую по нашему каналу.
$ go run channels.go 
ping

# По-умолчанию, отправление и получение блокируются,
# пока оба отправителя и получателя готовы. Это
# свойство позволило нам ждать в конце нашей программы
# сообщения `"ping"` без использования какой-либо
# другой синхронизации.
