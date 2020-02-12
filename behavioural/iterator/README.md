### Iterator

Iterator design pattern is useful for hiding underlying information under a certain collection. Iterator should provide user to be able to traverse through the collection.

#### Structure
1. Iterator (interface)
    - has method: hasNext(): bool, and getNext(): Iterator
2. IteratorImpl (concrete class)
    - has attribute: collection and iterator state (e.g. index)
    - has method: implementation of hasNext and getNext
3. Collection (interface)
    - has method: createIterator(): Iterator
4. CollectionImpl (concrete class)
    - has attribute: collection itself
    - has method: implementation of createIterator