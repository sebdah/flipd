namespace go flipd

/******************************************************
*                       Types
/*****************************************************/

typedef string Key

/******************************************************
*                       Enums
/*****************************************************/

// FeatureStatus indicates the status of the feature.
enum FeatureStatus {
    ENABLED = 1,
    DISABLED = 2
}

/******************************************************
*                      Objects
/*****************************************************/

// Feature is the core component used to define and communicate information
// about a feature.
struct Feature {
    // The key is a unique identifier for a feature. Typically features should
    // be grouped with a namespacing pattern like: "namespace:feature_name" or
    // such.
    //
    // Default:     None
    // Required:    True (InvalidInput exception is raised otherwise)
    1: optional Key key

    // Feature status of the new feature.
    //
    // Default:     FeatureStatus.DISABLED
    // Required:    False
    2: optional FeatureStatus status = FeatureStatus.DISABLED
}

/******************************************************
*               Requests and responses
/*****************************************************/

// DeregisterFeatureRequest is used when deregistering features in the service.
struct DeregisterFeatureRequest {
    // Feature holds the definition of the new feature to be registered.
    1: optional Key key
}

// GetFeaturesResponse is the response object returned when requesting a list
// of features registered with the flipd service.
struct GetFeaturesResponse {
    // List of registered features.
    1: optional list<Feature> features
}

// GetFeaturesRequest is used when feching a list of features from the flipd
// service.
struct GetFeaturesRequest {
    // Send a list of keys to fetch. If omitted all registered keys will be
    // returned. If the keys does not exist, they are silently ignored.
    //
    // Default:     None
    // Required:    True
    1: optional list<Key> keys
}

// RegisterFeatureRequest is used when registering new features in the service.
struct RegisterFeatureRequest {
    // Feature holds the definition of the new feature to be registered.
    1: optional Feature feature
}

/******************************************************
*                     Exceptions
/*****************************************************/

// DuplicateException is retuned when a uniqueness constraint is broken. Could
// for example be returned if the combination of namespace and feature name is
// already taken.
exception DuplicateException {
    1: optional string message
}

// InternalErrorException is used to communicate unexpected errors in the
// server.
exception InternalErrorException {
    1: optional string message
}

// InvalidInputException is thrown if the input from the client do not match
// the expectations of the server.
exception InvalidInputException {
    1: optional string message
}

// NotFound is thrown if the requested entity does not exist.
exception NotFoundException {
    1: optional string message
}

/******************************************************
*                    Services
/*****************************************************/
service flipd {
    // deregisterFeature removes a feature from the flipd service.
    void deregisterFeature(1: DeregisterFeatureRequest request)
        throws (
            1: InternalErrorException internalError
            2: InvalidInputException InvalidInput
            3: NotFoundException notFound)

    // getFeatures returns a list of features registered with the flipd
    // service.
    //
    // If no keys are provided, all features will be returned.
    GetFeaturesResponse getFeatures(1: GetFeaturesRequest request)
        throws (
            1: InternalErrorException internalError)

    // ping returns a basic string, "pong". Can be used to check if the service
    // is up.
    string ping()

    // registerFeature adds new features to the flipd service.
    void registerFeature(1: RegisterFeatureRequest request)
        throws (
            1: InternalErrorException internalError
            2: InvalidInputException InvalidInput
            3: DuplicateException duplicate)
}
