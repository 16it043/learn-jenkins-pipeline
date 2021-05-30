pipeline {
    agent any

    tools {
        go 'go'
    }

    stages {
        stage('Build') {
            steps {
                sh 'go env'
                echo 'Building...1'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...1'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...1'
            }
        }
    }
}