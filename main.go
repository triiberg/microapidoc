package microapidoc

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// always sorts by group and then by method
type GeneralDoc struct {
	BuildTag                 string
	Started                  string
	SearchControllersIn      string
	SearchModelsIn           string
	AllRoutes                []RouteInfo
	Name                     string
	BaseUrl                  string
	HeaderColor              string
	AuthHeaderDefaultOn      bool     // if true, the auth header will be added to all endpoints, except "noauth" endpoints
	AuthHeaderNames          []string // Authorization header like "Bearer <token>"
	HighlightResponseHeaders []string // like x-header:ok:green and x-header:error:red or
	TSModelPrefix            string   // prefix for the typescript model
}

type ConfigPayload struct {
	Title           string            `json:"title,omitempty"`
	BuildTag        string            `json:"buildTag,omitempty"`        // for the build tag
	Started         string            `json:"started,omitempty"`         // for the start time
	AuthDefaultMode bool              `json:"authDefaultMode,omitempty"` // if true, the auth header will be added to all endpoints, except "noauth" endpoints
	HeaderColor     string            `json:"headerColor,omitempty"`
	ResponseHeaders []ResponseHeaders `json:"responseHeaders,omitempty"`
	AuthHeaders     []AuthHeaders     `json:"authHeaders,omitempty"`
	Groups          []EndpointGroups  `json:"groups,omitempty"`
	ExampleJsons    map[string]string `json:"exampleJsons,omitempty"` // json examples for the models
}
type ResponseHeaders struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Color string `json:"color,omitempty"`
}
type AuthHeaders struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type InputParameter struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type Response struct {
}
type Endpoints struct {
	Name             string           `json:"name,omitempty"`
	Summary          string           `json:"description,omitempty"`
	Method           string           `json:"method,omitempty"`
	URL              string           `json:"url,omitempty"`
	Label            []string         `json:"label,omitempty"` // like tested:green and experimental:yellow
	AuthHeaderOn     bool             `json:"authHeaderOn,omitempty"`
	AuthHeaders      []InputParameter `json:"authHeaders,omitempty"`
	HeaderParameters []InputParameter `json:"headerParameters,omitempty"`
	PathParameters   []InputParameter `json:"pathParameters,omitempty"`
	QueryParameters  []InputParameter `json:"queryParameters,omitempty"`
	BodyParameters   []InputParameter `json:"bodyParameters,omitempty"`
	Response         Response         `json:"response,omitempty"`
}
type EndpointGroups struct {
	Name      string      `json:"name,omitempty"`
	Endpoints []Endpoints `json:"endpoints,omitempty"`
}

type OneComment struct {
	FileName         string
	FunctionName     string
	Group            string
	Summary          string
	GoodResponse     string
	BadResponse      string
	Label            []string
	AuthNotDefault   bool
	HeaderParameters []InputParameter
	PathParameters   []InputParameter
	QueryParameters  []InputParameter
	BodyParameters   []InputParameter
}

func GenerateJSONExamplesFromModels(root string) (map[string]string, error) {
	examples := make(map[string]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-Go files and test files
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("Failed to parse %s: %v\n", path, err)
			return nil
		}

		for _, decl := range node.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				jsonExample := make(map[string]interface{})
				for _, field := range structType.Fields.List {
					if len(field.Names) == 0 {
						continue // Skip fields without a name (composed or embedded fields)
					}
					fieldName := field.Names[0].Name
					jsonTag := ""
					if field.Tag != nil {
						tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
						jsonTag = tag.Get("json")
					}
					if jsonTag == "" || jsonTag == "-" {
						jsonTag = fieldName
					} else {
						jsonTag = strings.Split(jsonTag, ",")[0]
					}

					// Add a placeholder value based on the field type
					switch fieldType := field.Type.(type) {
					case *ast.Ident:
						switch fieldType.Name {
						case "string":
							jsonExample[jsonTag] = "example string"
						case "int", "int32", "int64":
							jsonExample[jsonTag] = 123
						case "float32", "float64":
							jsonExample[jsonTag] = 123.45
						case "bool":
							jsonExample[jsonTag] = true
						default:
							jsonExample[jsonTag] = nil
						}
					case *ast.ArrayType:
						jsonExample[jsonTag] = []interface{}{"example"}
					default:
						jsonExample[jsonTag] = nil
					}
				}

				// Use the package name as a prefix for the model name
				packageName := filepath.Base(filepath.Dir(path))
				modelName := fmt.Sprintf("%s.%s", packageName, typeSpec.Name.Name)

				jsonObj := "{"
				for k, v := range jsonExample {
					jsonObj += fmt.Sprintf("\"%s\": \"%v\",", k, v)
				}
				jsonObj = strings.TrimSuffix(jsonObj, ",") + "}" // Remove trailing comma and close the JSON object
				examples[modelName] = jsonObj
			}
		}

		return nil
	})

	return examples, err
}

