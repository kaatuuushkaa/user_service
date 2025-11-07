# User Service

Сервис для управления пользователями, их поинтами и выполнением заданий. Пользователь может получать награды за выполнение задач (подписки, приглашения и т.д.) и за ввод реферального кода.

---

## Основные возможности
- `POST /login` - получение access token(JWT) для дальнейшей авторизации пользователя
<img width="750" height="595" alt="image" src="https://github.com/user-attachments/assets/44296170-3935-4266-95d2-3596ced7845b" />
Обработана ситуация с несуществующим пользователем
<img width="752" height="569" alt="image" src="https://github.com/user-attachments/assets/6d796242-e3b0-4255-88f7-9ae351b4caca" />

- `GET /users/{id}/status` - получение всей доступной информации о пользователе
<img width="756" height="614" alt="image" src="https://github.com/user-attachments/assets/58bc33c3-e443-4cfb-a456-515e072aa098" />
С полученным токеном мы можем посмотреть информацию только о том пользователе, которому принадлежит токен
<img width="751" height="554" alt="image" src="https://github.com/user-attachments/assets/b6b9cd27-b101-4c5a-9ad5-1dfd515c5ea7" />

- `GET /users/leaderboard` - получение топа пользователей с самым большим балансом
<img width="751" height="749" alt="image" src="https://github.com/user-attachments/assets/f8bda2e4-8e1e-4318-9abe-7e5ed8b96965" />

- `POST /users/{id}/task/complete` - выполнение заданий
<img width="747" height="622" alt="image" src="https://github.com/user-attachments/assets/f6687318-8f59-4f91-ba64-2c512a0eca79" />
Если пользователя не существует - получим ошибку
<img width="756" height="626" alt="image" src="https://github.com/user-attachments/assets/b3082b6c-ddf6-413b-8ee2-a588dad7b22d" />

- `POST /users/{id}/referrer` - ввод реферального кода (может быть id другого пользователя)
<img width="749" height="756" alt="image" src="https://github.com/user-attachments/assets/c825060a-4db5-46b3-88a9-8a9513fc5c54" />
Реализована идемпотентность работы с реферальным кодом. При повторном вызове, баллы не изменятся и придет уведомление, что пользователь уже использовал рефералку
<img width="753" height="572" alt="image" src="https://github.com/user-attachments/assets/f9fb1cb7-8803-4866-985d-286ee056b462" />
Пользователь не может зарефералить себя сам
<img width="746" height="582" alt="image" src="https://github.com/user-attachments/assets/5d137678-5584-425f-9cbc-e55016535ef3" />

# Стек
Go, gin, gorm, jwt, docker-compose, golang-migrate, postgresql

# Запуск проекта
`docker-compose up --build`

Запустится 2 контейнера
<img width="943" height="115" alt="image" src="https://github.com/user-attachments/assets/4825d577-4ee0-4a84-bfc3-6d3fa7e7152b" />
Используем команду миграции для создания таблиц и внесения данных

`make migrate`


