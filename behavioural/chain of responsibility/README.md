### Chain of Responsibility

In this design pattern, we can have a chain of responsibility, so that each unit of process is handled by a handler and can be passed into another handler.
Having this design pattern, we can have 2 kinds of scenario:
1. We can create some sort of filter along the way of chain. For example, processing an authentication request from a particular user. We can do some filter along the way, like if the request is still below rate limit, or is the credential correct
2. We can cut off the chain process if it can be handled in the middle of the process. For example:
   1. a client call tech support, first it is handled by robotic answering machine
   2. if the answering machine cannot help, you can continue process to customer service.
   3. if customer services cannot do any favour, your question is passed into technical guy and in this process, your request is handled properly

#### Structure

1. Create an interface that have 2 method, setNext and handleRequest
2. Create classes that implements the interface, i.e. the handler, and each handle have next attribute with type of interface, then the next attr will be set using setNext method