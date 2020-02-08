### Singleton

Singleton relates to creating an object and we can use the object entirely in the application.

An object that's created sometimes should only created once. If we deal with Goroutines, we must take care of concurrency in creating object.

We can do it in 2 ways:
1. Use mutex and do locking. We must be careful here since we do checking first if the project is nil or not, then if it is nil, then we are ready to create by acquiring lock. We then need to check again whether the object is already created after locking phase, if it is not created, then we can safely assume that we can create it and it will be created once.
2. The other thing is we can use sync.Once and we can get rid of the additional nil object checking