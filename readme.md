# Echomers

## Description

This is a simple Rest Api created with go and echo framework.

### Technologies

- [Go](https://golang.org/)
- [Echo](https://echo.labstack.com/)
- [Gorm](https://gorm.io/)
- [Mysql](https://www.mysql.com/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)


## Features

- Create a user
- Get all users
- Get a user by id
- Update a user
- Delete a user

## Installation

1. Clone the repository
2. Run `docker-compose up` to start the server
3. The server will be running on `http://localhost:8080`

## Usage

### Create a user

```bash
curl -X POST http://localhost:8080/users -d '{"name": "John Doe", "email": " [email protected]"}'
```

### Get all users

```bash
curl http://localhost:8080/users
```

### Get a user by id

```bash
curl http://localhost:8080/users/1
```

### Update a user

```bash
curl -X PUT http://localhost:8080/users/1 -d '{"name": "Jane Doe", "email": " [email protected]"}'
```' | echo -n 1 | sha256sum
```

### Delete a user

```bash
curl -X DELETE http://localhost:8080/users/1
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
