# Telegram-бот

Телеграм бот, реагирующий на входные команды и выдающит ответ по заданной команде.

Бот поддерживает следующие команды:

- /ping : бот отвечает строкой pong.
- /ping/my : бот отвечает пользователю строкой с счетчиком пингов (сколько пингов отправилась от данного юзера за все время)
- /time : бот отвечает текущим временем.

# Установка и запуск

Для запуска бота (проекта) необходимо иметь:

- Установленный Go (версии 1.24.2)
- Docker
- Доступ к командной строке

### Склонируйте репозиторий в удобную папку:
```bash
git clone https://github.com/dltzk/intership-lamp-labs-2025.git
```

Проект уже готов к запуску, но мы чето забыли..
Ёшки матрёшки, нам нужно создать бота и заполучить его токен!

- Чтобы создать бота, переходим в телеграм и ищем бота @BotFather

![Bot screenshot](https://sun9-35.userapi.com/impg/OLtmjs3OFaWwwslM0OXjgIlcFKUi0Cdu-iGe6g/Y8oLYP4cngw.jpg?size=592x113&quality=95&sign=a398017aee51f6408c9fb9a5fbc44e7d&type=album)

- Жмем /start, бот выдает нам все доступные команды:

![Commands screenshot](https://sun9-55.userapi.com/impg/SA8ZP0fCuFpq-ogYAfElWHK7o3-AYqDWW9Ak_Q/MYNRSbR2Tv0.jpg?size=474x783&quality=95&sign=5c5281af5664035d39156aa5f621afc4&type=album)

- Вводим команду /newbot, называем его, и получаем от бота следующее сообщение:

![Answer screenshot](https://sun9-13.userapi.com/impg/4-wcHsNO20yP9ZiQh_CK36b1q8j2sTe9OFJ_uA/TLmVVMAwu_U.jpg?size=476x516&quality=95&sign=11cc28608a90b8bfe723d46e59dea952&type=album)

Это значит, что мы сделали все верно, бот создан, и теперь нужно скопировать токен и записать его в нужный файлик.

- В корневом разделе проекта создаем файл .env, и записываем туда:
```bash
TELEGRAM_TOKEN=ВАШ_ТОКЕН
```

Все готово, теперь переходим в папку deployment и выполняем deploy.ps1:
```bash
cd deployment
.\deploy
```

Когда сбилдится и запустится, можно переходить в личку бота и писать доступные команды.
Пример вызова команд:

![Example screenshot](https://sun9-67.userapi.com/impg/nFajbp0CIrUbbc6B04kh5WMSEqjlUbRfqJifSg/bxaMqv8107A.jpg?size=261x774&quality=95&sign=8292bf7022479af8df96a258b5343929&type=album)
