package main

import (
	"github.com/hortonworks/gohadoop/hadoop_yarn"
	"github.com/hortonworks/gohadoop/hadoop_yarn/conf"
	"github.com/hortonworks/gohadoop/hadoop_yarn/yarn_client"
	"log"
	"time"
)

/*

2014/08/28 18:45:50 Successfully connected <clientId:ee71c86d-f50a-4052-53c0-14e700bd9c15, server:36citusm1.in.cfops.it:8032>
2014/08/28 18:44:06 Received rpcResponseHeaderProto = {CallId:0xc2080871a0 Status:SUCCESS ServerIpcVersionNum:0xc2080871a8 ExceptionClassName:<nil> ErrorMsg:<nil> ErrorDetail:<nil> ClientId:[111 140 57 186 25 245 69 157 88 133 44 28 18 95 114 104] RetryCount:0xc2080871ac XXX_unrecognized:[]}

2014/08/28 18:44:07 Successfully connected <clientId:fa10c6ae-84fd-4417-6816-4c1ee3594ba3, server:36citusm1.in.cfops.it:8030>
2014/08/28 18:44:07 rmClient.RegisterApplicationMaster org.apache.hadoop.security.AccessControlException:SIMPLE authentication is not enabled.  Available:[TOKEN]:FATAL_UNAUTHORIZED:ServerDidNotSetErrorDetail

https://github.com/apache/hadoop-common/commit/68467a22
https://issues.apache.org/jira/browse/YARN-961

*/

func main() {
	var err error

	// Create YarnConfiguration
	conf, err := conf.NewYarnConfiguration()
	if err != nil {
		panic(err)
	}

	// Create YarnClient
	yarnClient, err := yarn_client.CreateYarnClient(conf)
	if err != nil {
		panic(err)
	}

	// Create new application to get ApplicationSubmissionContext
	_, asc, err := yarnClient.CreateNewApplication()
	if err != nil {
		panic(err)
	}

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
