package datas

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID          uint      `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"<-:create"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string
	Description string
}

type ToDoRepository struct {
	db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) *ToDoRepository {
	result := &ToDoRepository{
		db: db,
	}

	return result
}

func (r *ToDoRepository) Create(data *Todo) {
	r.db.Create(data)
}

func (r *ToDoRepository) Update(data *Todo) {
	r.db.Save(data)
}

func (r *ToDoRepository) Delete(id int) {
	r.db.Delete(&Todo{}, id)
}

func (r *ToDoRepository) GetAll() *[]Todo {
	var result []Todo
	r.db.Find(&result)
	return &result
}

func (r *ToDoRepository) GetById(id int) *[]Todo {
	var result []Todo
	r.db.Find(&result, id)
	return &result
}
