package interfaces

// Repos contains details about repositories connected to an event.
// This is primarily used to aid with mapping the application code's filepath
// to the equivilent path inside of a repository.
//
// {
//     "/abs/path/to/sentry": {
//         "name": "getsentry/sentry",
//         "prefix": "src",
//         "revision": "..." // optional
//     }
// }
type Repos map[string]Repo

// Repo contains details about one repository for the Repos interface.
type Repo struct {
	Name     string  `kv:"name"     in:"name"     json:"name"`
	Prefix   string  `kv:"prefix"   in:"prefix"   json:"prefix"`
	Revision *string `kv:"revision" in:"revision" json:"revision"`
}

func init() {
	Register(&Repos{})
}

func (*Repos) KeyAlias() string {
	return "repos"
}

func (*Repos) KeyCanonical() string {
	return "sentry.interfaces.Repos"
}
