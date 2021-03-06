package ccv3_test

import (
	"fmt"
	"net/http"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	. "code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Route", func() {
	var client *Client

	BeforeEach(func() {
		client, _ = NewTestClient()
	})

	Describe("CreateRoute", func() {
		var (
			route      Route
			warnings   Warnings
			executeErr error
			spaceGUID  string
			domainGUID string
			host       string
			path       string
			ccv3Route  Route
		)

		BeforeEach(func() {
			host = ""
			path = ""
		})

		JustBeforeEach(func() {
			spaceGUID = "space-guid"
			domainGUID = "domain-guid"
			ccv3Route = Route{SpaceGUID: spaceGUID, DomainGUID: domainGUID, Host: host, Path: path}
			route, warnings, executeErr = client.CreateRoute(ccv3Route)
		})

		When("the request succeeds", func() {
			When("no additional flags", func() {
				BeforeEach(func() {
					host = ""
					response := `{
  "guid": "some-route-guid",
  "relationships": {
    "space": {
	  "data": { "guid": "space-guid" }
    },
    "domain": {
	  "data": { "guid": "domain-guid" }
    }
  },
	"host": ""
}`

					expectedBody := `{
  "relationships": {
  	"space": {
      "data": { "guid": "space-guid" }
    },
    "domain": {
	  "data": { "guid": "domain-guid" }
    }
  }
}`

					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodPost, "/v3/routes"),
							VerifyJSON(expectedBody),
							RespondWith(http.StatusCreated, response, http.Header{"X-Cf-Warnings": {"warning-1"}}),
						),
					)
				})

				It("returns the given route and all warnings", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(warnings).To(ConsistOf("warning-1"))

					Expect(route).To(Equal(Route{
						GUID:       "some-route-guid",
						SpaceGUID:  "space-guid",
						DomainGUID: "domain-guid",
					}))
				})
			})

			When("hostname is passed in", func() {

				BeforeEach(func() {
					host = "cheesecake"
					response := `{
  "guid": "some-route-guid",
  "relationships": {
    "space": {
			"data": { "guid": "space-guid" }
    },
    "domain": {
			"data": { "guid": "domain-guid" }
    }
  },
	"host": "cheesecake"
}`

					expectedBody := `{
  "relationships": {
  	"space": {
      "data": { "guid": "space-guid" }
    },
    "domain": {
			"data": { "guid": "domain-guid" }
    }
  },
	"host": "cheesecake"
}`

					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodPost, "/v3/routes"),
							VerifyJSON(expectedBody),
							RespondWith(http.StatusCreated, response, http.Header{"X-Cf-Warnings": {"warning-1"}}),
						),
					)
				})

				It("returns the given route and all warnings", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(warnings).To(ConsistOf("warning-1"))

					Expect(route).To(Equal(Route{
						GUID:       "some-route-guid",
						SpaceGUID:  "space-guid",
						DomainGUID: "domain-guid",
						Host:       "cheesecake",
					}))
				})
			})

			When("path is passed in", func() {
				BeforeEach(func() {
					path = "lion"

					response := `{
	"guid": "this-route-guid",
	"relationships": {
		"space": {
			"data": {
				"guid": "space-guid"
			}
		},
		"domain": {
			"data": {
				"guid": "domain-guid"
			}
		}
	},
	"path": "lion"
}`
					expectedRequestBody := `{
	"relationships": {
		"space": {
			"data": {
				"guid": "space-guid"
			}
		},
		"domain": {
			"data": {
				"guid": "domain-guid"
			}
		}
	},
	"path": "lion"
}`

					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodPost, "/v3/routes"),
							VerifyJSON(expectedRequestBody),
							RespondWith(http.StatusCreated, response, http.Header{"X-Cf-Warnings": {"warning-1"}}),
						),
					)
				})
				When("the request succeeds", func() {
					It("returns the given route and all warnings", func() {
						Expect(executeErr).ToNot(HaveOccurred())
						Expect(warnings).To(ConsistOf("warning-1"))

						Expect(route).To(Equal(Route{
							GUID:       "this-route-guid",
							SpaceGUID:  "space-guid",
							DomainGUID: "domain-guid",
							Path:       "lion",
						}))
					})
				})
			})

		})

		When("the cloud controller returns errors and warnings", func() {
			BeforeEach(func() {
				response := `{
  "errors": [
    {
      "code": 10008,
      "detail": "The request is semantically invalid: command presence",
      "title": "CF-UnprocessableEntity"
    },
		{
      "code": 10010,
      "detail": "Isolation segment not found",
      "title": "CF-ResourceNotFound"
    }
  ]
}`
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodPost, "/v3/routes"),
						RespondWith(http.StatusTeapot, response, http.Header{"X-Cf-Warnings": {"this is a warning"}}),
					),
				)
			})

			It("returns the error and all warnings", func() {
				Expect(executeErr).To(MatchError(ccerror.MultiError{
					ResponseCode: http.StatusTeapot,
					Errors: []ccerror.V3Error{
						{
							Code:   10008,
							Detail: "The request is semantically invalid: command presence",
							Title:  "CF-UnprocessableEntity",
						},
						{
							Code:   10010,
							Detail: "Isolation segment not found",
							Title:  "CF-ResourceNotFound",
						},
					},
				}))
				Expect(warnings).To(ConsistOf("this is a warning"))
			})
		})
	})

	Describe("GetRoutes", func() {
		var (
			query      Query
			routes     []Route
			warnings   Warnings
			executeErr error
		)

		JustBeforeEach(func() {
			routes, warnings, executeErr = client.GetRoutes(query)
		})

		When("the request succeeds", func() {
			var (
				response1 string
				response2 string
			)

			BeforeEach(func() {
				response1 = fmt.Sprintf(`
				{
					"pagination": {
						"next": {
							"href": "%s/v3/routes?page=2"
						}
					},
					"resources": [
						{
							"guid": "route-1-guid"
						},
						{
							"guid": "route-2-guid"
						}
					]
				}`, server.URL())

				response2 = `
				{
					"pagination": {
						"next": null
					},
					"resources": [
						{
							"guid": "route-3-guid"
						}
					]
				}`
			})

			When("not passing any filters", func() {
				BeforeEach(func() {
					query = Query{}

					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodGet, "/v3/routes"),
							RespondWith(http.StatusOK, response1, http.Header{"X-Cf-Warnings": {"warning-1"}}),
						),
					)
					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodGet, "/v3/routes", "page=2"),
							RespondWith(http.StatusOK, response2, http.Header{"X-Cf-Warnings": {"warning-2"}}),
						),
					)
				})

				It("returns the given route and all warnings", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))

					Expect(routes).To(Equal([]Route{
						Route{
							GUID: "route-1-guid",
						},
						Route{
							GUID: "route-2-guid",
						},
						Route{
							GUID: "route-3-guid",
						},
					}))
				})
			})

			When("passing in a query", func() {
				BeforeEach(func() {
					query = Query{Key: "space_guids", Values: []string{"guid1", "guid2"}}

					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodGet, "/v3/routes", "space_guids=guid1,guid2"),
							RespondWith(http.StatusOK, response1, http.Header{"X-Cf-Warnings": {"warning-1"}}),
						),
					)
					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodGet, "/v3/routes", "page=2", "space_guids=guid1,guid2"),
							RespondWith(http.StatusOK, response2, http.Header{"X-Cf-Warnings": {"warning-2"}}),
						),
					)
				})

				It("passes query params", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))

					Expect(routes).To(Equal([]Route{
						Route{
							GUID: "route-1-guid",
						},
						Route{
							GUID: "route-2-guid",
						},
						Route{
							GUID: "route-3-guid",
						},
					}))
				})
			})
		})
	})
})
