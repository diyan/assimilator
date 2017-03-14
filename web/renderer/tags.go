package renderer

import (
	"encoding/json"
	"fmt"

	"github.com/diyan/assimilator/models"

	"github.com/flosch/pongo2"
)

type noOpNode struct{}

type assetURLNode struct {
	module   string
	path     string
	absolute bool
}

type reactConfigNode struct {
	Config ReactConfig
}

// ReactConfig ..
// TODO Consider use *url.URL, *mail.Address instead of string type and use custom JSON marshaller
type ReactConfig struct {
	DSN                string        `json:"dsn"`
	NeedsUpgrade       bool          `json:"needsUpgrade"`
	Version            SentryVersion `json:"version"`
	Features           []string      `json:"features"`
	IsAuthenticated    bool          `json:"isAuthenticated"`
	MediaURL           string        `json:"mediaUrl"`
	SingleOrganization bool          `json:"singleOrganization"`
	Messages           []string      `json:"messages"` // TODO double check this
	URLPrefix          string        `json:"urlPrefix"`
	User               User          `json:"user"`
}

// SentryVersion ..
type SentryVersion struct {
	foo              string
	Current          string `json:"current"` // TODO int.int.int
	Build            string `json:"build"`   // int.int.int
	UpgradeAvailable bool   `json:"upgradeAvailable"`
	Latest           string `json:"latest"` // int.int.int
}

// User ..
type User struct {
	// TODO investigate why embedding models.User does not work
	models.User
	ID       int    `db:"id" json:"id,string"`
	Username string `db:"username" json:"username"`
	// this column is called first_name for legacy reasons, but it is the entire
	//   display name
	Name        string `db:"first_name" json:"name"`
	Email       string `db:"email" json:"email"`
	IsSuperuser bool   `db:"is_superuser" json:"isSuperuser"`

	AvatarURL string      `json:"avatarUrl"`
	Options   UserOptions `json:"options"`
}

// UserOptions ..
type UserOptions struct {
	Timezone        string `json:"timezone"`        // TODO double check this
	StacktraceOrder string `json:"stacktraceOrder"` // default
	Language        string `json:"language"`
	Clock24Hours    bool   `json:"clock24Hours"`
}

type publicDsnNode struct {
	DSN string
}

type getUserContextNode struct {
	name string
}

// UserContext ..
type UserContext struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email"`
	IPAddress string `json:"ip_address"`
}

// RegisterTags ...
func RegisterTags() {
	pongo2.RegisterTag("load", noOpParser)
	// TODO implement those tags
	// https://www.florian-schlachter.de/post/pongo2/
	// https://github.com/flosch/pongo2/blob/master/tags_comment.go
	// https://github.com/flosch/pongo2/blob/master/filters_builtin.go
	// NOTE pongo2 failed to parse string arguments with single quotes
	//   {% asset_url 'sentry' 'js/ads.js' %}
	pongo2.RegisterTag("get_sentry_version", noOpParser)
	pongo2.RegisterTag("absolute_asset_url", absoluteAssetURLParser)
	pongo2.RegisterTag("asset_url", assetURLParser)
	pongo2.RegisterTag("crossorigin", noOpParser)
	pongo2.RegisterTag("locale_js_include", noOpParser)
	pongo2.RegisterTag("get_react_config", getReactConfigParser)
	pongo2.RegisterTag("show_system_status", noOpParser)
	pongo2.RegisterTag("serialize_detailed_org", noOpParser)
	pongo2.RegisterTag("public_dsn", publicDsnParser)
	pongo2.RegisterTag("convert_to_json", convertToJSONParser)
	pongo2.RegisterTag("get_user_context", getUserContextParser)
	pongo2.RegisterTag("trans", noOpParser)
	pongo2.RegisterTag("letter_avatar_svg", noOpParser)
	pongo2.RegisterTag("profile_photo_url", noOpParser)
	pongo2.RegisterTag("gravatar_url", noOpParser)
	pongo2.RegisterTag("url", noOpParser)
	pongo2.RegisterTag("feature", noOpParser)
	pongo2.RegisterTag("endfeature", noOpParser)
}

func (node *noOpNode) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	return nil
}

func noOpParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	return &noOpNode{}, nil
}

func (node *reactConfigNode) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	bytes, err := json.Marshal(node.Config)
	if err != nil {
		panic(err)
	}
	// TODO investigate how to disable HTML encoding in the config.user.avatarUrl
	writer.Write(bytes)
	return nil
}

func getReactConfigParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	node := &reactConfigNode{}
	node.Config = ReactConfig{
		DSN:          "http://535d385540bc43c496aa02d14db75fc6@localhost:9001/1",
		NeedsUpgrade: false,
		Version: SentryVersion{
			Current:          "8.12.0",
			Build:            "8.12.0",
			UpgradeAvailable: false,
			Latest:           "8.12.0",
		},
		Features:           []string{},
		IsAuthenticated:    true,
		MediaURL:           "/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/",
		SingleOrganization: true,
		Messages:           []string{},
		URLPrefix:          "http://localhost:9000",
		User: User{
			Username:    "admin",
			Name:        "alexey.diyan@gmail.com",
			AvatarURL:   "https://secure.gravatar.com/avatar/16b920fc5dcf06585b4c205373c2ca1c?s=32&d=mm",
			IsSuperuser: true,
			Options: UserOptions{
				Timezone:        "UTC",
				StacktraceOrder: "default",
				Language:        "en",
				Clock24Hours:    false,
			},
			ID:    1,
			Email: "alexey.diyan@gmail.com",
		},
	}
	return node, nil
}

