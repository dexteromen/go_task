# Go Language

### Go Gin Framework

https://gin-gonic.com/docs/

https://gorm.io/docs/index.html

#### Reference

https://blog.logrocket.com/rest-api-golang-gin-gorm/#what-is-gin

https://zaahidali.medium.com/mastering-rest-apis-with-the-go-gin-framework-b512ad785fdd

https://www.allhandsontech.com/programming/golang/web-app-sqlite-go/


# Task Manager API

The ***Task Manager API*** is a simple yet powerful API built using ***Go*** and ***Gin***. It provides endpoints to manage tasks with in-memory storage. This project is designed to help you learn Go language and understand the fundamentals of building APIs.

## Features

- Get a list of all tasks.
- Get the details of a specific task.
- Update a specific task.
- Delete a specific task.
- Create a new task.
- Update the status of the specific  task.

## API Endpoints

- `GET /tasks`
  - Get a list of all tasks.

- `GET /tasks/:id`
  - Get the details of a specific task.

- `PUT /tasks/:id`
  - Update a specific task.
  - Request Body: JSON object with the updated task details.

- `DELETE /tasks/:id`
  - Delete a specific task.

- `POST /tasks`
  - Create a new task.
  - Request Body: JSON object with the task's title, description, due date, and status.

- `PUT /tasks/:id/status`
  - Updates the status of the task .
  - Request Body: JSON object with the task's status.

### Examples

#### GET /tasks
- Send a GET request to `http://localhost:8080/tasks` to retrieve all tasks.

#### GET /tasks/:id
- Send a GET request to `http://localhost:8080/tasks/:id` to retrieve details of a specific task.

#### PUT /tasks/:id
- Send a PUT request to `http://localhost:8080/tasks/:id` with a JSON body containing the updated task details.

#### DELETE /tasks/:id
- Send a DELETE request to `http://localhost:8080/tasks/:id` to delete a specific task.

#### POST /tasks
- Send a POST request to `http://localhost:8080/tasks` with a JSON body containing the task's details to create a new task.

#### PUT /tasks/:id/status
- Send a POST request to `http://localhost:8080/tasks/:id/status` with a JSON body containing the task's status detail.