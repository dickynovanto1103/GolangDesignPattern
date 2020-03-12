### Mediator

Mediator design pattern is a design problem where we use a mediator component as a bridge for communication between 2 components.
This design pattern can solve problem related to tightly coupled components (2 components directly communicate each other) and it makes the component hard to reuse.

#### Structure
1. Mediator interface (interface)
    - has method: Notify. For notifying the other component
2. Mediator class (concrete class)
    - implementation of mediator interface
    - has methods to call the other component
    - has attributes: components
3. Component classes (concrete class)
    - has attribute: Mediator interface