func (node *publicDsnNode) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	writer.WriteString(node.DSN)
	return nil
}

func publicDsnParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	node := &publicDsnNode{}
	// TODO move publicDsnNode.DSN and reactConfigNode.Config.DSN to the settings
	node.DSN = "http://535d385540bc43c496aa02d14db75fc6@localhost:9001/1"
	return node, nil
}

type convertToJSONNode struct {
	name string
}

func (node *convertToJSONNode) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	obj, ok := ctx.Public[node.name]
	if !ok {
		// TODO return pongo's error or write to the log
		return nil
	}
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	writer.Write(bytes)
	return nil
}

func convertToJSONParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	node := &convertToJSONNode{}
	if arguments.Count() != 1 {
		return nil, arguments.Error(
			"Tag 'convert_to_json' requires 'var' argument as identifier.",
			nil)
	}
	if name := arguments.MatchType(pongo2.TokenIdentifier); name != nil {
		node.name = name.Val
	} else {
		return nil, arguments.Error(
			"Tag 'convert_to_json' requires 'var' argument as identifier.",
			nil)
	}
	return node, nil
}

func (node *getUserContextNode) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	_, ok := ctx.Public[node.name]
	if !ok {
		// TODO return pongo's error or write to the log
		return nil
	}
	// TODO stub implementation
	user := UserContext{
		ID:        1,
		Email:     "alexey.diyan@gmail.com",
		IPAddress: "192.169.100.1",
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	writer.Write(bytes)
	return nil
}

func getUserContextParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	node := &getUserContextNode{}
	if arguments.Count() != 1 {
		return nil, arguments.Error(
			"Tag 'get_user_context' requires 'request' argument as identifier.",
			nil)
	}
	if name := arguments.MatchType(pongo2.TokenIdentifier); name != nil {
		node.name = name.Val
	} else {
		return nil, arguments.Error(
			"Tag 'get_user_context' requires 'request' argument as identifier.",
			nil)
	}
	return node, nil
}

/*
JS 3
Raven.config('http://535d385540bc43c496aa02d14db75fc6@localhost:9000/1', {
  release: '8.0.6',
  whitelistUrls: ""
}).install();
Raven.setUserContext(
	{"ip_address":"192.169.100.1",
		"email":"alexey.diyan@gmail.com",
		"id":1});


Raven.config('{% public_dsn %}', {
  release: '{{ sentry_version.build }}',
  whitelistUrls: {% convert_to_json ALLOWED_HOSTS %}
}).install();
Raven.setUserContext({% get_user_context request %});

@register.simple_tag
def public_dsn():
    project_id = settings.SENTRY_FRONTEND_PROJECT or settings.SENTRY_PROJECT
    cache_key = 'dsn:%s' % (project_id,)

    result = default_cache.get(cache_key)
    if result is None:
        key = _get_project_key(project_id)
        if key:
            result = key.dsn_public
        else:
            result = ''
        default_cache.set(cache_key, result, 60)
    return result
*/

// TODO implement get_sentry_version pango2 tag that will set tuple below into
//   context as 'sentry_version' variable
// SentryVersion = namedtuple('SentryVersion', [
//    'current', 'latest', 'update_available', 'build',
// ])

func (node *assetURLNode) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	// TODO log on error, consider to return it
	url := ""
	if node.absolute {
		url = "http://localhost:9001/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/%s/%s"
	} else {
		url = "/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/%s/%s"
	}
	writer.WriteString(fmt.Sprintf(url, node.module, node.path))
	return nil
}

// Returns a versioned absolute asset URL (located within Sentry's static files).
//
// Example:
//   {% absolute_asset_url 'sentry' 'dist/sentry.css' %}
//   =>  "http://sentry.example.com/_static/74d127b78dc7daf2c51f/sentry/dist/sentry.css"
func absoluteAssetURLParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	pongoNode, err := assetURLParser(doc, start, arguments)
	if err != nil {
		return nil, err
	}
	node, ok := pongoNode.(*assetURLNode)
	if ok {
		node.absolute = true
	}
	return node, err
}

// Returns a versioned asset URL (located within Sentry's static files).
//
// Example:
//   {% asset_url 'sentry' 'dist/sentry.css' %}
//   =>  "/_static/74d127b78dc7daf2c51f/sentry/dist/sentry.css"
func assetURLParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	node := &assetURLNode{}
	if arguments.Count() != 2 {
		return nil, arguments.Error(
			"Tag 'absolute_asset_url' requires 'module' and 'path' arguments.",
			nil)
	}
	if module := arguments.MatchType(pongo2.TokenString); module != nil {
		node.module = module.Val
	} else {
		return nil, arguments.Error(
			"Tag 'absolute_asset_url' requires 'module' first argument as string.",
			nil)
	}
	if path := arguments.MatchType(pongo2.TokenString); path != nil {
		node.path = path.Val
	} else {
		return nil, arguments.Error(
			"Tag 'absolute_asset_url' requires 'path' second argument as string.",
			nil)
	}
	return node, nil
}
