package proc

type Mapping []Map

func (m Mapping) Len() int           { return len(m) }
func (m Mapping) Less(i, j int) bool { return m[i].Start < m[j].Start }
func (m Mapping) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
