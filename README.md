# сервис для регистрации пользователей

# запуск
Для старта:
```
docker compose build

docker compose up
```
# регистрация
Для регистрации нужно послать **POST** запрос на метод **/user/signin**
![image](https://github.com/user-attachments/assets/876794e5-9617-443f-8b3b-aed402987254)

Реализована проверка емейла через регулярные выражения

Сервис ответит что нужно подверждение емейла

**в логах сервиса можно найти код для подверждения емейла**
![image](https://github.com/user-attachments/assets/b98f8017-e02d-4fa2-a819-181e52f2c357)


# подтверждение
Для подверждения нужно послать **GET** запрос на метод **/user/confirm_email** с параметрами запроса
пример:
**http://localhost:6666/user/confirm_email?mail=danil@mail.ru&code=762354**

Дервис ответит
![image](https://github.com/user-attachments/assets/c43793ef-cc8d-4ba8-8cb2-3fe104e7af2a)

# база данных
В базе есть таблица users
![image](https://github.com/user-attachments/assets/aec42aa9-428d-47a8-b2e1-3f7cb30eb604)

Пока емейл не подвержден, юзер будет со статусом 2(awaiting_confirmation)
![image](https://github.com/user-attachments/assets/4c0feca0-b52a-4cb2-88a4-a19d91feada7)

