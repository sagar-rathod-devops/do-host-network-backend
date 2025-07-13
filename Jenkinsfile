pipeline {
    agent any

    environment {
        IMAGE_NAME = 'do-host-network'
        IMAGE_TAG = 'latest'
    }

    stages {
        stage('Checkout') {
            steps {
                branch 'main' git 'https://github.com/sagar-rathod-devops/do-host-network-backend.git'
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
                    dockerImage = docker.build("${IMAGE_NAME}:${IMAGE_TAG}")
                }
            }
        }

        stage('Docker Run') {
            steps {
                sh 'docker run -d -p 8080:8080 --name go-container my-go-app:latest'
            }
        }
    }

    post {
        always {
            echo 'Pipeline execution completed.'
        }
    }
}
