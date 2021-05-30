pipeline {
    agent any

    tools {
        go 'go'
    }

    stages {
        
        stage('Build') {
            steps {
                echo 'Building GO project...'
                sh 'go build .'
            }
        }
        
        stage('Build-Docker-Image') {
            steps {
                echo 'Building the docker image for build ${BUILD_ID}'
                sh 'docker build -t ghcr.io/16it043/hello-world:${BUILD_NUMBER} .'
                sh 'docker tag ghcr.io/16it043/hello-world:${BUILD_NUMBER} ghcr.io/16it043/hello-world:latest'
            }
        }
        
        stage('Deploy') {
            steps {
                echo 'Deploying the docker image'
                sh 'docker images'
                sh 'docker container run ghcr.io/16it043/hello-world:${BUILD_NUMBER}'
            }
        }
    }
}