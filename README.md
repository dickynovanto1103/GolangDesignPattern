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
5. Singleton
    - Use this when:
        1. We want an object to be used in many part of the code.
        2. We want to have stricter control over global variable
            - The only one who can create / access the variable is the singleton class itself
    - Disadvantage:
        1. Violate single responsibility principle.
            - The code now depends on singleton class and if the singleton class and the business logic change, we have more than 1 reason to change the class.
        2. Hard to unit test since we depend on the singleton
            - solution: Maybe we can find way to mock it or just don't use this pattern
        3. Requires extra attention to multithreading programming such that the object is guaranteed only created once
        4. The part of code know to much about the other code. Bad design in general 

#### Behavioural Design Pattern
1. Chain of Responsibility
    - Use this when:
        1. We want to have several actions and it's too big to be handled in one handler
            - In this process we can do it step by step, asking each handler whether they are responsible to do it, or the other handler's job
        2. We want to execute several handler in a certain order
        3. We want to dynamically set the handler order
    - Benefit:
        1. We can control the order of the handler
        2. Single Responsibility Principle
            - Previously, the request is handled in a big handler, and if we separate the handler into separate handlers, then each handler will only change if there is only 1 reason to change.
        3. Open Closed Principle
            - We can introduce new handler into the code without breaking the application / changing the business logic.
2. Command
    - We should use this:
        1. We should separate concern of handling request and show something in frontend
            - Showing something is the job of frontend code, and handling request must be processed in backend, and any request must go through 1 router, and it will route to specific interface implementation to handle the request.
    - Benefit:
        1. Single Responsibility Principle
            - We can create each request handler and each request handler only change on 1 reason only
        2. Open Closed Principle
            - We can easily add command implementation without changing main business logic code.