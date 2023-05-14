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
  { "error": "user already exists" }
```
***Logowanie*** 
```bash
  curl -X GET \
  http://localhost:8080/login \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "test@user.com",
    "password": "password"
  }'
```
Output:
```json
    { "message": "User logged in" }
```
### Głosowanie
```bash
curl -X POST \
  http://localhost:8080/vote \
  -H 'Content-Type: application/json' \
  -d '{
    "login": "user1",
    "university_name": "Collegium da Vinci"
}'
```
Output: 
```json
    { "message": "Vote created" }
```
Sprawdzane błędy:
```
    { "error": "user already voted for this university" } # user już głosował
    { "error": "user does not exist" } # user nie istnieje
    { "error": "university does not exist" } # podany uniwersytet nie istnieje
```
***Przeglądanie swoich głosów***
```bash
curl -X GET \
  http://localhost:8080/myvotes \
  -H 'Content-Type: application/json' \
  -d '{
    "login": "user1"
}'
```
Output
```json
[
    {
        "voted_id": "31ffc1bf-29f6-4e67-a435-aeb02c7944bd",
        "login": "user1",
        "university_name": "Collegium da Vinci",
        "voted_on": "2023-05-14 15:56:39.062 +0000 UTC"
    },
    {
        "voted_id": "01c3d675-f778-49a1-abfb-35486ebb870b",
        "login": "user1",
        "university_name": "Politechnika Poznanska",
        "voted_on": "2023-05-14 16:02:40.765 +0000 UTC"
    }
```
***Sprawdzanie dodanych uniwersytetów***
```bash
curl -X GET \
  http://localhost:8080/universities \
  -H 'Content-Type: application/json'
```
Output:
```json
[
    {
        "university_id": "5d913fb0-b91d-4bda-95d4-e899fc4ef840",
        "city": "Oxford",
        "country": "UK",
        "name": "University of Oxford"
    },
    {
        "university_id": "9f7ae968-f261-11ed-add5-2ed8f8039cac",
        "city": "Poznan",
        "country": "Poland",
        "name": "Politechnika Poznanska"
    },
```
