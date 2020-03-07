### Golang Design Pattern

This repository contains list of design pattern implemented in Golang
Reference:
1. https://refactoring.guru/design-patterns/
2. https://golangbyexample.com/all-design-patterns-golang/

#### Key Points
#### Creational Design Pattern
1. Abstract Factory
    - Use this when we want to handle with a lot kind of objects but also we want to omit dependency to concrete class when trying to create the objects
    - Benefit:
        1. We can be sure that objects created by the factory will be compatible with each other
        2. We can avoid tight coupling between client and concrete class
        3. Open Closed Principle: we can easily add objects without changing the existing code
        4. Single Responsibility Principle: creation of objects only handled in one class
2. Factory
    - Use this when:
        1. We don't know exactly which object we are going to be depend on:
            - Factory design pattern separates construction of objects and the code that uses the object, so any addition of objects can be done easily
        2. We provide a library: we wanted the client to be able to put their own object plugin with they provide their own behaviour of the object
        3. We can reuse objects
            - It's better to use factory's createObjects() method instead of constructor, since constructor should always return new object.
            - for example if we already created database connections and we want to reuse the connections
    - Benefit:
        1. Avoid tight coupling between concrete object and business logic code
        2. Single Responsibility Principle
            - We separate objects creation into separate place, so that the object creator only have 1 reason to change
        3. Open Closed Principle
            - We can easily add objects with its behaviour without changing the business logic code
3. Builder
    - Use this when:
        1. We need flexibility in creating objects
            - There are many parameters for creating an object and sometimes not all attribute is used.
        2. We want to create multiple objects with the similar steps of creation.
            - Builder interface provides steps of creating object, and director class guides the process of building it
        3. Constructing Composite object (we can create recursive objects)
    - Benefit:
        1. Reuse builder interface to build several objects with similar steps
        2. Build recursive objects
        3. Single Responsibility Principle: we can isolate building objects and business logic
4. Prototype
    - Use this when:
        1. We need to have a copy for an object, and copying a complex object will be such a hassle
        2. We need a copy for the object but it only exposes method for copying it through an interface. And also, the object doesn't expose all the attributes
    - Benefit:
        1. We can clone objects without depending to the concrete object class (copying through an interface)
        2. We can copy complex objects easily 