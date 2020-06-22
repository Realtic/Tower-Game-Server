package datastore

type genericStore interface {
	Read()
	Write()
}
