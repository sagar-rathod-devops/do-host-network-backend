pipeline {
    agent any

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
                sh 'go mod tidy'
                sh 'go build -o main .'
            }
        }

        stage('Docker Build') {
            steps {
                script {
                    docker.build("${IMAGE_NAME}:${IMAGE_TAG}")
                }
            }
        }

        stage('Docker Run') {
            steps {
                sh 'docker run -d -p 8080:8000 --name backend-container do-host-network-backend:latest'
            }
        }
    }

    post {
        always {
            echo 'Pipeline completed.'
        }
    }
}
