Tech Stacks
-
- GO >1.22
- Docker
- Postgres running in Docker
- Swagger for docs
- GOlang migrate for migrations
- chi - go get -u github.com/go-chi/chi/v5
- air-verse/air - hot reload
- direnv - load and unload environment variables
- db migration: https://github.com/golang-migrate/migrate
- go validator: https://github.com/go-playground/validator

Database ORM (Alternatives)
-
-  GOORM
-  sqlboiler
-  goose - another db migration tool
  

Architechture concepts:
-

1. ***Separation of concerns:*** each layer in your program should be separate by a clearn barrier the 
     + transport layer, 
     + the service layer, 
     + the storage layer

2. ***Dependency Inversion Principle(DIP):*** Injecting the dependencies in layers. Don't directly call them. it Promotes loose coupling and makes it easier to test programs

3. ***Adaptablity to Chnage:*** By organizing code in modular and flexible way, we can more easily introduce new features, refactor existing code, and respond to evolvong business requirements. system should be easy to change when it comes to add new features
   
4. ***Focus on Business Value:*** 
   

Layers
-
| TRANSPORT | (HTTP HANDLER)
| SERVICE   | 
| STORAGE   | (REPOSITORY) -> DB


Commands
-
```
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
```

```
migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" up
```

Resources
-

+ https://www.postgresql.org/docs/current/citext.html
