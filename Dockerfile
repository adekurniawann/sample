@Library('jenkins-shared-library') _
import com.shared.utils.PodTemplates

agentTemplates = new PodTemplates()

agentTemplates.generalTemplate {
  node(POD_LABEL) {

    checkout scm 
      stage ('Checkout') {
        checkoutCode()
      }
      stage ('Docker Build'){
        dockerBuild()
      }  
      stage ('Docker Push'){
        dockerPush()
      }
      stage ('Helm Deploy'){
        helmDeploy()
      }
  }
}
