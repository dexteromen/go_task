package main

import (
	"log"      // Package log implements simple logging
	"net/http" // Package http provides HTTP client and server implementations
	"time"     // Package time provides functionality for measuring and displaying time

	"github.com/gin-gonic/gin" // Gin is a web framework written in Go
	"github.com/google/uuid"   // Package uuid generates unique identifiers

	"example.com/tasks/models" // Importing the models package from the tasks module
)

func main() {
	// Connect to the database
	err := models.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Run database migrations
	err = models.AutoMigrate()
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	// Create a new Gin router
	r := gin.Default()

	// Define the routes and their handlers
	r.POST("/tasks", createTask)                 // Create a new task
	r.GET("/tasks", getTasks)                    // Get all tasks
	r.GET("/tasks/:id", getTask)                 // Get a task by ID
	r.PUT("/tasks/:id", updateTask)              // Update a task by ID
	r.DELETE("/tasks/:id", deleteTask)           // Delete a task by ID
	r.PUT("/tasks/:id/status", updateTaskStatus) // Update the status of a task by ID

	// Start the HTTP server
	r.Run()
}

func createTask(c *gin.Context) {
	// Declare a variable to hold the task data
	var task models.Task

	// Bind the JSON payload to the task variable
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new unique ID for the task
	task.ID = uuid.New().String()

	// Set the creation and update timestamps
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	// Create the task in the database
	if err := models.CreateTask(task); err != nil {
		// Check if the error is due to a duplicate task title
		if err.Error() == "task with the same title already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			log.Println("Error creating task:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Respond with the created task
	c.JSON(http.StatusCreated, gin.H{"message": "The task has been created", "task": task})
}

// getTasks handles the GET request to retrieve all tasks
func getTasks(c *gin.Context) {
	// Retrieve all tasks from the database
	tasks, err := models.GetTasks()
	if err != nil {
		// If there is an error, respond with an internal server error status
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Respond with the retrieved tasks
	c.JSON(http.StatusOK, tasks)
}

// getTask handles the GET request to retrieve a task by its ID
func getTask(c *gin.Context) {
	// Get the task ID from the URL parameter
	id := c.Param("id")

	// Retrieve the task from the database by its ID
	task, err := models.GetTaskByID(id)
	if err != nil {
		// If the task is not found, respond with a not found status
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Respond with the retrieved task
	c.JSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	// Get the task ID from the URL parameter
	id := c.Param("id")

	// Declare a variable to hold the task data
	var task models.Task

	// Retrieve the task from the database by its ID
	task, err := models.GetTaskByID(id)
	if err != nil {
		// If the task is not found, respond with a not found status
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Bind the JSON payload to the task variable
	if err := c.ShouldBindJSON(&task); err != nil {
		// If there is an error binding the JSON, respond with a bad request status
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the task ID and update the timestamp
	task.ID = id
	task.UpdatedAt = time.Now()

	// Update the task in the database
	if err := models.UpdateTask(task); err != nil {
		// If there is an error updating the task, respond with an internal server error status
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated task
	c.JSON(http.StatusOK, task)
}

func updateTaskStatus(c *gin.Context) {
	// Get the task ID from the URL parameter
	id := c.Param("id")

	// Declare a variable to hold the status update request data
	var request struct {
		Status string `json:"status"`
	}

	// Bind the JSON payload to the request variable
	if err := c.ShouldBindJSON(&request); err != nil {
		// If there is an error binding the JSON, respond with a bad request status
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the task status in the database
	if err := models.UpdateTaskStatus(id, request.Status); err != nil {
		// If there is an error updating the task status, respond with an internal server error status
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Task status updated successfully"})
}

func deleteTask(c *gin.Context) {
	// Get the task ID from the URL parameter
	id := c.Param("id")

	// Retrieve the task from the database by its ID to check if it exists
	_, err := models.GetTaskByID(id)
	if err != nil {
		// If the task is not found, respond with a not found status
		c.JSON(http.StatusNotFound, gin.H{"error": "Task Not Found!"})
		return
	}

	// Delete the task from the database by its ID
	if err := models.DeleteTask(id); err != nil {
		// If there is an error deleting the task, respond with an internal server error status
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
