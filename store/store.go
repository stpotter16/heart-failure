package store

import heartfailure "github.com/stpotter16/heart-failure/heart-failure"

type Store interface {
	InsertRecord(record heartfailure.Record) (string, error)
}
