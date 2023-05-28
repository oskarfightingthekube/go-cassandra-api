# go-cassandra-api
Projekt na studia jest to API napisane w Golangu dla Cassandry - pozwala na dodawanie użytkowników, wyświetlanie ich, głosowanie na uniwersytety czy wyszukiwanie ich po nazwie bądź po typie kierunku. 
## Jak odpalić?
Wystarczy wejść w folder repo i wpisać 
```golang
go run main.go
```
API będzie nasłuchiwać na portcie :8080 (localhost:8080)
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
    "password": "insert_pasword",
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
    "email": "user1", # wpisujemy email, ale trzeba podać login - do poprawki
    "password": "password"
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
***Dodawanie Uniwersytetów***
```bash
  curl -X POST \
  http://localhost:8080/adduniversity \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Uniwersytet im. Adama Mickiewicza",
    "country": "Poland",
    "city": "Poznan"
  }'
```
Funkcja oczywiście sprawdza, czy Uniwersytet już istnieje.

***Szukanie uniwesytetu po typie studiów**

```bash
  curl -X GET \
  'http://localhost:8080/majors' \
  -H 'Content-Type: application/json' \
  -d '{
    "type": "Technical"
  }'
```
Output:
```json
[
    {
        "major_id": "a60e2746-3aa8-468a-a2a2-6cff7c7878ac",
        "name": "Engineering",
        "type": "Technical",
        "university_id": "0cda384f-292d-4978-861c-8661855e22a7",
        "university_name": "Massachusetts Institute of Technology"
    },
    {
        "major_id": "16f6c8bd-335f-4482-8a54-c024ca83a1ac",
        "name": "Computer Science",
        "type": "Technical",
        "university_id": "5d913fb0-b91d-4bda-95d4-e899fc4ef840",
        "university_name": "University of Oxford"
    }
]
```
***Szukanie uniwersytetu po typie oferowanych kierunków***
```bash
curl -X GET \
  'http://localhost:8080/majors/name' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Computer Science"
        }'
```
Output:
```json
[
    {
        "major_id": "16f6c8bd-335f-4482-8a54-c024ca83a1ac",
        "name": "Computer Science",
        "type": "Technical",
        "university_id": "5d913fb0-b91d-4bda-95d4-e899fc4ef840",
        "university_name": "University of Oxford"
    }
]
```
***Szukanie kierunków studiów po nazwie Uniwesytetu***
```bash
  curl -X GET \
  'http://localhost:8080/departments/university' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Harvard University"
}'
```
Output:
```
[
    {
        "department_id": "0356d901-dc2a-4414-bed3-2fa2234d3938",
        "name": "Law",
        "university_id": "b7b4050a-4e66-4cba-badf-846f1d07f0c5",
        "university_name": "Harvard University"
    },
    {
        "department_id": "36dfaccc-4a07-401c-af74-f66198232b4a",
        "name": "Business",
        "university_id": "b7b4050a-4e66-4cba-badf-846f1d07f0c5",
        "university_name": "Harvard University"
    },
    [...]
```
***Wyświetlnie wszystkich wydziałów***
```bash
 curl -X GET \
  'localhost:8080/departments' \
  -H 'Content-Type: application/json' 
```
Output:
```json
[
    {
        "department_id": "0356d901-dc2a-4414-bed3-2fa2234d3938",
        "name": "Law",
        "university_id": "b7b4050a-4e66-4cba-badf-846f1d07f0c5",
        "university_name": "Harvard University"
    },
    {
        "department_id": "36dfaccc-4a07-401c-af74-f66198232b4a",
        "name": "Business",
        "university_id": "b7b4050a-4e66-4cba-badf-846f1d07f0c5",
        "university_name": "Harvard University"
    },
[...]
```

