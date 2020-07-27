package main

// FieldSpec ...
type FieldSpec struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// ObjectSpec ...
type ObjectSpec struct {
	Name        string                           `json:"name"`
	Type        string                           `json:"type"`
	Fields      []FieldSpec                      `json:"fieldSpecList"`
	Constraints map[string]map[string]Constraint `json:"constraints"`
}

// DataSetSpec ...
type DataSetSpec struct {
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	ObjectTypes []ObjectSpec `json:"objectSpecList"`
}

func (ds *DataSetSpec) withType(objectSpec *ObjectSpec) *DataSetSpec {
	dss := DataSetSpec{Name: ds.Name, Type: ds.Type, ObjectTypes: append(ds.ObjectTypes, *objectSpec)}
	return &dss
}
