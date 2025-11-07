package node

import (
	"fmt"
	"log"
	"os"
)

type NodeLogger struct {
	*log.Logger
	nodeId   string
	getClock func() int64
	file     *os.File
}

func NewNodeLogger(nodeId string, getClock func() int64, filePath string) (*NodeLogger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	nl := &NodeLogger{
		nodeId:   nodeId,
		getClock: getClock,
		file:     file,
	}

	nl.Logger = log.New(file, fmt.Sprintf("[Node: %s] [Clock: %d] ", nl.nodeId, getClock()), log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	return nl, nil
}

// // Println prepends the dynamic clock/node prefix and delegates to the embedded logger
// func (l *NodeLogger) Println(v ...interface{}) {
// 	prefix := fmt.Sprintf("[Clock: %d] [Node: %s] ", l.getClock(), l.nodeId)
// 	l.Logger.Println(prefix + fmt.Sprint(v...))
// }

// // Printf prepends the dynamic clock/node prefix and delegates to the embedded logger
// func (l *NodeLogger) Printf(format string, v ...interface{}) {
// 	prefix := fmt.Sprintf("[Clock: %d] [Node: %s] ", l.getClock(), l.nodeId)
// 	l.Logger.Printf(prefix+format, v...)
// }

// // Fatalf behaves like Printf but exits after closing file
// func (l *NodeLogger) Fatalf(format string, v ...interface{}) {
// 	prefix := fmt.Sprintf("[Clock: %d] [Node: %s] ", l.getClock(), l.nodeId)
// 	l.Logger.Printf(prefix+format, v...)
// 	l.file.Close()
// 	os.Exit(1)
// }

func (l *NodeLogger) Close() error {
	// Flush any buffered data to disk before closing
	if err := l.file.Sync(); err != nil {
		return err
	}
	return l.file.Close()
}
