# Progress Journal

## 01-10-2020

- Basic structure, copying from my current company, I took part a bit on the research for this, we tried to implement Uncle Bob's SOLID architecture into golang.

- Learned how to do postgres connection. In general, using gorm as the ORM makes life easier because they have their own auto migration. I only learned the basics though.

- Learned that mux is more idiomatic than echo or gin, still don't really understand what is the meaning of "idiomatic".

- Learned using mux and separate the routes.

- Learned that docker expose PORT 9999 does not work, tried it for the db and it won't connect to dbeaver.

- Learned that if windows GOARCH is amd64, I need to use CGO_ENABLED=0 GOOS=linux on build. [More information.](https://stackoverflow.com/questions/20829155/how-to-cross-compile-from-windows-to-linux)

- Decided to make a journal everytime I work on this project.

I'll admit I haven't really grasp on how golang works, most of the codes I still took reference from others' people work.

## 08-10-2020

- Finally able to fully CRUD with database, got problem first because I forgot to set .env as "db" for db host, it is needed for docker-compose.

- Create utils library for response error and response with json.

- Still confused about best practice. From what I learned, there are so many ways to solve single issue in golang.

- Use gorm .First() instead of .Find() when looking for a single object, otherwise zero result will not return error.

- Always return correct http status code. [Reference](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml)

- Added DeletedAt, CreatedAt, UpdatedAt. Should have been from the start, if only I read the documentation thoroughly.

Decent progress I think, next time I will finish CRUD for supplier entity and trying gorm transaction for product entity.

## 14-10-2020

- Fixed bug on update when the id is not exist, I have to call .First() first to check if the data exist, and then actually update the data.

- Finished supplier module CRUD, there is not much of a difference as store module!!!, I wonder if abstracting these modules is the best way to do or keep them separated?
