### Adapter Design Pattern

Adapter design pattern is a way to provide some kind of conversion works due to incompatibility issue of 2 parties.

For example, a client wants to use library that receives JSON. But the client can only provide data in form of XML. So, the client will call the adapter to process the XML input into JSON then calling the library.

#### Structure
1. Client Interface (interface)
    - has method: execute
2. Adapter classes (concrete class)
    - implements methods of client interface
    - has attribute: the other service that the client wants to call
3. Other services (concrete class / Interface)