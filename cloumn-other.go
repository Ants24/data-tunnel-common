package datatunnelcommon

type UnKnown struct {
	Val string
}

func (t *UnKnown) ID() Type       { return UNKNOWN }
func (t *UnKnown) Name() string   { return "unknown" }
func (t *UnKnown) String() string { return "unknown" }

func (t *UnKnown) Value() string { return t.Val }
