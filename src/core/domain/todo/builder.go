package todo

import (
	"github.com/google/uuid"
	"projeto-docker/src/core/domain/errors"
	"projeto-docker/src/core/messages"
	"projeto-docker/src/utils/validator"
	"time"
)

type builder struct {
	fields        []string
	errorMessages []string
	todo          *todo
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		todo:          &todo{},
	}
}

func (b *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		b.fields = append(b.fields, messages.TodoID)
		b.errorMessages = append(b.errorMessages, messages.TodoIDErrorMessage)
		return b
	}
	b.todo.id = &id
	return b
}

func (b *builder) WithDescription(description string) *builder {
	b.todo.description = description
	return b
}

func (b *builder) WithIsActive(isActive bool) *builder {
	b.todo.isActive = isActive
	return b
}

func (b *builder) WithCreatedAt(createdAt *time.Time) *builder {
	b.todo.createdAt = createdAt
	return b
}

func (b *builder) WithUpdatedAt(updatedAt *time.Time) *builder {
	b.todo.updatedAt = updatedAt
	return b
}

func (b *builder) Build() (Todo, errors.Error) {
	if len(b.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(b.errorMessages, map[string]interface{}{
			"fields": b.fields})
	}
	return b.todo, nil
}
