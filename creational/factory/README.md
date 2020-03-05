### Factory

Factory design pattern is a design pattern where we can provide a factory to create an object that object can be part of some class and the behaviour can be flexible based on the objects and the behaviour it carries.

#### Structure
1. ObjectFactory (concrete class)
    - has method: createObject(param)
2. ObjectInterface (interface)
    - has method: behaviour of the object
3. Objects (concrete class)
    - implements methods in ObjectInterface