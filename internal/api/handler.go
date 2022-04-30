package api

import (
	"io"
	"log"
	"os"

	"github.com/erauer/wd-launcher/internal/ipc"
	"github.com/erauer/wd-launcher/internal/warp"
)

type Handler struct {
	parser    ipc.Parser
	projects  *warp.Projects
	logger    *log.Logger
	responder *ipc.Responder
}

func NewHandler(parser ipc.Parser, responder *ipc.Responder, projects *warp.Projects, logger *log.Logger) *Handler {
	return &Handler{parser: parser, projects: projects, responder: responder, logger: logger}
}

func (h *Handler) Process(message string) error {
	parsed, err := h.parser.ParseRequest(message)
	if err != nil {
		return err
	}

	switch command := parsed.(type) {
	case ipc.Exit:
		os.Exit(0)
	case ipc.Interrupt:
		h.logger.Println("ignoring interupt")
	case ipc.Search:
		if err := h.search(command.Name); err != nil {
			return err
		}
	case ipc.Activate:
		if err := h.activate(command.Index); err != nil {
			return err
		}
	default:
		h.logger.Printf("Unhandled: %+v\n", message)
	}

	return nil
}

func (h *Handler) search(name string) error {
	projects := h.projects.Search(name, h.logger)

	for _, project := range projects {
		if err := h.responder.Append(project.ID, project.Name, project.Path); err != nil {
			return err
		}
	}

	if err := h.responder.Finished(); err != nil {
		return err
	}

	return nil
}

func (h *Handler) activate(index int) error {
	if err := h.projects.Activate(index, h.logger); err != nil {
		return err
	}

	multi := io.MultiWriter(os.Stdout, h.logger.Writer())
	responder := ipc.NewResponder(multi)

	responder.Close()

	return nil
}
