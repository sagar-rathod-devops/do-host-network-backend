pipeline {
    agent any

    environment {
        GO_VERSION = '1.20'
        APP_NAME = 'do-host-network-backend'
        DOCKER_IMAGE = "sagar-rathod/${APP_NAME}:latest"
    }

    tools {
        go "${GO_VERSION}"
    }

    stages {

        stage('Checkout') {
            steps {
                git 'https://github.com/sagar-rathod-devops/do-host-network-backend.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -o main .'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Docker Build') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE}")
                }
            }
        }

        stage('Docker Push') {
            when {
                expression { return env.BRANCH_NAME == 'main' }
            }
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh "echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin"
                    sh "docker push ${DOCKER_IMAGE}"
                }
            }
        }
    }

    post {
        success {
            echo "✅ Jenkins pipeline completed successfully."
        }
        failure {
            echo "❌ Pipeline failed."
        }
    }
}
