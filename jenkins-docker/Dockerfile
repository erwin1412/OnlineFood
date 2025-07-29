FROM jenkins/jenkins:lts

USER root

# Install Docker CLI & Compose
RUN apt-get update && \
    apt-get install -y docker.io docker-compose curl apt-transport-https ca-certificates gnupg

# Install Google Cloud SDK
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - && \
    apt-get update && apt-get install -y google-cloud-sdk kubectl

RUN usermod -aG docker jenkins

USER jenkins
