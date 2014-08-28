package main

import (
	"github.com/hortonworks/gohadoop/hadoop_yarn"
	"github.com/hortonworks/gohadoop/hadoop_yarn/conf"
	"github.com/hortonworks/gohadoop/hadoop_yarn/yarn_client"
	"log"
	"time"
)

func main() {
	var err error

	// Create YarnConfiguration
	conf, _ := conf.NewYarnConfiguration()

	// Create YarnClient
	yarnClient, _ := yarn_client.CreateYarnClient(conf)

	// Create new application to get ApplicationSubmissionContext
	_, asc, _ := yarnClient.CreateNewApplication()

	// Some useful information
	queue := "default"
	appName := "simple-go-yarn-app"
	appType := "GO_HADOOP_UNMANAGED"
	unmanaged := true
	clc := hadoop_yarn.ContainerLaunchContextProto{}

	// Setup ApplicationSubmissionContext for the application
	asc.AmContainerSpec = &clc
	asc.ApplicationName = &appName
	asc.Queue = &queue
	asc.ApplicationType = &appType
	asc.UnmanagedAm = &unmanaged

	// Submit!
	err = yarnClient.SubmitApplication(asc)
	if err != nil {
		log.Fatal("yarnClient.SubmitApplication ", err)
	}
	log.Println("Successfully submitted unmanaged application: ", asc.ApplicationId)
	time.Sleep(1 * time.Second)

	appReport, err := yarnClient.GetApplicationReport(asc.ApplicationId)
	if err != nil {
		log.Fatal("yarnClient.GetApplicationReport ", err)
	}
	appState := appReport.GetYarnApplicationState()
	for appState != hadoop_yarn.YarnApplicationStateProto_ACCEPTED {
		log.Println("Application in state ", appState)
		time.Sleep(1 * time.Second)
		appReport, err = yarnClient.GetApplicationReport(asc.ApplicationId)
		appState = appReport.GetYarnApplicationState()
		if appState == hadoop_yarn.YarnApplicationStateProto_FAILED || appState == hadoop_yarn.YarnApplicationStateProto_KILLED {
			log.Fatal("Application in state ", appState)
		}
	}
	log.Printf("Application in state %v", appState)

	// Create AMRMClient
	var attemptId int32
	attemptId = 1
	applicationAttemptId := hadoop_yarn.ApplicationAttemptIdProto{ApplicationId: asc.ApplicationId, AttemptId: &attemptId}

	rmClient, _ := yarn_client.CreateAMRMClient(conf, &applicationAttemptId)
	log.Printf("Created RM client: %v", rmClient)

	// Wait for ApplicationAttempt to be in Launched state
	appAttemptReport, err := yarnClient.GetApplicationAttemptReport(&applicationAttemptId)
	appAttemptState := appAttemptReport.GetYarnApplicationAttemptState()
	for appAttemptState != hadoop_yarn.YarnApplicationAttemptStateProto_APP_ATTEMPT_LAUNCHED {
		log.Println("ApplicationAttempt in state ", appAttemptState)
		time.Sleep(1 * time.Second)
		appAttemptReport, err = yarnClient.GetApplicationAttemptReport(&applicationAttemptId)
		appAttemptState = appAttemptReport.GetYarnApplicationAttemptState()
	}
	log.Printf("ApplicationAttempt in state %v", appAttemptState)

	// SaslRpcClient.java
	ipcClient := rmClient.Client
	_ = ipcClient

	// Register with ResourceManager
	log.Println("About to register application master.")
	err = rmClient.RegisterApplicationMaster("", -1, "")
	if err != nil {
		log.Fatal("rmClient.RegisterApplicationMaster ", err)
	}
	log.Println("Successfully registered application master.")

	// Add resource requests
	const numContainers = int32(1)
	memory := int32(128)
	resource := hadoop_yarn.ResourceProto{Memory: &memory}
	rmClient.AddRequest(1, "*", &resource, numContainers)

	// Now call ResourceManager.allocate
	allocateResponse, err := rmClient.Allocate()
	if err == nil {
		log.Println("allocateResponse: ", *allocateResponse)
	}
	log.Println("#containers allocated: ", len(allocateResponse.AllocatedContainers))

	numAllocatedContainers := int32(0)
	allocatedContainers := make([]*hadoop_yarn.ContainerProto, numContainers, numContainers)
	for numAllocatedContainers < numContainers {
		// Sleep for a while
		log.Println("Sleeping...")
		time.Sleep(3 * time.Second)
		log.Println("Sleeping... done!")

		// Try to get containers now...
		allocateResponse, err = rmClient.Allocate()
		if err == nil {
			log.Println("allocateResponse: ", *allocateResponse)
		}

		for _, container := range allocateResponse.AllocatedContainers {
			allocatedContainers[numAllocatedContainers] = container
			numAllocatedContainers++
		}

		log.Println("#containers allocated: ", len(allocateResponse.AllocatedContainers))
		log.Println("Total #containers allocated so far: ", numAllocatedContainers)
	}
	log.Println("Final #containers allocated: ", numAllocatedContainers)

	// Now launch containers
	containerLaunchContext := hadoop_yarn.ContainerLaunchContextProto{Command: []string{"/bin/date"}}
	log.Println("containerLaunchContext: ", containerLaunchContext)
	for _, container := range allocatedContainers {
		log.Println("Launching container: ", *container, " ", container.NodeId.Host, ":", container.NodeId.Port)
		nmClient, err := yarn_client.CreateAMNMClient(*container.NodeId.Host, int(*container.NodeId.Port))
		if err != nil {
			log.Fatal("hadoop_yarn.DialContainerManagementProtocolService: ", err)
		}
		log.Println("Successfully created nmClient: ", nmClient)
		err = nmClient.StartContainer(container, &containerLaunchContext)
		if err != nil {
			log.Fatal("nmClient.StartContainer: ", err)
		}
	}

	// Wait for Containers to complete
	numCompletedContainers := int32(0)
	for numCompletedContainers < numContainers {
		// Sleep for a while
		log.Println("Sleeping...")
		time.Sleep(3 * time.Second)
		log.Println("Sleeping... done!")

		allocateResponse, err = rmClient.Allocate()
		if err == nil {
			log.Println("allocateResponse: ", *allocateResponse)
		}

		for _, containerStatus := range allocateResponse.CompletedContainerStatuses {
			log.Println("Completed container: ", *containerStatus, " exit-code: ", *containerStatus.ExitStatus)
			numCompletedContainers++
		}
	}
	log.Println("Containers complete: ", numCompletedContainers)

	// Unregister with ResourceManager
	log.Println("About to unregister application master.")
	finalStatus := hadoop_yarn.FinalApplicationStatusProto_APP_SUCCEEDED
	err = rmClient.FinishApplicationMaster(&finalStatus, "done", "")
	if err != nil {
		log.Fatal("rmClient.RegisterApplicationMaster ", err)
	}
	log.Println("Successfully unregistered application master.")
}

