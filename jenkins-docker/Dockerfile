# jenkins-docker/Dockerfile

FROM jenkins/jenkins:lts

# Switch to root user to install dependencies
USER root

# Reduce vulnerabilities by using a minimal base image and cleaning up
RUN apt-get update && apt-get install -y --no-install-recommends \
    docker.io \
    docker-compose && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Add Jenkins user to the Docker group
RUN usermod -aG docker jenkins

# Switch back to Jenkins user
USER jenkins

# Ensure the base image is updated

USER root

# Install Docker CLI & Compose
RUN apt-get update && \
    apt-get install -y docker.io docker-compose && \
    usermod -aG docker jenkins

USER jenkins
