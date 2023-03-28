# API para a faculdade Passei Direto com GOLANG

A aplicação foi desenvolvida em **Go** com banco de dados  **Postgres**. Ela possui um docker-compose para build do banco de dados **Postgres** e do administrador **pgAdmin**. Além de contar com a documentação construida com o **Swagger**.

### Run Database

```
$ docker-compose up --build
```

### Run API 

```
$ go run main.go
```

### API Docs

- JSON file at docs/swagger.json

##### A API está disponível na porta 5000
 - http://localhost:5000/api/v1/...

##### O swagger está disponível na porta 5000
- http://localhost:5000/swagger/index.html#/

##### O pgAdmin está disponível na porta 9000
- http://localhost:900
- User: admin@admin.com
- Password: 123456

![]


Os filtros solicitados estão contidos nos controladores de **Aluno** e  **Disciplina**, como está exibido nas imagens abaixo.

![]

![]