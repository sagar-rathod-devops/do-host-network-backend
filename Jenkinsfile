pipeline {
    agent any

    environment {
        APP_NAME = 'do-host-network-backend'
        DOCKER_IMAGE = "sagar-rathod/${APP_NAME}:latest"
    }

    stages {

        stage('Install Go') {
    steps {
        sh '''
        if ! command -v go &> /dev/null
        then
          echo "Installing Go locally..."
          wget https://golang.org/dl/go1.20.12.linux-amd64.tar.gz
          rm -rf go
          mkdir go
          tar -C go -xzf go1.20.12.linux-amd64.tar.gz --strip-components=1
          export PATH=$PATH:$(pwd)/go/bin
        fi
        go version
        '''
    }
}


        stage('Checkout') {
    steps {
        git branch: 'main', url: 'https://github.com/sagar-rathod-devops/do-host-network-backend.git'
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
