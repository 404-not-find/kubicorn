// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package gamelift provides the client and types for making API
// requests to Amazon GameLift.
//
// Amazon GameLift is a managed service for developers who need a scalable,
// dedicated server solution for their multiplayer games. Amazon GameLift provides
// tools for the following tasks: (1) acquire computing resources and deploy
// game servers, (2) scale game server capacity to meet player demand, (3) host
// game sessions and manage player access, and (4) track in-depth metrics on
// player usage and server performance.
//
// The Amazon GameLift service API includes two important function sets:
//
//    * Manage game sessions and player access – Retrieve information on available
//    game sessions; create new game sessions; send player requests to join
//    a game session.
//
//    * Configure and manage game server resources – Manage builds, fleets,
//    queues, and aliases; set autoscaling policies; retrieve logs and metrics.
//
// This reference guide describes the low-level service API for Amazon GameLift.
// You can use the API functionality with these tools:
//
//    * The Amazon Web Services software development kit (AWS SDK (http://aws.amazon.com/tools/#sdk))
//    is available in multiple languages (http://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-supported.html#gamelift-supported-clients)
//    including C++ and C#. Use the SDK to access the API programmatically from
//    an application, such as a game client.
//
//    * The AWS command-line interface (http://aws.amazon.com/cli/) (CLI) tool
//    is primarily useful for handling administrative actions, such as setting
//    up and managing Amazon GameLift settings and resources. You can use the
//    AWS CLI to manage all of your AWS services.
//
//    * The AWS Management Console (https://console.aws.amazon.com/gamelift/home)
//    for Amazon GameLift provides a web interface to manage your Amazon GameLift
//    settings and resources. The console includes a dashboard for tracking
//    key resources, includings builds and fleets, and displays usage and performance
//    metrics for your games as customizable graphs.
//
//    * Amazon GameLift Local is a tool for testing your game's integration
//    with Amazon GameLift before deploying it on the service. This tools supports
//    a subset of key API actions, which can be called from either the AWS CLI
//    or programmatically. See Testing an Integration (http://docs.aws.amazon.com/gamelift/latest/developerguide/integration-testing-local.html).
//
// MORE RESOURCES
//
//    * Amazon GameLift Developer Guide (http://docs.aws.amazon.com/gamelift/latest/developerguide/)
//    – Learn more about Amazon GameLift features and how to use them.
//
//    * Lumberyard and Amazon GameLift Tutorials (https://gamedev.amazon.com/forums/tutorials)
//    – Get started fast with walkthroughs and sample projects.
//
//    * GameDev Blog (http://aws.amazon.com/blogs/gamedev/) – Stay up to date
//    with new features and techniques.
//
//    * GameDev Forums (https://gamedev.amazon.com/forums/spaces/123/gamelift-discussion.html)
//    – Connect with the GameDev community.
//
//    * Amazon GameLift Document History (http://docs.aws.amazon.com/gamelift/latest/developerguide/doc-history.html)
//    – See changes to the Amazon GameLift service, SDKs, and documentation,
//    as well as links to release notes.
//
// API SUMMARY
//
// This list offers a functional overview of the Amazon GameLift service API.
//
// Managing Games and Players
//
// These actions allow you to start new game sessions, find existing game sessions,
// track status and other game session information, and enable access for players
// to join game sessions.
//
//    * Discover existing game sessions
//
// SearchGameSessions – Get all available game sessions or search for game sessions
//    that match a set of criteria.
//
//    * Start a new game session
//
// Game session placement – Use a queue to process requests for new game sessions
//    and place them on the best available fleet. Placement requests are asynchronous;
//    game sessions are started whenever acceptable resources become available.
//
//
// StartGameSessionPlacement – Request a new game session placement and add
//    one or more players to it.
//
// DescribeGameSessionPlacement – Get details on a placement request, including
//    status.
//
// StopGameSessionPlacement – Cancel a placement request.
//
// CreateGameSession – Request a new game session on a specific fleet. Available
//    in Amazon GameLift Local.
//
//    * Manage game session data
//
// DescribeGameSessions – Retrieve metadata for one or more game sessions, including
//    length of time active and current player count. Available in Amazon GameLift
//    Local.
//
// DescribeGameSessionDetails – Retrieve metadata and the game session protection
//    setting for one or more game sessions.
//
// UpdateGameSession – Change game session settings, such as maximum player
//    count and join policy.
//
// GetGameSessionLogUrl – Get the location of saved logs for a game session.
//
//    * Manage player sessions
//
// CreatePlayerSession – Send a request for a player to join a game session.
//    Available in Amazon GameLift Local.
//
// CreatePlayerSessions – Send a request for multiple players to join a game
//    session. Available in Amazon GameLift Local.
//
// DescribePlayerSessions – Get details on player activity, including status,
//    playing time, and player data. Available in Amazon GameLift Local.
//
// Setting Up and Managing Game Servers
//
// When setting up Amazon GameLift resources for your game, you first create
// a game build (http://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html)
// and upload it to Amazon GameLift. You can then use these actions to configure
// and manage a fleet of resources to run your game servers, scale capacity
// to meet player demand, access performance and utilization metrics, and more.
//
//    * Manage game builds
//
// CreateBuild – Create a new build using files stored in an Amazon S3 bucket.
//    (Update uploading permissions with RequestUploadCredentials.) To create
//    a build and upload files from a local path, use the AWS CLI command upload-build.
//
// ListBuilds – Get a list of all builds uploaded to a Amazon GameLift region.
//
// DescribeBuild – Retrieve information associated with a build.
//
// UpdateBuild – Change build metadata, including build name and version.
//
// DeleteBuild – Remove a build from Amazon GameLift.
//
//    * Manage fleets
//
// CreateFleet – Configure and activate a new fleet to run a build's game servers.
//
// ListFleets – Get a list of all fleet IDs in a Amazon GameLift region (all
//    statuses).
//
// DeleteFleet – Terminate a fleet that is no longer running game servers or
//    hosting players.
//
// View / update fleet configurations.
//
// DescribeFleetAttributes / UpdateFleetAttributes – View or change a fleet's
//    metadata and settings for game session protection and resource creation
//    limits.
//
// DescribeFleetPortSettings / UpdateFleetPortSettings – View or change the
//    inbound permissions (IP address and port setting ranges) allowed for a
//    fleet.
//
// DescribeRuntimeConfiguration / UpdateRuntimeConfiguration – View or change
//    what server processes (and how many) to run on each instance in a fleet.
//
//    * Control fleet capacity
//
// DescribeEC2InstanceLimits – Retrieve maximum number of instances allowed
//    for the current AWS account and the current usage level.
//
// DescribeFleetCapacity / UpdateFleetCapacity – Retrieve the capacity settings
//    and the current number of instances in a fleet; adjust fleet capacity
//    settings to scale up or down.
//
// Autoscale – Manage autoscaling rules and apply them to a fleet.
//
// PutScalingPolicy – Create a new autoscaling policy, or update an existing
//    one.
//
// DescribeScalingPolicies – Retrieve an existing autoscaling policy.
//
// DeleteScalingPolicy – Delete an autoscaling policy and stop it from affecting
//    a fleet's capacity.
//
//    * Access fleet activity statistics
//
// DescribeFleetUtilization – Get current data on the number of server processes,
//    game sessions, and players currently active on a fleet.
//
// DescribeFleetEvents – Get a fleet's logged events for a specified time span.
//
// DescribeGameSessions – Retrieve metadata associated with one or more game
//    sessions, including length of time active and current player count.
//
//    * Remotely access an instance
//
// DescribeInstances – Get information on each instance in a fleet, including
//    instance ID, IP address, and status.
//
// GetInstanceAccess – Request access credentials needed to remotely connect
//    to a specified instance in a fleet.
//
//    * Manage fleet aliases
//
// CreateAlias – Define a new alias and optionally assign it to a fleet.
//
// ListAliases – Get all fleet aliases defined in a Amazon GameLift region.
//
// DescribeAlias – Retrieve information on an existing alias.
//
// UpdateAlias – Change settings for a alias, such as redirecting it from one
//    fleet to another.
//
// DeleteAlias – Remove an alias from the region.
//
// ResolveAlias – Get the fleet ID that a specified alias points to.
//
//    * Manage game session queues
//
// CreateGameSessionQueue – Create a queue for processing requests for new game
//    sessions.
//
// DescribeGameSessionQueues – Get data on all game session queues defined in
//    a Amazon GameLift region.
//
// UpdateGameSessionQueue – Change the configuration of a game session queue.
//
// DeleteGameSessionQueue – Remove a game session queue from the region.
//
// See https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01 for more information on this service.
//
// See gamelift package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/gamelift/
//
// Using the Client
//
// To use the client for Amazon GameLift you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := gamelift.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon GameLift client GameLift for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/gamelift/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.CreateAlias(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("CreateAlias result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.CreateAliasWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package gamelift
