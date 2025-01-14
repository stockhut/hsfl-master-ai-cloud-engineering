# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## loadbalancer-0.2.0 - 2024-01-17
#### Features
- **(load-balancer)** Load config from yaml file - (56a5077) - TuftedDeer

- - -

## loadbalancer-0.1.0 - 2023-12-27
#### Bug Fixes
- **(load-balancer)** fix race condition when passing healthyReplicas to balancing strategy - (fcd94ca) - TuftedDeer
- **(load-balancer)** Implement basic load balancer with round robin and least connections strategies - (8c0342e) - TuftedDeer
- **(load-balancer/round-robin)** fix race condition when updating index - (122bbd7) - TuftedDeer
#### Continuous Integration
- **(load-balancer)** Add Dockerfile - (ca3b68b) - alex
#### Documentation
- **(load-balancer)** add some docs about thread safety - (aaaf2f4) - TuftedDeer
#### Features
- **(load-balancer)** implement ip hash strategy - (0328cc0) - TuftedDeer
- **(load-balancer)** initialize replica list with running container instances - (27a3eea) - TuftedDeer
- **(load-balancer)** turn LeastConnections.m into sync.RWMutex (was sync.Mutex) - (ebd6b8c) - TuftedDeer
- **(load-balancer)** make healthyLock sync.RWMutex (was sync.Mutex) - (5a01d23) - TuftedDeer
- **(load-balancer)** WIP: add container orchestration capabilities to load balancer - (de6d1bb) - TuftedDeer
- **(load-balancer)** Mark host as unhealthy when request forwarding fails - (be4e111) - TuftedDeer
- **(load-balancer)** Implement healthchecks - (b226931) - TuftedDeer
- **(orchestration)** move orchestration code from load-balancer to own project - (9b51c80) - TuftedDeer
#### Miscellaneous Chores
- go mod tidy - (f45fa76) - TuftedDeer
#### Tests
- **(load-balancer/least-connections)** pass correct mutex type - (526d5e0) - TuftedDeer

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).