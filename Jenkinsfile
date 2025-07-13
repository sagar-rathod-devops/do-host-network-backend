pipeline {
    agent any

    stages {
        stage('Go Version') {
            steps {
                sh 'go version'
            }
        }
    }

    environment {
        IMAGE_NAME = 'do-host-network-backend'
        IMAGE_TAG = 'latest'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout([$class: 'GitSCM',
                    branches: [[name: '*/main']],
                    userRemoteConfigs: [[url: 'https://github.com/sagar-rathod-devops/do-host-network-backend.git']]
                ])
            }
        }

        stage('Build Go App') {
            steps {
                sh 'go version'
                sh 'go mod tidy'
                sh 'go build -o main .'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .'
            }
        }

        stage('Docker Run') {
            steps {
                sh '''
                    docker stop backend-container || true
                    docker rm backend-container || true
                    docker run -d -p 8080:8000 --name backend-container ${IMAGE_NAME}:${IMAGE_TAG}
                '''
            }
        }
    }

    post {
        always {
            echo '✅ Pipeline completed!'
        }
    }
}
