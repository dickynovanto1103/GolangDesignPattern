### Builder

Builder design pattern is pattern for building complex objects (containing many attributes) and the client wants some flexibility of attribute to be set.

#### Structure
1. BuilderInterface (interface)
    - has method: SetAttributes() --> should be separated per attribute, GetObject()
2. Builder class (concrete class)
    - implementation of builder interface
3. Director (concrete class)
    - has attribute: builder
    - has method: Build(), SetBuilder(builder)
