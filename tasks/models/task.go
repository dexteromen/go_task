package models

import (
	// "fmt"
	"log"
	"time"
)

// Task represents a task with its details.
type Task struct {
	ID          string    `json:"id" gorm:"primaryKey"`                   // Unique identifier for the task
	Title       string    `json:"title" binding:"required" gorm:"unique"` // Title of the task, required field with unique constraint
	Description string    `json:"description"`                            // Description of the task
	DueDate     string    `json:"due_date" binding:"required"`            // Due date of the task, required field
	Status      string    `json:"status"`                                 // Status of the task
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`       // Timestamp when the task was created
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`       // Timestamp when the task was last updated
}

// AutoMigrate automatically migrates the database schema for the Task model.
func AutoMigrate() error {
	// Attempt to auto-migrate the Task model schema in the database.
	err := DB.AutoMigrate(&Task{})
	if err != nil {
		// Log an error message if the migration fails.
		log.Println("Error migrating database:", err)
		return err
	}
	// Return nil if the migration is successful.
	return nil
}

// GetTasks retrieves all tasks from the database.
func GetTasks() ([]Task, error) {
	var tasks []Task
	// Attempt to find all tasks in the database.
	if err := DB.Find(&tasks).Error; err != nil {
		// Log an error message if the retrieval fails.
		log.Println("Error retrieving tasks:", err)
		return nil, err
	}
	// Return the list of tasks if the retrieval is successful.
	return tasks, nil
}

// GetTaskByID retrieves a task from the database by its ID.
func GetTaskByID(id string) (Task, error) {
	var task Task
	// Attempt to find the task in the database by its ID.
	if err := DB.First(&task, "id = ?", id).Error; err != nil {
		// Log an error message if the retrieval fails.
		log.Println("Error retrieving task by ID:", err)
		return task, err
	}
	// Return the task if the retrieval is successful.
	return task, nil
}

// CreateTask creates a new task in the database.
func CreateTask(task Task) error {
	// Set the CreatedAt and UpdatedAt fields to the current time.
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	// Attempt to create the new task in the database.
	if err := DB.Create(&task).Error; err != nil {
		// Log an error message if the creation fails.
		log.Println("Error creating task:", err)
		return err
	}
	// Return nil if the task creation is successful.
	return nil
}

// UpdateTask updates an existing task in the database.
func UpdateTask(task Task) error {
	// Attempt to save the updated task in the database.
	if err := DB.Save(&task).Error; err != nil {
		// Log an error message if the update fails.
		log.Println("Error updating task:", err)
		return err
	}
	// Return nil if the task update is successful.
	return nil
}

// DeleteTask deletes a task from the database by its ID.
func DeleteTask(id string) error {
	// Attempt to delete the task from the database by its ID.
	if err := DB.Delete(&Task{}, "id = ?", id).Error; err != nil {
		// Log an error message if the deletion fails.
		log.Println("Error deleting task:", err)
		return err
	}
	// Return nil if the task deletion is successful.
	return nil
}

// UpdateTaskStatus updates the status of a task in the database by its ID.
func UpdateTaskStatus(id string, status string) error {
	var task Task
	// Attempt to find the task in the database by its ID.
	if err := DB.First(&task, "id = ?", id).Error; err != nil {
		// Log an error message if the retrieval fails.
		log.Println("Error retrieving task by ID:", err)
		return err
	}
	// Update the status of the task.
	task.Status = status
	task.UpdatedAt = time.Now()

	// Attempt to save the updated task in the database.
	if err := DB.Save(&task).Error; err != nil {
		// Log an error message if the update fails.
		log.Println("Error updating task status:", err)
		return err
	}
	// Return nil if the task status update is successful.
	return nil
}
