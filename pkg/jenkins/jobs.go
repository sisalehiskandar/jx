package jenkins

import (
	"github.com/jenkins-x/jx/pkg/gits"
)

func CreateFolderXml(folderUrl string, name string) string {
	return `<?xml version='1.0' encoding='UTF-8'?>
<com.cloudbees.hudson.plugins.folder.Folder plugin="cloudbees-folder@6.2.1">
  <actions>
    <io.jenkins.blueocean.service.embedded.BlueOceanUrlAction plugin="blueocean-rest-impl@1.3.3">
      <blueOceanUrlObject class="io.jenkins.blueocean.service.embedded.BlueOceanUrlObjectImpl">
        <mappedUrl>blue/organizations/jenkins</mappedUrl>
        <modelObject class="com.cloudbees.hudson.plugins.folder.Folder" reference="../../../.."/>
      </blueOceanUrlObject>
    </io.jenkins.blueocean.service.embedded.BlueOceanUrlAction>
  </actions>
  <description></description>
  <properties>
    <org.jenkinsci.plugins.pipeline.modeldefinition.config.FolderConfig plugin="pipeline-model-definition@1.2.4">
      <dockerLabel></dockerLabel>
      <registry plugin="docker-commons@1.9"/>
    </org.jenkinsci.plugins.pipeline.modeldefinition.config.FolderConfig>
  </properties>
  <folderViews class="com.cloudbees.hudson.plugins.folder.views.DefaultFolderViewHolder">
    <views>
      <hudson.model.AllView>
        <owner class="com.cloudbees.hudson.plugins.folder.Folder" reference="../../../.."/>
        <name>All</name>
        <filterExecutors>false</filterExecutors>
        <filterQueue>false</filterQueue>
        <properties class="hudson.model.View$PropertyList"/>
      </hudson.model.AllView>
    </views>
    <tabBar class="hudson.views.DefaultViewsTabBar"/>
  </folderViews>
  <healthMetrics>
    <com.cloudbees.hudson.plugins.folder.health.WorstChildHealthMetric>
      <nonRecursive>false</nonRecursive>
    </com.cloudbees.hudson.plugins.folder.health.WorstChildHealthMetric>
  </healthMetrics>
  <icon class="com.cloudbees.hudson.plugins.folder.icons.StockFolderIcon"/>
</com.cloudbees.hudson.plugins.folder.Folder>
`
}

func createBranchSource(info *gits.GitRepositoryInfo, credentials string) string {
	credXml := ""
	if credentials != "" {
		credXml = `		  <credentialsId>` + credentials + `</credentialsId>
`
	}
	return `
	    <source class="org.jenkinsci.plugins.github_branch_source.GitHubSCMSource" plugin="github-branch-source@2.3.1">
		  <id>b50ee5d4-cb45-42de-9140-d79330bab9ac</id>` + credXml + `
		  <repoOwner>` + info.Organisation + `</repoOwner>
		  <repository>` + info.Name + `</repository>
		  <traits>
			<org.jenkinsci.plugins.github__branch__source.BranchDiscoveryTrait>
			  <strategyId>1</strategyId>
			</org.jenkinsci.plugins.github__branch__source.BranchDiscoveryTrait>
			<org.jenkinsci.plugins.github__branch__source.OriginPullRequestDiscoveryTrait>
			  <strategyId>1</strategyId>
			</org.jenkinsci.plugins.github__branch__source.OriginPullRequestDiscoveryTrait>
			<org.jenkinsci.plugins.github__branch__source.ForkPullRequestDiscoveryTrait>
			  <strategyId>1</strategyId>
			  <trust class="org.jenkinsci.plugins.github_branch_source.ForkPullRequestDiscoveryTrait$TrustContributors"/>
			</org.jenkinsci.plugins.github__branch__source.ForkPullRequestDiscoveryTrait>
		  </traits>
		</source>
`
}

func CreateMultiBranchProjectXml(info *gits.GitRepositoryInfo, credentials string) string {
	return `<?xml version='1.0' encoding='UTF-8'?>
<org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject plugin="workflow-multibranch@2.16">
  <actions/>
  <description></description>
  <properties>
	<org.jenkinsci.plugins.pipeline.modeldefinition.config.FolderConfig plugin="pipeline-model-definition@1.2.4">
	  <dockerLabel></dockerLabel>
	  <registry plugin="docker-commons@1.9"/>
	</org.jenkinsci.plugins.pipeline.modeldefinition.config.FolderConfig>
  </properties>
  <folderViews class="jenkins.branch.MultiBranchProjectViewHolder" plugin="branch-api@2.0.15">
	<owner class="org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject" reference="../.."/>
  </folderViews>
  <healthMetrics>
	<com.cloudbees.hudson.plugins.folder.health.WorstChildHealthMetric plugin="cloudbees-folder@6.2.1">
	  <nonRecursive>false</nonRecursive>
	</com.cloudbees.hudson.plugins.folder.health.WorstChildHealthMetric>
  </healthMetrics>
  <icon class="jenkins.branch.MetadataActionFolderIcon" plugin="branch-api@2.0.15">
	<owner class="org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject" reference="../.."/>
  </icon>
  <orphanedItemStrategy class="com.cloudbees.hudson.plugins.folder.computed.DefaultOrphanedItemStrategy" plugin="cloudbees-folder@6.2.1">
	<pruneDeadBranches>true</pruneDeadBranches>
	<daysToKeep>-1</daysToKeep>
	<numToKeep>-1</numToKeep>
  </orphanedItemStrategy>
  <triggers/>
  <disabled>false</disabled>
  <sources class="jenkins.branch.MultiBranchProject$BranchSourceList" plugin="branch-api@2.0.15">
	<data>
	  <jenkins.branch.BranchSource>
` + createBranchSource(info, credentials) + `
		<strategy class="jenkins.branch.DefaultBranchPropertyStrategy">
		  <properties class="empty-list"/>
		</strategy>
	  </jenkins.branch.BranchSource>
	</data>
	<owner class="org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject" reference="../.."/>
  </sources>
  <factory class="org.jenkinsci.plugins.workflow.multibranch.WorkflowBranchProjectFactory">
	<owner class="org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject" reference="../.."/>
	<scriptPath>Jenkinsfile</scriptPath>
  </factory>
</org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject>
`
}
