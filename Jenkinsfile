pipeline {
    agent any

    environment {
        APP_NAME = "do-host-network-backend"
        PORT = "8000"
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/sagar-rathod-devops/do-host-network-backend.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -o app main.go'
            }
        }

        stage('Run App') {
            steps {
                // Kill any app already running on port 8000, then run new one
                sh '''
                fuser -k ${PORT}/tcp || true
                nohup ./app > app.log 2>&1 &
                echo "App is running on port ${PORT}"
                '''
            }
        }
    }

    post {
        success {
            echo "✅ App deployed on EC2 at port ${PORT}"
        }
        failure {
            echo "❌ Pipeline failed"
        }
    }
}
