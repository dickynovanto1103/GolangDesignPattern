### Flyweight

Flyweight design pattern is a pattern of reusing something that is similar across a large number of object to save memory usage.

Flyweight pattern has 2 states:
1. Extrinsic state: the states of this object can be changed and differs across objects 
2. Intrinsic state: this object is the one that will be used across objects and it is immutable

#### Structure
1. Flyweight (interface)
2. Flyweight implementations (concrete class)
3. Flyweight Factory (concrete class): we can create a map inside the factory struct and it acts like a cache