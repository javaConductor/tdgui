package main

import "encoding/json"

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

func (os *ObjectSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(&ObjectSpec{
		Name:        os.Name,
		Type:        "ObjectSpec",
		Fields:      os.Fields,
		Constraints: os.Constraints,
	})
}

func (os *ObjectSpec) UnmarshalJSON(b []byte) error {
	var o ObjectSpec
	if err := json.Unmarshal(b, &o); err != nil {
		return err
	}
	os.Type = "ObjectSpec"
	return nil
}

// DataSetSpec ...
type DataSetSpec struct {
	Name           string       `json:"name"`
	Type           string       `json:"type"`
	ArtifactId     string       `json:"artifactId"`
	ObjectSpecList []ObjectSpec `json:"objectSpecList"`
}

func (dss *DataSetSpec) withType(objectSpec *ObjectSpec) *DataSetSpec {
	newDss := DataSetSpec{Name: dss.Name, Type: dss.Type, ObjectSpecList: append(dss.ObjectSpecList, *objectSpec)}
	return &newDss
}

func (dss *DataSetSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(&DataSetSpec{
		Name:           dss.Name,
		Type:           "DataSetSpec",
		ArtifactId:     dss.ArtifactId,
		ObjectSpecList: dss.ObjectSpecList,
	})
}

func (dss *DataSetSpec) UnmarshalJSON(b []byte) error {
	var d DataSetSpec
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}

	dss.Type = "DataSetSpec"
	return nil
}
