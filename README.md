# User Service

Сервис для управления пользователями, их поинтами и выполнением заданий. Пользователь может получать награды за выполнение задач (подписки, приглашения и т.д.) и за ввод реферального кода.

---

## Основные возможности
- `POST /auth/signup` - регистрация пользователя
<img width="761" height="578" alt="image" src="https://github.com/user-attachments/assets/d6f03890-3c6a-4bf6-bf49-2cfbf42cd989" />


- `POST /auth/signin` - получение access token(JWT) для дальнейшей авторизации пользователя
<img width="758" height="612" alt="image" src="https://github.com/user-attachments/assets/791a4f64-e190-4c37-abbc-4b603bd50e68" />

Обработана ситуация с несуществующим пользователем или неправильным паролем
<img width="761" height="574" alt="image" src="https://github.com/user-attachments/assets/3b0e0852-8ead-4252-a471-62e3484ff102" />

- `GET /users/{id}/status` - получение всей доступной информации о пользователе
<img width="757" height="572" alt="image" src="https://github.com/user-attachments/assets/0f5a7a5a-0d25-4d97-b6a3-e24edd8cacdd" />

- `GET /users/leaderboard` - получение топа пользователей с самым большим балансом
<img width="757" height="642" alt="image" src="https://github.com/user-attachments/assets/d3d34936-59d0-4471-b11a-ef04447d222c" />



- `POST /users/{id}/task/complete` - выполнение заданий
<img width="757" height="598" alt="image" src="https://github.com/user-attachments/assets/2fb414b4-843d-406f-a853-a44d9cb7c105" />

Если пользователя не существует - получим ошибку
<img width="749" height="508" alt="image" src="https://github.com/user-attachments/assets/b846747b-01df-475c-93c0-6ea34753376a" />


- `POST /users/{id}/referrer` - ввод реферального кода (может быть id другого пользователя)
<img width="749" height="651" alt="image" src="https://github.com/user-attachments/assets/344d0519-1eab-428d-8bd0-2c35263ebdb2" />

Реализована идемпотентность работы с реферальным кодом. При повторном вызове, баллы не изменятся и придет уведомление, что пользователь уже использовал рефералку
<img width="755" height="530" alt="image" src="https://github.com/user-attachments/assets/2378d483-b520-4d78-b814-dd14b60bc553" />

Пользователь не может зарефералить себя сам
<img width="762" height="572" alt="image" src="https://github.com/user-attachments/assets/c3982c32-abaa-4ffc-a551-1f815d50b95a" />


# Стек
Go 1.23+
Gin (github.com/gin-gonic/gin)
Gorm Driver Postgres (gorm.io/driver/postgres)
JWT (github.com/golang-jwt/jwt/v5)
BCrypt (golang.org/x/crypto/bcrypt)

# Запуск проекта
`docker-compose up --build`

Запустится 2 контейнера
<img width="943" height="115" alt="image" src="https://github.com/user-attachments/assets/4825d577-4ee0-4a84-bfc3-6d3fa7e7152b" />
Используем команду миграции для создания таблиц и внесения данных

`make migrate`

# Идеи для доработки
- добавить роли пользователей для разграничения права доступа
- middleware для проверки роли
- CRUD для задач
- добавить Redis для хэширования часто запрашиваемых данных
- добавить Kafka для обеспечения пропускной способности, есди запросы будут расти
