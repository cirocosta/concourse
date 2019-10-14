package tracing_test

import (
	"context"

	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/dbfakes"
	"github.com/concourse/concourse/atc/tracing"
	"github.com/concourse/concourse/atc/tracing/tracingfakes"
	"go.opentelemetry.io/api/key"
	"go.opentelemetry.io/api/trace"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tracer", func() {

	var (
		tracer tracing.Tracer
		span   trace.Span
	)

	BeforeEach(func() {
		fakeSpan := new(tracingfakes.FakeSpan)
		fakeTracer := new(tracingfakes.FakeTracer)

		fakeTracer.StartReturns(
			context.Background(),
			fakeSpan,
		)

		tracer = tracing.Tracer{fakeTracer}
	})

	Describe("StartSpan", func() {

		Context("with empty context", func() {

			It("starts root span", func() {

			})

		})

		Context("with span-based context", func() {

			It("starts child span", func() {

			})

		})

	})

	Describe("BuildRootSpan", func() {

		var (
			build db.Build
		)

		Context("without a build (programming error)", func() {
			It("panics", func() {
				Expect(func() {
					tracer.BuildRootSpan(build)
				}).To(Panic())
			})
		})

		Context("with a build", func() {

			BeforeEach(func() {
				fakeBuild := new(dbfakes.FakeBuild)

				fakeBuild.NameReturns("build-name")
				fakeBuild.JobNameReturns("job-name")
				fakeBuild.PipelineNameReturns("pipeline-name")
				fakeBuild.TeamNameReturns("team-name")

				build = fakeBuild
			})

			JustBeforeEach(func() {
				span = tracer.BuildRootSpan(build)
			})

			It("creates a span", func() {
				Expect(span).ToNot(BeNil())
			})

			It("has build attributes set", func() {
				fakeSpan := span.(*tracingfakes.FakeSpan)
				kvs := fakeSpan.SetAttributesArgsForCall(0)

				Expect(kvs).ToNot(BeEmpty())
				Expect(kvs).To(ConsistOf(
					key.New("team").String("team-name"),
					key.New("pipeline").String("pipeline-name"),
					key.New("job").String("job-name"),
					key.New("build").String("build-name"),
				))
			})
		})
	})

})
