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