### Command

Command design pattern can be used to decouple between frontend and backend.

For example, suppose there is button class, and there is many concrete button as a subclass of the button class. If our design is limited to this, this is okay, but let's say a command can be executed not only by using button but we can use shortcut (copy button and ctrl + c), so we might be duplicating copy logic code. So this is not good.

Therefore, we should create another layer, called command as a middleware, to route request to exact receiver.

#### Structure
1. Invoker (concrete class)
    - has attribute: command
    - has method: setCommand(command), and one method for invoking command
2. Command (interface)
    - has method: execute()
3. Command Implementations (concrete class, impl of Command)
    - has attribute: receiver
    - has method: execute(), implementation of the execute method
3. Receiver (interface)
    - has method: methods depends on type of receiver. Should further doing interface segregation if there is supposed to be many receiver such that there is no interface pollution.
