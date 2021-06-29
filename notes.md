
# postgres connection
[pgx postgres driver](https://habr.com/ru/company/oleg-bunin/blog/461935/)
    buff size limit 4KB!
    
# pretty project for example
[photoprism](https://github.com/photoprism/photoprism)

# project layout
[project layout](https://github.com/golang-standards/project-layout)

# awesome go
[awesome go](https://awesome-go.com/)

#go jwt 
[go jwt](https://github.com/appleboy/gin-jwt) 

# go echo
[go echo](https://echo.labstack.com/guide/)

# sqlx
[sqlx](https://github.com/jmoiron/sqlx)
# Plan
## Todo 	
- API
- Documentation
- Tests - show example how to write test for routes
- Business logic
- Migrations
- Basic role system
- swag docs can't set query params with string\[string]=string

## Resolved
- routes: create, login
- middleware: jwt
- business logic
    - authorization
    - model of user, workout, workout_type, workout_value
    - password checking
- sqlx: extension of database/sql for better life
- basic query builder (find, orderBy, offsetRows, andWhere, maxResult method in repository) (deprecated)
- in FindBy - set flexible field selection (deprecated)
- parser query parameters (deprecated)
#Business logic
### basic entities
 - user
 - workout
 - workout type
 - workout value
## Workout scheme
##### workout type - `(run, gym)` and
##### workout template - `(workout scheme of certain units for workout type)`
have a finite set of values that can be stored in the cache (redis).
##### workout_template - `will store in map and maybe will move in db or cache`

#API
 - https
 - routes
 - cors
 - http 2.0
 
# In progress
 - CRUD
   - elements to make CRUD
      - entities
         - workout
         - workout_type 
         - workout_value
      - controllers
         - workout
         - workout_type
         - workout_value
      - routes
         - GET
         - POST
         - PUT
         - DELETE
      - repositories
         - workout
         - workout_type
         - workout_value
      - migration
         - create tables
            - workout
            - workout_type
            - workout_value
 - JWT
   - create secret key in env
 - Logging
   - use echo
 - describe a thing like DTO in symfony (to hide private fields as password)
 - describe methods for basic api functional 
     - getById, update, sort, filter, per page (offset),        
     - get all
     - get by id
     - get by a field ?
     - order by
     - filter