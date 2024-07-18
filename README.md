# Highload Social Messenger Counter

## Description

This project is a part of [Highload Social Messenger](https://github.com/eydeveloper/highload-social-messenger). It extends the messenger
with unread message counter functionality.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed Docker
- You have installed Go

## Setup

To set up this project, follow these steps:

1. Clone the repository:
    ```shell
    git clone https://github.com/eydeveloper/highload-social-messenger-counter.git
    ```

2. Start the Docker services:
    ```shell
    make up
    ```

3. Start the server by running the following command:
    ```shell
    go run cmd/main.go
    ```

## Saga pattern for distributed transactions

This project leverages the Saga pattern to manage distributed transactions within a messaging system. The primary focus is on ensuring data consistency and reliability when updating the unread message counters for recipients in dialogues.

### Problem statement

In a distributed messaging system, each time a message is sent within a dialogue, the unread message counter for the recipient needs to be incremented. This update must be consistent and reliable across multiple microservices or database instances. Traditional transaction management methods may not be feasible or efficient in such a distributed environment, leading to the need for a robust alternative.

### Solution

The Saga pattern addresses the challenges of maintaining data consistency in a distributed system by breaking down a transaction into a series of smaller, manageable sub-transactions. Each sub-transaction is treated as a standalone unit of work, which can be committed or rolled back independently.
