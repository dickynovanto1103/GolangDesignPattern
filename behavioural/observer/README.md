### Observer

Observer pattern is like subscription to a topic and we can provide notification to all subscribers (observers) once there is an event.

#### Structure
1. Subject (Interface)
    - has method: AddObserver, RemoveObserver, and NotifyObservers
2. Subject concrete (Concrete class)
    - implements the subject's methods
3. Observer (Interface)
    - has method: Update (giving info to the observer)
4. Observer concrete (Concrete class)
    - implements the observer's methods