package reader

import (
	"encoding/csv"
	"fmt"
	"github.com/jatis/oms/handler/reader/worker"
	"github.com/jatis/oms/init/service"
	"io"
	"os"
	"sync"
)

type Opts struct {
	Service *service.Service
}

type Handler struct {
	options     *Opts
	listenErrCh chan error
	worker      int
	rows        chan []interface{}
	wg          *sync.WaitGroup
	workerType  string
}

type ReaderHandler interface {
	Run(filepath string)
	Error() <-chan error
}

func New(workerType string, o *Opts) ReaderHandler {
	handler := &Handler{
		options:     o,
		listenErrCh: make(chan error, 1),
		worker:      100,
		rows:        make(chan []interface{}, 0),
		wg:          new(sync.WaitGroup),
		workerType:  workerType,
	}

	go worker.New(&worker.Opts{
		Rows:       handler.rows,
		Wg:         handler.wg,
		Service:    handler.options.Service,
		WorkerType: handler.workerType,
		Error:      handler.listenErrCh,
	}).Register()

	return handler
}

func (h *Handler) Run(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		h.listenErrCh <- fmt.Errorf("failed to open file %s: %v", filepath, err)
		return
	}

	reader := csv.NewReader(f)
	defer f.Close()

	h.readRow(reader)

	h.wg.Wait()
}

func (h *Handler) Error() <-chan error {
	return h.listenErrCh
}

func (h *Handler) readRow(reader *csv.Reader) {
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			h.listenErrCh <- err
			break
		}

		rowOrdered := make([]interface{}, 0)
		for _, each := range row {
			rowOrdered = append(rowOrdered, each)
		}

		h.wg.Add(1)
		h.rows <- rowOrdered
	}

	close(h.rows)
}
