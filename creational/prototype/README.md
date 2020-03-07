### Prototype

Prototype is creational design pattern with the purpose of copying objects more easily

#### Structure

1. Prototype (interface)
    - has method: clone
2. Real objects class (concrete class)
    - implements Prototype interface
3. Prototype wrapper class (concrete class)
    - contains prototype(s)
    - can provides methods like getItem(criteria) and the implementations is if found then clone

Note: the implementation in this example still depends on concrete class, if the app is getting complex, having a prototype wrapper can be beneficial