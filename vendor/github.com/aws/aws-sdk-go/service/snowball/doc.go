// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package snowball provides the client and types for making API
// requests to Amazon Import/Export Snowball.
//
// AWS Snowball is a petabyte-scale data transport solution that uses secure
// appliances to transfer large amounts of data between your on-premises data
// centers and Amazon Simple Storage Service (Amazon S3). The Snowball commands
// described here provide access to the same functionality that is available
// in the AWS Snowball Management Console, which enables you to create and manage
// jobs for Snowball. To transfer data locally with a Snowball appliance, you'll
// need to use the Snowball client or the Amazon S3 API adapter for Snowball.
// For more information, see the User Guide (http://docs.aws.amazon.com/AWSImportExport/latest/ug/api-reference.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30 for more information on this service.
//
// See snowball package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/snowball/
//
// Using the Client
//
// To use the client for Amazon Import/Export Snowball you will first need
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
//   svc := snowball.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Import/Export Snowball client Snowball for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/snowball/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.CancelCluster(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("CancelCluster result:")
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
//   result, err := svc.CancelClusterWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package snowball
