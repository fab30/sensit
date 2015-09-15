/*
Package server contains all the types and methods that handle serving an http
API to receive the callbacks from the sensit server.
For more information about these callbacks, please refer to https://api.sensit.io/v1

The API server provides two endpoints :
  /ping                             // responds Pong to the request. This endpoint is usefull to healthcheck the server
  /api/v1/{deviceID}/temperature    // receives the callbacks

*/
package server