/*

  <name>hadoop.rpc.protection</name>
  <value>authentication</value>
  <description>A comma-separated list of protection values for secured sasl
      connections. Possible values are authentication, integrity and privacy.
      authentication means authentication only and no integrity or privacy;
      integrity implies authentication and integrity are enabled; and privacy
      implies all of authentication, integrity and privacy are enabled.
      hadoop.security.saslproperties.resolver.class can be used to override
      the hadoop.rpc.protection for a connection at the server side.

*/

/*
Two types of protocols support Kerberos authentication: RPC and HTTP. The RPC protocol uses SASL GSSAPI while the HTTP protocol uses SPNego.
*/

/*
org.apache.hadoop.ipc.RPC
Server IPC version 9
*/

/*
2014-08-27 20:31:09,521 INFO org.apache.hadoop.ipc.Server: Socket Reader #1 for port 8030: readAndProcess from client 10.36.10.7 threw exception [org.apache.hadoop.security.AccessControlException: SIMPLE authentication is not enabled.  Available:[TOKEN]]
org.apache.hadoop.security.AccessControlException: SIMPLE authentication is not enabled.  Available:[TOKEN]
	at org.apache.hadoop.ipc.Server$Connection.initializeAuthContext(Server.java:1542)
	at org.apache.hadoop.ipc.Server$Connection.readAndProcess(Server.java:1498)
	at org.apache.hadoop.ipc.Server$Listener.doRead(Server.java:750)
	at org.apache.hadoop.ipc.Server$Listener$Reader.doRunLoop(Server.java:624)
	at org.apache.hadoop.ipc.Server$Listener$Reader.run(Server.java:595)
*/

/*
2014-08-27 20:35:43,711 WARN org.apache.hadoop.yarn.server.resourcemanager.rmapp.RMAppImpl: The specific max attempts: 0 for application: 3 is invalid, because it is out of the range [1, 2]. Use the global max attempts instead.
*/