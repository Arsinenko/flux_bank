# Development Plan for Orchestra Service (orch-go)

This document outlines the development roadmap for the `orch-go` service, which serves as the orchestrator and emulator for the banking system. It manages autonomous agents and interacts with backend services via gRPC.

## 1. Foundation & Architecture
* [ ] **Initialize Main Entry Point**: Implement `cmd/orchestrator/main.go` to bootstrap the application, load config, and handle graceful shutdown.
* [ ] **Dependency Injection**: Set up a DI container (e.g., manual wiring or a library) to manage clients, repositories, and services.
* [ ] **Configuration Management**: Implement loading configuration (env variables/files) for gRPC service endpoints, agent settings, and log levels.

## 2. Core Domain & Infrastructure
* [ ] **Complete Repositories**: Ensure all infrastructure repositories (Account, Transaction, etc.) correctly wrap their respective gRPC clients.
    * *Note*: Existing repositories in `internal/infrastructure/repository` act as adapters to external gRPC services.
* [ ] **Implement Service Layer**: Create domain services in `internal/services` to encapsulate higher-level logic.
    * Example: `AuthService`, `TransactionService` (handling complex flows if necessary).

## 3. Autonomous Agent System
The core purpose of this service is to run agents that emulate user behavior.
* [ ] **Agent Framework Design**:
    * Define `Agent` interface/struct.
    * Define `Behavior` strategy pattern (e.g., `Saver`, `Spender`, `Investor`).
* [ ] **Agent Scheduler/Runner**: Implement a mechanism to spawn and manage the lifecycle of multiple agents.
* [ ] **Behavior Implementation**:
    * Implement specific behaviors (randomized transactions, periodic logins, loan applications).
    * Ensure agents use the **Service Layer** to perform actions, not raw repositories.
* [ ] **Simulation Loop**: Create the main loop that ticks agents or schedules their tasks.

## 4. Analytics Integration (Future)
Planning for the future connection to an Analytics Service.
* [ ] **Event Schema Design**: Define internal structures for tracking significant events (e.g., `AgentAction`, `TransactionCompleted`).
* [ ] **Event Bus / Publisher**: Create an internal abstraction for publishing events.
    * *Initially*: Can just log to stdout/file.
    * *Future*: Will be replaced by a gRPC client or Message Queue producer sending to the Analytics Service.
* [ ] **Instrumentation**: Add hooks in the **Service Layer** or **Agent Behaviors** to emit events when actions are performed.

## 5. Testing & Verification
* [ ] **Unit Tests**: Test agent behaviors and service logic.
* [ ] **Integration Tests**: Verify interactions with the mocked gRPC backend.
