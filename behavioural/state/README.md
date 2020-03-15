### State

State design pattern is a design pattern where the state can change the behaviour of an object. This is like finite state machine.

#### Structure
1. State (interface)
    - Has methods regarding to changing state of an object
2. State implementation (concrete class)
    - Implements state interface
    - Has attribute: object
3. Object (concrete class)
    - Has attribute: State