### Bridge

Bridge design pattern is a design when there is abstraction and implementation. The bridge consist of abstraction and implementation interface.
Abstraction here is a concrete class, meanwhile Implementation here is an interface.

Abstraction: the high level class that will only delegate work into the implementation (platform)
Real application: Abstraction: GUI, implementation: API

Notice that the term abstraction and implementation here is different with the term we usually use (abstraction = interface, implemnetation = concrete class)

#### Structure

1. Bridge, consist of:
    - Abstraction: concrete class that will only delegate works to the implementation
        - have attribute: the implementation interface
        - have several methods to call the interface
    - Implementation: interface that will be implemented by the other classes
2. Implementation's subclass