package models

// Payload holds the metrics we receive in a request.
type Payload struct {
	User      *User      // payload is specific to a user
	Server    *Server    // payload is specific to a server or mobile app
	MobileApp *MobileApp // payload is specific to a server or mobile app
	Metric    *Metric    // the collected metric from server or mobile app
}
