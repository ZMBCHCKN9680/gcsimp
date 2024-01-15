package arlecchino

import (
	"github.com/genshinsim/gcsim/pkg/core/glog"
)

func (c *char) a1() {
	if c.Base.Ascension < 1 {
		return
	}
	c.Core.Log.NewEvent("a1 initialised. does nothing", glog.LogCharacterEvent, c.Index)
}

func (c *char) a4() {
	if c.Base.Ascension < 4 {
		return
	}
	c.Core.Log.NewEvent("a4 initialised. does nothing", glog.LogCharacterEvent, c.Index)
}
