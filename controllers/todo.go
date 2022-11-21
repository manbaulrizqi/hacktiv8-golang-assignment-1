package controllers

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"latihan1/datas"
	"latihan1/params/request"
	"latihan1/params/views"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// GetTodos godoc
// @Summary Get All TODOS
// @Schemes
// @Description get all todos
// @Tags TODOS
// @Accept json
// @Produce json
// @Success 200 {object} views.GetTodosSuccessSwag
// @Router /todos [get]
func GetAll(c *gin.Context) {
	dsn := `sqlserver://sa:Pass123*@localhost/SQLSERVER14?database=GolangTraining`
	//dsn := `server=localhost\SQLSERVER14;user id=sa;password=Pass123*;database=GolangTraining;encrypt=disable`
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	todoRepository := datas.NewToDoRepository(db)
	dbResults := todoRepository.GetAll()
	var results []request.CreateTodo
	if dbResults != nil {
		for _, todo := range *dbResults {
			results = append(results, request.CreateTodo{
				Title:       todo.Title,
				Description: todo.Description,
			})
		}
	}

	resp := views.GeneralSuccessPayload{
		Status:  http.StatusCreated,
		Message: "GETALL TODO SUCCESS",
		Payload: results,
	}

	c.JSON(resp.Status, resp)

}

// @BasePath /api/v1
// CreateTodo godoc
// @Summary Create TODO
// @Schemes
// @Description create todo
// @Tags TODOS
// @Accept json
// @Produce json
// @Param request body request.CreateTodo  true  "Request Body"
// @Success 200 {object} views.CreateTodoSuccessSwag
// @Failure      400  {object}  views.CreateTodoFailureSwag
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	var req request.CreateTodo

	err := c.ShouldBindJSON(&req)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	var payload = views.CreateTodoPayload{
		ID:          1,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}

	dsn := `sqlserver://sa:Pass123*@localhost/SQLSERVER14?database=GolangTraining`
	//dsn := `server=localhost\SQLSERVER14;user id=sa;password=Pass123*;database=GolangTraining;encrypt=disable`
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	todoRepository := datas.NewToDoRepository(db)
	todoRepository.Create(&datas.Todo{
		Title:       req.Title,
		Description: req.Description,
	})

	resp := views.GeneralSuccessPayload{
		Status:  http.StatusCreated,
		Message: "CREATE TODO SUCCESS",
		Payload: payload,
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// GetTodos godoc
// @Summary Get All TODOS
// @Schemes
// @Description get all todos
// @Param id path int true "Todo ID"
// @Tags TODOS
// @Accept json
// @Produce json
// @Success 200 {object} views.GetTodosSuccessSwag
// @Router /todos/{id} [get]
func GetByID(c *gin.Context) {

	dataId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	dsn := `sqlserver://sa:Pass123*@localhost/SQLSERVER14?database=GolangTraining`
	//dsn := `server=localhost\SQLSERVER14;user id=sa;password=Pass123*;database=GolangTraining;encrypt=disable`
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	todoRepository := datas.NewToDoRepository(db)
	dbResults := todoRepository.GetById(dataId)
	var results []request.CreateTodo
	if dbResults != nil {
		for _, todo := range *dbResults {
			results = append(results, request.CreateTodo{
				Title:       todo.Title,
				Description: todo.Description,
			})
		}
	}

	resp := views.GeneralSuccessPayload{
		Status:  http.StatusCreated,
		Message: "GET TODO SUCCESS",
		Payload: results,
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// UpdateTodo godoc
// @Summary Update TODO
// @Schemes
// @Description update todo
// @Tags TODOS
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param request body request.CreateTodo  true  "Request Body"
// @Success 200 {object} views.UpdateTodoSuccessSwag
// @Failure      400  {object}  views.UpdateTodoFailureSwag
// @Router /todos/{id} [put]
func UpdateByID(c *gin.Context) {
	var req request.CreateTodo

	err := c.ShouldBindJSON(&req)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	dataId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	dsn := `sqlserver://sa:Pass123*@localhost/SQLSERVER14?database=GolangTraining`
	//dsn := `server=localhost\SQLSERVER14;user id=sa;password=Pass123*;database=GolangTraining;encrypt=disable`
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	todoRepository := datas.NewToDoRepository(db)
	data := &datas.Todo{
		Title:       req.Title,
		Description: req.Description,
	}
	data.ID = uint(dataId)
	todoRepository.Update(data)

	resp := views.GeneralSuccessPayload{
		Status:  http.StatusCreated,
		Message: "UPDATE TODO SUCCESS",
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// UpdateTodo godoc
// @Summary Delete TODO
// @Schemes
// @Description delete todo
// @Tags TODOS
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} views.UpdateTodoSuccessSwag
// @Failure      400  {object}  views.UpdateTodoFailureSwag
// @Router /todos/{id} [delete]
func DeleteByID(c *gin.Context) {
	dataId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	dsn := `sqlserver://sa:Pass123*@localhost/SQLSERVER14?database=GolangTraining`
	//dsn := `server=localhost\SQLSERVER14;user id=sa;password=Pass123*;database=GolangTraining;encrypt=disable`
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	todoRepository := datas.NewToDoRepository(db)
	todoRepository.Delete(dataId)

	resp := views.GeneralSuccessPayload{
		Status:  http.StatusCreated,
		Message: "DELETE TODO SUCCESS",
	}

	c.JSON(resp.Status, resp)
}
