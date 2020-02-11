package collection

import "GolangDesignPattern/behavioural/iterator/iterator"

type Collection interface {
	CreateIterator() iterator.Iterator
}
