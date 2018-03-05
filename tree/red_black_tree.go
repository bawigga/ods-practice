package tree

import (
	log "github.com/Sirupsen/logrus"
)

func init() {
	//log.SetLevel(log.DebugLevel)
	log.SetLevel(log.InfoLevel)
}

type (
	RBTree struct {
		nodes []int
	}
	RBTNode struct {
		color byte
	}
)
