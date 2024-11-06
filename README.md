### Project Configuration

This project serves as an academic exercise to explore and experiment with various tools, applying best practices and clean principles in software development. While some challenges may arise that don't necessarily align with widely adopted practices in web development, the focus is on learning, testing, and understanding how different technologies and methodologies work together in real-world applications.
- **.env Files**: The `.env` files in the `build` directory are not added to the `.gitignore` file. They are used to provide sample configuration.

- **API Authentication**: There is a script to generate a token for API usage, which requires Bearer token authentication.

- **Directory Structure**:
  - **`build` Directory**: Contains all Dockerfiles to run the application.
    - **Subdirectories**:
      - **`Global`**: Contains a Docker Compose configuration for running the application with databases. Optimal for use in Docker.
      - **`Samples`**: Contains Docker Compose configurations for databases, messaging queues, etc. Optimal for local development. You can copy the configurations from `samples` to `Global`, respecting the networking setup, to adjust the deployment.

- **Docker Configuration**: The Docker configuration files are pushed to the repository to avoid loss of deployment details. This ensures that all deployment-related information is temporarily stored in the repository.

- **Missing Work**:
  - `Resolve the TODOs in the code.`
  - `Define the authentication to taskManager and consumers`
