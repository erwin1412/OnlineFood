pipeline {
  agent any

  environment {
    GOOGLE_APPLICATION_CREDENTIALS = credentials('gcp-service-account') // ID dari Jenkins Credentials
    PROJECT_ID = "your-gcp-project-id"
    IMAGE_NAME = "onlinefood"
    IMAGE_TAG = "latest"
    CLUSTER_NAME = "your-gke-cluster"
    CLUSTER_ZONE = "asia-southeast2-a"
  }

  stages {
    stage('Checkout') {
      steps {
        git branch: 'main', url: 'https://github.com/erwin1412/OnlineFood.git'
      }
    }

    stage('Auth GCP') {
      steps {
        sh '''
          echo $GOOGLE_APPLICATION_CREDENTIALS > key.json
          gcloud auth activate-service-account --key-file=key.json
          gcloud config set project $PROJECT_ID
        '''
      }
    }

    stage('Build Image') {
      steps {
        sh '''
          docker build -t gcr.io/$PROJECT_ID/$IMAGE_NAME:$IMAGE_TAG .
        '''
      }
    }

    stage('Push Image') {
      steps {
        sh '''
          gcloud auth configure-docker
          docker push gcr.io/$PROJECT_ID/$IMAGE_NAME:$IMAGE_TAG
        '''
      }
    }

    stage('Deploy to GKE') {
      steps {
        sh '''
          gcloud container clusters get-credentials $CLUSTER_NAME --zone $CLUSTER_ZONE
          kubectl apply -f k8s/
        '''
      }
    }
  }
}
