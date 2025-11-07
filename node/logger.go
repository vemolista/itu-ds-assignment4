package node

import (
	"fmt"
	"log"
	"os"
)

type NodeLogger struct {
	logger   *log.Logger
	nodeId   string
	getClock func() int64
	file     *os.File
}

func NewNodeLogger(nodeId string, getClock func() int64, filePath string) (*NodeLogger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logger := log.New(file, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	return &NodeLogger{
		logger:   logger,
		nodeId:   nodeId,
		getClock: getClock,
		file:     file,
	}, nil
}

func (l *NodeLogger) Println(v ...interface{}) {
	prefix := fmt.Sprintf("[Node: %s] [Clock: %d] ", l.nodeId, l.getClock())
	l.logger.SetPrefix(prefix)
	l.logger.Println(v...)
}

func (l *NodeLogger) Printf(format string, v ...interface{}) {
	prefix := fmt.Sprintf("[Node: %s] [Clock: %d] ", l.nodeId, l.getClock())
	l.logger.SetPrefix(prefix)
	l.logger.Printf(format, v...)
}

func (l *NodeLogger) Close() error {
	// Flush any buffered data to disk before closing
	if err := l.file.Sync(); err != nil {
		return err
	}
	return l.file.Close()
}
