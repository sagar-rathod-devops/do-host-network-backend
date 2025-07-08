pipeline {
    agent any

    environment {
        AWS_REGION = 'ap-south-1'
        IMAGE_NAME = 'do-host-network-backend'
        ACCOUNT_ID = '248189939111'
        ECR_REPO = "${ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${IMAGE_NAME}"
        EC2_HOST = 'ubuntu@54.235.0.39' // use 'ubuntu' for Ubuntu AMIs
        SSH_KEY = 'ec2-ssh-key' // Jenkins credential ID (.pem key for EC2)
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/sagar-rathod-devops/do-host-network-backend.git'
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
                sh 'docker build -t $IMAGE_NAME .'
            }
        }

        stage('Login to ECR') {
            steps {
                withAWS(credentials: 'aws-ecr-creds', region: "${AWS_REGION}") {
                    sh '''
                        aws ecr get-login-password --region $AWS_REGION | \
                        docker login --username AWS --password-stdin $ECR_REPO
                    '''
                }
            }
        }

        stage('Tag & Push to ECR') {
            steps {
                sh '''
                    docker tag $IMAGE_NAME:latest $ECR_REPO:latest
                    docker push $ECR_REPO:latest
                '''
            }
        }

        stage('Deploy on Ubuntu EC2') {
            steps {
                sshagent (credentials: ['ec2-ssh-key']) {
                    sh """
                    ssh -o StrictHostKeyChecking=no $EC2_HOST << 'EOF'
                        sudo apt update -y
                        sudo apt install -y docker.io awscli
                        sudo usermod -aG docker ubuntu
                        newgrp docker <<EONG
                            aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $ECR_REPO
                            docker pull $ECR_REPO:latest
                            docker stop $IMAGE_NAME || true
                            docker rm $IMAGE_NAME || true
                            docker run -d --name $IMAGE_NAME -p 8000:8000 $ECR_REPO:latest
                        EONG
                    EOF
                    """
                }
            }
        }
    }

    post {
        success {
            echo '✅ Image pushed to ECR and deployed on Ubuntu EC2 (port 8000).'
        }
        failure {
            echo '❌ Pipeline failed.'
        }
    }
}
