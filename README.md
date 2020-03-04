### Golang Design Pattern

This repository contains list of design pattern implemented in Golang

#### Key Points
#### Creational Design Pattern
1. Abstract Factory
    - Use this when we want to handle with a lot kind of objects but also we want to omit dependency to concrete class when trying to create the objects
    - Benefit:
        1. We can be sure that objects created by the factory will be compatible with each other
        2. We can avoid tight coupling between client and concrete class
        3. Open Closed Principle: we can easily add objects without changing the existing code
        4. Single Responsibility Principle: creation of objects only handled in one class
 