// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package swf provides the client and types for making API
// requests to Amazon Simple Workflow Service.
//
// The Amazon Simple Workflow Service (Amazon SWF) makes it easy to build applications
// that use Amazon's cloud to coordinate work across distributed components.
// In Amazon SWF, a task represents a logical unit of work that is performed
// by a component of your workflow. Coordinating tasks in a workflow involves
// managing intertask dependencies, scheduling, and concurrency in accordance
// with the logical flow of the application.
//
// Amazon SWF gives you full control over implementing tasks and coordinating
// them without worrying about underlying complexities such as tracking their
// progress and maintaining their state.
//
// This documentation serves as reference only. For a broader overview of the
// Amazon SWF programming model, see the Amazon SWF Developer Guide (http://docs.aws.amazon.com/amazonswf/latest/developerguide/).
//
// See swf package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/swf/
//
// Using the Client
//
// To use the client for Amazon Simple Workflow Service you will first need
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
//   svc := swf.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Simple Workflow Service client SWF for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/swf/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.CountClosedWorkflowExecutions(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("CountClosedWorkflowExecutions result:")
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
//   result, err := svc.CountClosedWorkflowExecutionsWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package swf
