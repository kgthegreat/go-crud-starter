package controllers

import (
	"net/http"

	"strconv"

	"github.com/go-chi/chi"
	"github.com/kgthegreat/meeteffective/app"
	"github.com/kgthegreat/meeteffective/models"
	"github.com/kgthegreat/meeteffective/repositories"
)

type TopicController struct {
	*app.App
	repositories.TopicRepository
}

func NewTopicController(app *app.App, topicRepository repositories.TopicRepository) *TopicController {
	return &TopicController{app, topicRepository}
}

func (topicController *TopicController) GetAll(w http.ResponseWriter, r *http.Request) {
	topics, err := topicController.TopicRepository.GetAll()
	if err != nil {

		NewAppError(&AppError{false, "Could Not Get Topics", http.StatusBadRequest, err}, w)
	}
	renderTemplate(w, "topics", &ResponseData{Topics: topics})
}

func (topicController *TopicController) GetById(w http.ResponseWriter, r *http.Request) {

	topicIdStr := chi.URLParam(r, "topicId")
	topicId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		NewAppError(&AppError{false, "Invalid request", http.StatusBadRequest, err}, w)
		return
	}
	topic, err := topicController.TopicRepository.FindById(topicId)
	if err != nil {
		NewAppError(&AppError{false, "Could not find topic", http.StatusNotFound, err}, w)
		return
	}

	topics, err := topicController.TopicRepository.GetAll()
	renderTemplate(w, "topics_show",
		struct {
			Topics []*models.Topic
			Topic  *models.Topic
		}{topics, topic})

}

func (topicController *TopicController) Create(w http.ResponseWriter, r *http.Request) {

	title := retrieveFormData(r, w, "topic")

	if len(title) < 1 {
		NewAppError(&AppError{false, "Title is too short", http.StatusBadRequest, nil}, w)
		return
	}

	topic := &models.Topic{
		Title: title,
	}

	err := topicController.TopicRepository.Create(topic)
	if err != nil {
		NewAppError(&AppError{false, "Could not create topic", http.StatusBadRequest, err}, w)
		return
	}

	http.Redirect(w, r, "/", 301)

}

func (topicController *TopicController) DeleteById(w http.ResponseWriter, r *http.Request) {

	topicIdStr := chi.URLParam(r, "topicId")
	topicId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		NewAppError(&AppError{false, "Invalid request", http.StatusBadRequest, err}, w)
		return
	}
	err = topicController.TopicRepository.DeleteById(topicId)
	if err != nil {
		NewAppError(&AppError{false, "Could not delete topic", http.StatusNotFound, err}, w)
		return
	}

	http.Redirect(w, r, "/", 301)

}
