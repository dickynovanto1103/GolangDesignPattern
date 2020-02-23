### Visitor

Visitor design pattern comes with the problem of endless possibility of growing interface, as for one user requests to add a functionality, and another user requests another functionality.
Therefore, in order to keep the interface small, we can apply visitor design pattern, that is to create only 1 method in interface, i.e. Accept(visitor), and then execute visitor.ExecuteAnyMethod().
This way, visitors acts like a plugin to the interface.

#### Structure
1. Visitor (interface)
    - has method: VisitObject(Object)
2. Implementations of Visitor (concrete class)
    - implements the visitor's methods
    - e.g. AreaCalculator and implements each method for each struct of Shape interface
3. Element (interface) --> class that is visited by the visitor
    - has method: Accept(visitor)
4. Implementations of Element (concrete class)
    - implements the element's methods
    - e.g. Circle, Square that will accept visitor, and then will just call visitor.VisitObject(square/circle) and it will execute the function of the struct that is in the parameter
