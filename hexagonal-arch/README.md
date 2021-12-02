```bash

── cmd
├── pkg
└── internal
    ├── core
    │   ├── domain
    │   │   ├── game.go
    │   │   └── board.go
    │   ├── ports
    │   │   ├── repositories.go
    │   │   └── services.go
    │   └── services
    │       └── games
    │           └── service.go
    ├── handlers
    └── repositories

```

* All core components(services, domain, and ports) are placed in `./internal/core` directory
* `Domain models`(structs involved in business logic) - structs define each entity that is part of the domain and can be used across the application
* `Ports` - ports will be interfaces definition used to communicate with actors