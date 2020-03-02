### Proxy

Proxy design pattern acts like a substitute for 3rd party library / services. Any access to a particular services must go through Proxy and then if allowed, then the proxy can redirect the request to the real library / services.

Benefit of proxy:
1. Filtering request
2. Rate limiting
3. Caching content

#### Structure
1. service (interface)
2. proxy (concrete class)
    - implementation of service
3. real service (concrete class)
    - implementation of service