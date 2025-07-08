pipeline {
    agent {
        label 'linux' // Ensure this is a Linux Jenkins agent
    }

    environment {
        AWS_REGION = 'ap-south-1'
        IMAGE_NAME = 'do-host-network-backend'
        ACCOUNT_ID = '248189939111'
        ECR_REPO = "${ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${IMAGE_NAME}"
        EC2_HOST = 'ubuntu@54.235.0.39' // Replace with your EC2 Public IP
        SSH_KEY = 'ec2-ssh-key' // Jenkins credential ID for .pem SSH key
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
                withCredentials([[$class: 'AmazonWebServicesCredentialsBinding', credentialsId: 'aws-ecr-creds']]) {
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
                sshagent (credentials: [SSH_KEY]) {
                    sh """
                        ssh -o StrictHostKeyChecking=no $EC2_HOST << EOF
                            sudo apt update -y
                            sudo apt install -y docker.io awscli
                            sudo systemctl enable docker
                            sudo systemctl start docker
                            aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $ECR_REPO
                            docker stop $IMAGE_NAME || true
                            docker rm $IMAGE_NAME || true
                            docker pull $ECR_REPO:latest
                            docker run -d --name $IMAGE_NAME -p 8000:8000 $ECR_REPO:latest
                        EOF
                    """
                }
            }
        }
    }

    post {
        success {
            echo '✅ Deployment to EC2 on port 8000 completed successfully.'
        }
        failure {
            echo '❌ Pipeline failed. Check logs.'
        }
    }
}
