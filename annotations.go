package entris

import "encoding/json"

const (
	skipAnootationName      = "entris-skip"
	SkipCreate         uint = iota
	SkipUpdate
	SkipType
	SkipQuery
	SkipField
	SkipEdge
	SkipAll
)

func (a *annotation) Name() string {
	return a.name
}

func (a *annotation) decode(v interface{}) error {
	buffer, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return json.Unmarshal(buffer, a)
}

func Skip(skips ...uint) *annotation {
	return &annotation{
		name:    skipAnootationName,
		Content: skips,
	}
}
