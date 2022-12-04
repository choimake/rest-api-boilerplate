package createtodo

import (
	"rest-api-boilerplate/internal/usecase/createtodo"
)

func InputTranslator(req *Request) *createtodo.Input {

	i := &createtodo.Input{
		UserId:      req.UserId,
		Content:     req.Content,
		Media:       req.Media,
		MediaHeader: req.MediaHeader,
	}

	return i
}

func outputTranslator(o *createtodo.Output) Todo {

	return Todo{
		Id:       o.Todo.Id,
		UserId:   o.Todo.UserId,
		Content:  o.Todo.Content,
		MediaUrl: o.Todo.MediaUrl,
	}
}
