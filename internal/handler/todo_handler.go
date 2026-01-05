package handler

import (
	"net/http"
	"strconv"
	"todolist/internal/domain"
	"todolist/internal/service"
	"todolist/pkg/response" // El helper que pide tu manual

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService *service.TodoService
}

// NewTodoHandler inyecta el servicio
func NewTodoHandler(todoService *service.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	// Pasamos el contexto de la petición original
	todos, err := h.todoService.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Error al obtener todos", err)
		return
	}

	response.Success(c, http.StatusOK, "Todos recuperados", todos)
}

func (h *TodoHandler) Create(c *gin.Context) {
	// Definimos una estructura local para leer el JSON de entrada
	var input struct {
		Content string `json:"content" binding:"required"`
		Status  bool   `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, "Datos de entrada inválidos", err)
		return
	}

	// Convertimos el input a nuestra Entidad de Dominio
	todo := &domain.Todo{
		Content: input.Content,
		Status:  input.Status,
	}

	// Llamamos al servicio
	if err := h.todoService.Create(c.Request.Context(), todo); err != nil {
		response.Error(c, http.StatusInternalServerError, "No se pudo crear el todo", err)
		return
	}

	response.Success(c, http.StatusCreated, "Todo creado con éxito", todo)
}

func (h *TodoHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32) // Convertimos string a uint

	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID inválido", err)
		return
	}

	todo, err := h.todoService.GetById(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "No se encontró el recurso", err)
		return
	}

	response.Success(c, http.StatusOK, "Todo encontrado", todo)
}
