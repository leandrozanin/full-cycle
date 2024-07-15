pipeline { 
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go1.19'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                sh 'go test ./...'
            }
        }
        stage("functional-test") {
            steps {
                echo 'FUNCTIONAL TEST EXECUTION STARTED'
                sh 'go test ./...'
            }
        }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'go get ./...'
                sh 'go build math.go'
            }
        }
        // stage('deliver') {
        //     agent any
        //     steps {
        //         withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
        //         sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
        //         sh 'docker push shadowshotx/product-go-micro'
        //         }
        //     }
        // }
    }
}