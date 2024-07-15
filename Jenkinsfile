// Run on an agent where we want to use Go
node {
    // Ensure the desired Go version is installed on this agent,
    // using the name defined in the Global Tool Configuration
    def root = tool type: 'go', name: 'go-1.20'

    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
       withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                
                stage('Checkout'){
                    sh 'go version'
                    echo 'Checking'
                    sh 'git clone https://github.com/leandrozanin/full-cycle'
                    sh 'cd ./full-cycle && go mod tidy'
                    sh 'cd ./full-cycle && go test ./...'
                    
                   // sh 'git checkout develop'
                }
                
       }
    }
   
}