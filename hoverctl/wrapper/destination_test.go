package wrapper

import (
	"testing"

	"github.com/SpectoLabs/hoverfly/core/handlers/v2"
	"github.com/SpectoLabs/hoverfly/core/util"
	. "github.com/onsi/gomega"
)

func Test_GetDestination_GetsDestinationFromHoverfly(t *testing.T) {
	RegisterTestingT(t)

	hoverfly.DeleteSimulation()
	hoverfly.PutSimulation(v2.SimulationViewV3{
		v2.DataViewV3{
			RequestResponsePairs: []v2.RequestMatcherResponsePairViewV3{
				v2.RequestMatcherResponsePairViewV3{
					RequestMatcher: v2.RequestMatcherViewV3{
						Method: &v2.RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("GET"),
						},
						Path: &v2.RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("/api/v2/hoverfly/destination"),
						},
					},
					Response: v2.ResponseDetailsViewV3{
						Status: 200,
						Body:   `{"destination": "test.com"}`,
					},
				},
			},
		},
		v2.MetaView{
			SchemaVersion: "v2",
		},
	})

	destination, err := GetDestination(target)
	Expect(err).To(BeNil())

	Expect(destination).To(Equal("test.com"))
}

func Test_GetDestination_ErrorsWhen_HoverflyNotAccessible(t *testing.T) {
	RegisterTestingT(t)

	_, err := GetDestination(inaccessibleTarget)

	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Could not connect to Hoverfly at something:1234"))
}

func Test_GetDestination_ErrorsWhen_HoverflyReturnsNon200(t *testing.T) {
	RegisterTestingT(t)

	hoverfly.DeleteSimulation()
	hoverfly.PutSimulation(v2.SimulationViewV3{
		v2.DataViewV3{
			RequestResponsePairs: []v2.RequestMatcherResponsePairViewV3{
				v2.RequestMatcherResponsePairViewV3{
					RequestMatcher: v2.RequestMatcherViewV3{
						Method: &v2.RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("GET"),
						},
						Path: &v2.RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("/api/v2/hoverfly/destination"),
						},
					},
					Response: v2.ResponseDetailsViewV3{
						Status: 400,
						Body:   "{\"error\":\"test error\"}",
					},
				},
			},
		},
		v2.MetaView{
			SchemaVersion: "v2",
		},
	})

	_, err := GetDestination(target)
	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Could not retrieve destination\n\ntest error"))
}

func Test_SetDestination_SetsDestinationAndPrintsDestination(t *testing.T) {
	RegisterTestingT(t)

	hoverfly.DeleteSimulation()
	hoverfly.PutSimulation(v2.SimulationViewV3{
		v2.DataViewV3{
			RequestResponsePairs: []v2.RequestMatcherResponsePairViewV3{
				v2.RequestMatcherResponsePairViewV3{
					RequestMatcher: v2.RequestMatcherViewV3{
						Method: &v2.RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("PUT"),
						},
						Path: &v2.RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("/api/v2/hoverfly/destination"),
						},
					},
					Response: v2.ResponseDetailsViewV3{
						Status: 200,
						Body:   `{"destination": "new.com"}`,
					},
				},
			},
		},
		v2.MetaView{
			SchemaVersion: "v2",
		},
	})

	destination, err := SetDestination(target, "new.com")
	Expect(err).To(BeNil())

	Expect(destination).To(Equal("new.com"))
}

func Test_SetDestination_ErrorsWhen_HoverflyNotAccessible(t *testing.T) {
	RegisterTestingT(t)

	_, err := SetDestination(inaccessibleTarget, "something")

	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Could not connect to Hoverfly at something:1234"))
}

func Test_SetDestination_ErrorsWhen_HoverflyReturnsNon200(t *testing.T) {
	RegisterTestingT(t)

	hoverfly.DeleteSimulation()
	hoverfly.PutSimulation(v2.SimulationViewV3{
		v2.DataViewV3{
			RequestResponsePairs: []v2.RequestMatcherResponsePairViewV3{
				v2.RequestMatcherResponsePairViewV3{
					RequestMatcher: v2.RequestMatcherViewV3{
						Method: &v2.RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("PUT"),
						},
						Path: &v2.RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("/api/v2/hoverfly/destination"),
						},
					},
					Response: v2.ResponseDetailsViewV3{
						Status: 400,
						Body:   "{\"error\":\"test error\"}",
					},
				},
			},
		},
		v2.MetaView{
			SchemaVersion: "v2",
		},
	})

	_, err := SetDestination(target, "new.com")
	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Could not set destination\n\ntest error"))
}
