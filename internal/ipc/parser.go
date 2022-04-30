package ipc

import (
	"encoding/json"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseRequest(message string) (interface{}, error) {
	switch message {
	case "\"Exit\"":
		return Exit{}, nil
	case "\"Interrupt\"":
		return Interrupt{}, nil
	}

	var request Request
	if err := json.Unmarshal([]byte(message), &request); err != nil {
		return nil, err
	}

	switch {
	case request.Activate > 0:
		return Activate{Index: request.Activate}, nil
	case request.Context > 0:
		return Context{Index: request.Context}, nil
	case request.Complete > 0:
		return Complete{Index: request.Complete}, nil
	case request.Quit > 0:
		return Quit{Index: request.Quit}, nil
	case len(request.Search) > 0:
		return Search{Name: request.Search}, nil
	}

	return nil, nil
}
