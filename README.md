### Project Configuration

This project serves as an academic exercise to explore and experiment with various tools, applying best practices and clean principles in software development. While some challenges may arise that don't necessarily align with widely adopted practices in web development, the focus is on learning, testing, and understanding how different technologies and methodologies work together in real-world applications.
- **.env Files**: The `.env` files in the `build` directory are not added to the `.gitignore` file. They are used to provide sample configuration. The reason behind is to provide an example if the repository wants to be tested locally.

- **Directory Structure**:
  - **`build` Directory**: Contains all Dockerfiles to run the application.
    - **Subdirectories**:
      - **`Global`**: Contains a Docker Compose configuration for running the application with databases. Optimal for use in Docker.
      - **`Samples`**: Contains Docker Compose configurations for databases, messaging queues, etc. Optimal for local development. You can copy the configurations from `samples` to `Global`, respecting the networking setup, to adjust the deployment.

- **Docker Configuration**: The Docker configuration files are pushed to the repository to avoid loss of deployment details. This ensures that all deployment-related information is temporarily stored in the repository.

# Work In Progress

- **Enhance user accounts and transactions**:  
  Expanding the functionality to cover more aspects of user accounts and their transactions.

- **Investigate JOINs in the database**:  
  Determine how to incorporate JOINs in the bounded contexts and implement a solution where `GET /users` returns both user details and associated accounts, potentially using read models.

- **Jenkins, CI/CD, and Kubernetes**:  
  Currently working on integrating Jenkins for continuous integration and deployment (CI/CD) and setting up Kubernetes for container orchestration.
    
    - **Testing**:
    Work in testing. Narrowly related with Jenkins and K8s.

- **Request Authentication Management**:  
  Developing a factory to handle different authentication methods, including Basic Authentication and Bearer Token.

- **Extend persistence and bus functionalities**:  
  Expanding the persistence layer and bus functionality to support other databases and messaging systems like MySQL, RabbitMQ, MongoDB, Redis, etc. Some of these systems may require the application of the RPC pattern.

- **Secrets Management**:  
  Implementing a solution to securely store credentials, transitioning from current storage in `.json` files to tools like Vault or AWS Secrets Manager.


