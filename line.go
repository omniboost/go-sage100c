package sage100c

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ianlopshire/go-fixedwidth"
)

type Lines []Line

type Line struct {
	*TypeI
	*TypeM
	*TypeC
}

func (l *Line) UnmarshalFixedWidth(data []byte) error {
	r := []rune(string(data))
	if r[0] == 'C' {
		l2 := TypeC{}
		err := fixedwidth.Unmarshal(data, &l2)
		if err != nil {
			return err
		}

		l.TypeC = &l2
		return nil
	} else if r[0] == 'M' {
		l2 := TypeM{}
		err := fixedwidth.Unmarshal(data, &l2)
		if err != nil {
			return err
		}

		l.TypeM = &l2
		return nil
	} else if r[0] == 'I' {
		l2 := TypeI{}
		err := fixedwidth.Unmarshal(data, &l2)
		if err != nil {
			return err
		}

		l.TypeI = &l2
		return nil
	}

	return errors.New(fmt.Sprintf("Unknown type: %s", string(r[0])))
}

func (l Line) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	if l.TypeI != nil {
		return fixedwidth.Marshal(l.TypeI)
	}
	if l.TypeM != nil {
		return fixedwidth.Marshal(l.TypeM)
	}
	if l.TypeC != nil {
		return fixedwidth.Marshal(l.TypeC)
	}

	return []byte{}, nil
}

func (l Line) Type() string {
	if l.TypeI != nil {
		return l.TypeI.Type
	}
	if l.TypeM != nil {
		return l.TypeM.Type
	}
	if l.TypeC != nil {
		return l.TypeC.Type
	}

	return ""
}

func (l *Line) UnmarshalJSON(data []byte) error {
	s := struct {
		Type string
	}{}

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	switch s.Type {
	case "I":
		l.TypeI = &TypeI{}
		return json.Unmarshal(data, l.TypeI)
	case "M":
		l.TypeM = &TypeM{}
		return json.Unmarshal(data, l.TypeM)
	case "C":
		l.TypeC = &TypeC{}
		return json.Unmarshal(data, l.TypeC)
	}

	return errors.New(fmt.Sprintf("Unknow line type %s", s.Type))
}

func (l Line) MarshalJSON() ([]byte, error) {
	if l.TypeI != nil {
		return json.Marshal(l.TypeI)
	}
	if l.TypeM != nil {
		return json.Marshal(l.TypeM)
	}
	if l.TypeC != nil {
		return json.Marshal(l.TypeC)
	}

	return []byte{}, nil
}
