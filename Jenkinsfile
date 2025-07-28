pipeline {
  agent any

  environment {
    PROJECT_ID = 'exalted-arcanum-465715-h7'
    REGION = 'asia-southeast2'
    REPO = 'onlinefood-repo'
  }

  stages {
    stage('Checkout') {
      steps {
        git credentialsId: 'github-pat', url: 'https://github.com/erwin1412/OnlineFood.git'
      }
    }

    stage('Generate') {
      steps {
        sh 'chmod +x generate.sh && ./generate.sh'
      }
    }

    stage('Build Images') {
      steps {
        sh 'docker-compose build'
      }
    }

    stage('Auth GCP') {
      steps {
        withCredentials([file(credentialsId: 'gcp-sa-key', variable: 'GOOGLE_APPLICATION_CREDENTIALS')]) {
          sh 'gcloud auth activate-service-account --key-file=$GOOGLE_APPLICATION_CREDENTIALS'
          sh 'gcloud auth configure-docker $REGION-docker.pkg.dev'
        }
      }
    }

    stage('Push Images') {
      steps {
        sh 'docker-compose push'
      }
    }

    stage('Deploy') {
      steps {
        sh '''
          echo "Contoh deploy ke GKE:"
          echo "gcloud container clusters get-credentials CLUSTER_NAME --zone ZONE --project $PROJECT_ID"
          echo "kubectl apply -f k8s/"
        '''
      }
    }
  }
}
