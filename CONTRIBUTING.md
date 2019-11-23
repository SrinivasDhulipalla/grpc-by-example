# Contributing

Project is open to contributions.
Even after working with gRPC for several years, I often find myself needing a reference point.
This repository is meant to serve as that canonical reference point.

For the sake of simplicity, I'm trying to focus in on the core gRPC libraries.
I've tried to exclude the large variety of code generators from this project as each can vary greatly.
Instead, I tried to focus on the key concepts such as creating a client or server, setting up mutual TLS, etc.

## Repository Structure

The repository is structured as follows.
Should you chose to contribute another language, then please follow the template that has been laid out.

```
grpc-by-example/
│
├── {{ language }}
│   ├── ...
│   └── {{ example }}
│       └── {{ main/index file }}
├── go
│   ├── go.mod
│   ├── go.sum
│   ├── plaintext-server
│   │   └── main.go
│   └── tls-server
│       └── main.go
└── nodejs
    ├── package-lock.json
    ├── package.json
    ├── plaintext-server
    │   └── index.js
    └── tls-server
        └── index.js
```

## Committing

Generally best practice to have a ticket associated with the pull request.
Not a requirement, but can help me be prepared to accept your pull request if I know it's coming.
