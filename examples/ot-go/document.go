package ot

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Document struct {
	Content    string
	Operations [][]Applier
}

func (doc *Document) Recieve(message Message) ([]Applier, error) {
	if message.Revision < 0 || message.Revision >= len(doc.Operations) {
		return nil, errors.New("revision not in history")
	}
	concurentOperations := doc.Operations[message.Revision:]
	operation := message.Operation
	var err error
	for _, authorityOp := range concurentOperations {
		operation, _, err = Transform(operation, authorityOp)
		if err != nil {
			return nil, err
		}
	}
	nextDoc, err := Apply(doc.Content, operation...)
	if err != nil {
		return nil, err
	}
	doc.Content = nextDoc
	doc.Operations = append(doc.Operations, operation)
	return operation, nil
}

type Message struct {
	Revision  int       `json:"revision"`
	Operation []Applier `json:"operation"`
}

func (message *Message) UnmarshalJSON(data []byte) error {
	var obj struct {
		Revision   int           `json:"revision"`
		Opperation []interface{} `json:"operation"`
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	message.Revision = obj.Revision
	for _, op := range obj.Opperation {
		switch o := op.(type) {
		case string:
			message.Operation = append(message.Operation, Insert(o))
		case float64:
			if o < 0 {
				message.Operation = append(message.Operation, Delete(int(o)))
			} else {
				message.Operation = append(message.Operation, Retain(int(o)))
			}
		default:
			return fmt.Errorf("unknown op type %v %t", o, o)
		}
	}
	return nil
}