func ParseFiles(root string) (comments []OneComment, err error) {

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-Go files and test files
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("Failed to parse %s: %v\n", path, err)
			return nil
		}

		for _, decl := range node.Decls {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}

			if funcDecl.Doc != nil {
				comment := OneComment{
					FileName:     path,
					FunctionName: funcDecl.Name.Name,
				}
				SearchTheTag(funcDecl.Doc.List, &comment)
				comments = append(comments, comment)

			}
		}

		return err
	})

	return comments, err
}

func regexDescription(input string) string {
	re := regexp.MustCompile(`#[\w]+ (.+)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func SearchTheTag(input []*ast.Comment, comment *OneComment) string {

	for _, c := range input {
		if strings.Contains(c.Text, "#Group") {
			comment.Group = regexDescription(c.Text)
		}
		if strings.Contains(c.Text, "#GoodResponse") {
			comment.GoodResponse = regexDescription(c.Text)
		}
		if strings.Contains(c.Text, "#BadResponse") {
			comment.BadResponse = regexDescription(c.Text)
		}
		if strings.Contains(c.Text, "#Summary") {
			comment.Summary = regexDescription(c.Text)
		}
		if strings.Contains(c.Text, "#AuthNotDefault") {
			comment.AuthNotDefault = true
		}
		if strings.Contains(c.Text, "#PathParameters") {
			x := regexDescription(c.Text)
			sp := strings.Split(x, " ")
			if len(sp) == 1 {
				comment.PathParameters = append(comment.PathParameters, InputParameter{
					Name: sp[0],
				})
			} else if len(sp) == 2 {
				comment.PathParameters = append(comment.PathParameters, InputParameter{
					Name: sp[0],
					Type: sp[1],
				})
			}
		}
		if strings.Contains(c.Text, "#BodyParameters") {
			x := regexDescription(c.Text)
			sp := strings.Split(x, " ")
			if len(sp) == 1 {
				comment.BodyParameters = append(comment.BodyParameters, InputParameter{
					Name: sp[0],
				})
			} else if len(sp) == 2 {
				comment.BodyParameters = append(comment.BodyParameters, InputParameter{
					Name: sp[0],
					Type: sp[1],
				})
			}
		}
		if strings.Contains(c.Text, "#HeaderParameters") {
			x := regexDescription(c.Text)
			sp := strings.Split(x, " ")
			if len(sp) == 1 {
				comment.HeaderParameters = append(comment.HeaderParameters, InputParameter{
					Name: sp[0],
				})
			} else if len(sp) == 2 {
				comment.HeaderParameters = append(comment.HeaderParameters, InputParameter{
					Name: sp[0],
					Type: sp[1],
				})
			}
		}
		if strings.Contains(c.Text, "#QueryParameters") {
			x := regexDescription(c.Text)
			sp := strings.Split(x, " ")
			if len(sp) == 1 {
				comment.QueryParameters = append(comment.QueryParameters, InputParameter{
					Name: sp[0],
				})
			} else if len(sp) == 2 {
				comment.QueryParameters = append(comment.QueryParameters, InputParameter{
					Name: sp[0],
					Type: sp[1],
				})
			}
		}
		if strings.Contains(c.Text, "#Label") {
			x := regexDescription(c.Text)
			sp := strings.Split(x, ",")
			for _, s := range sp {
				s = strings.TrimSpace(s)
				comment.Label = append(comment.Label, s)
			}
		}

	}
	return ""
}

func UniqueOrderedGroups(endpoints []OneComment) []string {
	seen := make(map[string]bool)
	var result []string

	for _, e := range endpoints {
		if !seen[e.Group] {
			seen[e.Group] = true
			result = append(result, e.Group)
		}
	}

	sort.Strings(result)

	return result
}

type RouteInfo struct {
	Method      string
	Path        string
	HandlerFunc string
}

type Microapidoc struct {
	Doc GeneralDoc
}

func NewMicroapidoc(settings GeneralDoc) *Microapidoc {
	return &Microapidoc{Doc: settings}
}

func (c *Microapidoc) DocHAndler(ctx *gin.Context) {

	payload := ConfigPayload{
		Title:           c.Doc.Name,
		HeaderColor:     c.Doc.HeaderColor,
		BuildTag:        c.Doc.BuildTag,
		Started:         time.Now().Format(time.RFC3339),
		AuthDefaultMode: c.Doc.AuthHeaderDefaultOn,
	}
	for _, rc := range c.Doc.HighlightResponseHeaders {
		splitted := strings.Split(rc, ":")
		if len(splitted) == 3 {
			payload.ResponseHeaders = append(payload.ResponseHeaders, ResponseHeaders{
				Name:  splitted[0],
				Value: splitted[1],
				Color: splitted[2],
			})
		}
	}
	comments, err := ParseFiles(c.Doc.SearchControllersIn)
	if err != nil {
		fmt.Println("Error walking files:", err)

	}

	groupNames := UniqueOrderedGroups(comments)
	var groups []EndpointGroups
	for _, groupName := range groupNames {
		var endpoints []Endpoints
		for _, comment := range comments {
			if comment.Group == groupName {
				endpoint := Endpoints{
					Name:    comment.FunctionName,
					Summary: comment.Summary,
					Label:   comment.Label,
					Method:  "UNKNOWN",
					URL:     fmt.Sprintf("%s/%s", c.Doc.BaseUrl, "???????"),

					QueryParameters:  comment.QueryParameters,
					PathParameters:   comment.PathParameters,
					HeaderParameters: comment.HeaderParameters,
					BodyParameters:   comment.BodyParameters,
				}
				if comment.AuthNotDefault {
					endpoint.AuthHeaderOn = !c.Doc.AuthHeaderDefaultOn
				} else {
					endpoint.AuthHeaderOn = c.Doc.AuthHeaderDefaultOn
				}
				if endpoint.AuthHeaderOn {
					for _, authHeader := range c.Doc.AuthHeaderNames {
						endpoint.AuthHeaders = append(endpoint.AuthHeaders, InputParameter{
							Name: authHeader,
							Type: "string",
						})
					}
				}
				for _, route := range c.Doc.AllRoutes {
					funcDistilled := strings.ReplaceAll(route.HandlerFunc, "controllers.(*Controller).", "")
					funcDistilled = strings.ReplaceAll(funcDistilled, "-fm", "")
					if strings.Contains(funcDistilled, comment.FunctionName) {
						endpoint.Method = route.Method
						endpoint.URL = fmt.Sprintf("%s%s", c.Doc.BaseUrl, route.Path)
						break
					}
				}

				endpoints = append(endpoints, endpoint)

			}
		}
		groups = append(groups, EndpointGroups{
			Name:      groupName,
			Endpoints: endpoints,
		})

	}

	payload.Groups = groups

	// generate json examples
	jsons, err := GenerateJSONExamplesFromModels(c.Doc.SearchModelsIn)
	if err != nil {
		fmt.Println("Error generating JSON examples:", err)
	} else {
		payload.ExampleJsons = jsons
	}

	ctx.JSON(200, payload)
	return
}
