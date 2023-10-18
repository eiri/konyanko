// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "animes"
				if l := len("animes"); len(elem) >= l && elem[0:l] == "animes" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListAnimeRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleCreateAnimeRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteAnimeRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleReadAnimeRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PATCH":
							s.handleUpdateAnimeRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/episodes"
						if l := len("/episodes"); len(elem) >= l && elem[0:l] == "/episodes" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListAnimeEpisodesRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			case 'e': // Prefix: "episodes"
				if l := len("episodes"); len(elem) >= l && elem[0:l] == "episodes" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListEpisodeRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleCreateEpisodeRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteEpisodeRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleReadEpisodeRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PATCH":
							s.handleUpdateEpisodeRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "item"
							if l := len("item"); len(elem) >= l && elem[0:l] == "item" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleReadEpisodeItemRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 'r': // Prefix: "release-group"
							if l := len("release-group"); len(elem) >= l && elem[0:l] == "release-group" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleReadEpisodeReleaseGroupRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 't': // Prefix: "title"
							if l := len("title"); len(elem) >= l && elem[0:l] == "title" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleReadEpisodeTitleRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				}
			case 'i': // Prefix: "items"
				if l := len("items"); len(elem) >= l && elem[0:l] == "items" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListItemRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleCreateItemRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteItemRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleReadItemRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PATCH":
							s.handleUpdateItemRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/episode"
						if l := len("/episode"); len(elem) >= l && elem[0:l] == "/episode" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleReadItemEpisodeRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			case 'r': // Prefix: "release-groups"
				if l := len("release-groups"); len(elem) >= l && elem[0:l] == "release-groups" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListReleaseGroupRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleCreateReleaseGroupRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteReleaseGroupRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleReadReleaseGroupRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PATCH":
							s.handleUpdateReleaseGroupRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/episodes"
						if l := len("/episodes"); len(elem) >= l && elem[0:l] == "/episodes" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListReleaseGroupEpisodesRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "animes"
				if l := len("animes"); len(elem) >= l && elem[0:l] == "animes" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListAnime"
						r.summary = "List Animes"
						r.operationID = "listAnime"
						r.pathPattern = "/animes"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateAnime"
						r.summary = "Create a new Anime"
						r.operationID = "createAnime"
						r.pathPattern = "/animes"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteAnime"
							r.summary = "Deletes a Anime by ID"
							r.operationID = "deleteAnime"
							r.pathPattern = "/animes/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadAnime"
							r.summary = "Find a Anime by ID"
							r.operationID = "readAnime"
							r.pathPattern = "/animes/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateAnime"
							r.summary = "Updates a Anime"
							r.operationID = "updateAnime"
							r.pathPattern = "/animes/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/episodes"
						if l := len("/episodes"); len(elem) >= l && elem[0:l] == "/episodes" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ListAnimeEpisodes
								r.name = "ListAnimeEpisodes"
								r.summary = "List attached Episodes"
								r.operationID = "listAnimeEpisodes"
								r.pathPattern = "/animes/{id}/episodes"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'e': // Prefix: "episodes"
				if l := len("episodes"); len(elem) >= l && elem[0:l] == "episodes" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListEpisode"
						r.summary = "List Episodes"
						r.operationID = "listEpisode"
						r.pathPattern = "/episodes"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateEpisode"
						r.summary = "Create a new Episode"
						r.operationID = "createEpisode"
						r.pathPattern = "/episodes"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteEpisode"
							r.summary = "Deletes a Episode by ID"
							r.operationID = "deleteEpisode"
							r.pathPattern = "/episodes/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadEpisode"
							r.summary = "Find a Episode by ID"
							r.operationID = "readEpisode"
							r.pathPattern = "/episodes/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateEpisode"
							r.summary = "Updates a Episode"
							r.operationID = "updateEpisode"
							r.pathPattern = "/episodes/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "item"
							if l := len("item"); len(elem) >= l && elem[0:l] == "item" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ReadEpisodeItem
									r.name = "ReadEpisodeItem"
									r.summary = "Find the attached Item"
									r.operationID = "readEpisodeItem"
									r.pathPattern = "/episodes/{id}/item"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						case 'r': // Prefix: "release-group"
							if l := len("release-group"); len(elem) >= l && elem[0:l] == "release-group" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ReadEpisodeReleaseGroup
									r.name = "ReadEpisodeReleaseGroup"
									r.summary = "Find the attached ReleaseGroup"
									r.operationID = "readEpisodeReleaseGroup"
									r.pathPattern = "/episodes/{id}/release-group"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						case 't': // Prefix: "title"
							if l := len("title"); len(elem) >= l && elem[0:l] == "title" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ReadEpisodeTitle
									r.name = "ReadEpisodeTitle"
									r.summary = "Find the attached Anime"
									r.operationID = "readEpisodeTitle"
									r.pathPattern = "/episodes/{id}/title"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				}
			case 'i': // Prefix: "items"
				if l := len("items"); len(elem) >= l && elem[0:l] == "items" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListItem"
						r.summary = "List Items"
						r.operationID = "listItem"
						r.pathPattern = "/items"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateItem"
						r.summary = "Create a new Item"
						r.operationID = "createItem"
						r.pathPattern = "/items"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteItem"
							r.summary = "Deletes a Item by ID"
							r.operationID = "deleteItem"
							r.pathPattern = "/items/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadItem"
							r.summary = "Find a Item by ID"
							r.operationID = "readItem"
							r.pathPattern = "/items/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateItem"
							r.summary = "Updates a Item"
							r.operationID = "updateItem"
							r.pathPattern = "/items/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/episode"
						if l := len("/episode"); len(elem) >= l && elem[0:l] == "/episode" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ReadItemEpisode
								r.name = "ReadItemEpisode"
								r.summary = "Find the attached Episode"
								r.operationID = "readItemEpisode"
								r.pathPattern = "/items/{id}/episode"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'r': // Prefix: "release-groups"
				if l := len("release-groups"); len(elem) >= l && elem[0:l] == "release-groups" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListReleaseGroup"
						r.summary = "List ReleaseGroups"
						r.operationID = "listReleaseGroup"
						r.pathPattern = "/release-groups"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateReleaseGroup"
						r.summary = "Create a new ReleaseGroup"
						r.operationID = "createReleaseGroup"
						r.pathPattern = "/release-groups"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteReleaseGroup"
							r.summary = "Deletes a ReleaseGroup by ID"
							r.operationID = "deleteReleaseGroup"
							r.pathPattern = "/release-groups/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadReleaseGroup"
							r.summary = "Find a ReleaseGroup by ID"
							r.operationID = "readReleaseGroup"
							r.pathPattern = "/release-groups/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateReleaseGroup"
							r.summary = "Updates a ReleaseGroup"
							r.operationID = "updateReleaseGroup"
							r.pathPattern = "/release-groups/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/episodes"
						if l := len("/episodes"); len(elem) >= l && elem[0:l] == "/episodes" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ListReleaseGroupEpisodes
								r.name = "ListReleaseGroupEpisodes"
								r.summary = "List attached Episodes"
								r.operationID = "listReleaseGroupEpisodes"
								r.pathPattern = "/release-groups/{id}/episodes"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
