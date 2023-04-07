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
	SkipSchema
)

func (a *skipAnnotation) Name() string {
	return a.name
}

func (a *skipAnnotation) decode(v interface{}) error {
	buffer, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return json.Unmarshal(buffer, a)
}

func Skip(skips ...uint) *skipAnnotation {
	return &skipAnnotation{
		name:  skipAnootationName,
		Skips: skips,
	}
}
