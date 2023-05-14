# go-cassandra-api
Projekt na studia jest to API napisane w Golangu dla Cassandry - pozwala na dodawanie użytkowników, wyświetlanie ich, głosowanie na uniwersytety czy wyszukiwanie ich po nazwie bądź po typie kierunku. 

###  Zarządzanie użytkownikami
**Wyświetalnie wszystkich dodanych użytkowników:**
```bash
curl localhost:8080/users
```

Output:

```json
  {
    "user_id": "41167c0e-f25a-11ed-8246-2ed8f8039cac",
    "email": "user1@gmail.com",
    "login": "user1",
    "password": "$2a$10$ctL5wMiS/Sr1.rbvfEuRK.sIjARhgWWFeCvuvcgK9kjvOtbbIB8we"
  },
```

**Wyświetlanie specyficznego użytkownika**
```bash
curl localhost:8080/users/472189c2-f25a-11ed-8247-2ed8f8039cac
```
Output:
```json
  {
    "user_id": "472189c2-f25a-11ed-8247-2ed8f8039cac", 
    "email": "user2@gmail.com", 
    "login": "user2", 
    "password": "$2a$10$s1V/kEnXz.q/wUXNHzkBuua.LxE5VloT9z5pKAFGuWoAg4Yr56DIe"
  }
```
***Dodanie użytkownika do bazy***
```bash
curl -X POST \
  http://localhost:8080/adduser \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "test@user.com",
    "login": "test",
    "password": "password"
  }'
```
Output:
```json
  {
    "message":"User created",
    "user_uid":"600c3ac2-f26e-11ed-9538-2ed8f8039cac" (utworzyłem już tego użytkownika)
  }
```
Jeśli użytkownik istnieje (login lub/oraz e-mail już jest w bazie
```json
  {
    "error": "user already exists"
  }
```
***Logowanie*** 
```bash
  curl -X GET \
  http://localhost:8080/login \
  -H 'Content-Type: application/json' \
  -d '{
    "login": "test@user.com",
    "password": "password"
  }'
```
Output:
```json
    {
      "message": "User logged in"
    }
```
