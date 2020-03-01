### Facade Design Pattern

Facade design pattern provides a class as an interface for client to use to get rid of complexity calling another 3rd party library.

In the code provided, let's say we have 2 services, iteminfo, and itemstock, and if the client only cares about ordering an item.
Thus it's better if we provide an interface for client to use, and the facade methods handle initialization of services' objects and calling the interface methods that the service provide.

#### Structure
1. 3rd party services / libraries
2. Facade (concrete class)
    - has methods that will call 3rd party libraries
3. (optional) Additional Facade (concrete class)
    - we can have this once the main facade is becoming to big and we need to separate based on functionality