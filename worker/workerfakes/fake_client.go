// This file was generated by counterfeiter
package workerfakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/worker"
)

type FakeClient struct {
	FindOrCreateBuildContainerStub        func(lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, map[string]string) (worker.Container, error)
	findOrCreateBuildContainerMutex       sync.RWMutex
	findOrCreateBuildContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 worker.Identifier
		arg5 worker.Metadata
		arg6 worker.ContainerSpec
		arg7 atc.ResourceTypes
		arg8 map[string]string
	}
	findOrCreateBuildContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	FindOrCreateResourceGetContainerStub        func(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, outputPaths map[string]string, resourceType string, version atc.Version, source atc.Source, params atc.Params) (worker.Container, error)
	findOrCreateResourceGetContainerMutex       sync.RWMutex
	findOrCreateResourceGetContainerArgsForCall []struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		outputPaths   map[string]string
		resourceType  string
		version       atc.Version
		source        atc.Source
		params        atc.Params
	}
	findOrCreateResourceGetContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	FindOrCreateResourceCheckContainerStub        func(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, resourceType string, source atc.Source) (worker.Container, error)
	findOrCreateResourceCheckContainerMutex       sync.RWMutex
	findOrCreateResourceCheckContainerArgsForCall []struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		resourceType  string
		source        atc.Source
	}
	findOrCreateResourceCheckContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	FindOrCreateResourceTypeCheckContainerStub        func(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, resourceType string, source atc.Source) (worker.Container, error)
	findOrCreateResourceTypeCheckContainerMutex       sync.RWMutex
	findOrCreateResourceTypeCheckContainerArgsForCall []struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		resourceType  string
		source        atc.Source
	}
	findOrCreateResourceTypeCheckContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	FindOrCreateContainerForIdentifierStub        func(logger lager.Logger, id worker.Identifier, metadata worker.Metadata, containerSpec worker.ContainerSpec, resourceTypes atc.ResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate, resourceSources map[string]worker.ArtifactSource) (worker.Container, []string, error)
	findOrCreateContainerForIdentifierMutex       sync.RWMutex
	findOrCreateContainerForIdentifierArgsForCall []struct {
		logger                lager.Logger
		id                    worker.Identifier
		metadata              worker.Metadata
		containerSpec         worker.ContainerSpec
		resourceTypes         atc.ResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
		resourceSources       map[string]worker.ArtifactSource
	}
	findOrCreateContainerForIdentifierReturns struct {
		result1 worker.Container
		result2 []string
		result3 error
	}
	FindOrCreateVolumeForResourceCacheStub        func(logger lager.Logger, vs worker.VolumeSpec, resourceCache *dbng.UsedResourceCache) (worker.Volume, error)
	findOrCreateVolumeForResourceCacheMutex       sync.RWMutex
	findOrCreateVolumeForResourceCacheArgsForCall []struct {
		logger        lager.Logger
		vs            worker.VolumeSpec
		resourceCache *dbng.UsedResourceCache
	}
	findOrCreateVolumeForResourceCacheReturns struct {
		result1 worker.Volume
		result2 error
	}
	FindInitializedVolumeForResourceCacheStub        func(logger lager.Logger, resourceCache *dbng.UsedResourceCache) (worker.Volume, bool, error)
	findInitializedVolumeForResourceCacheMutex       sync.RWMutex
	findInitializedVolumeForResourceCacheArgsForCall []struct {
		logger        lager.Logger
		resourceCache *dbng.UsedResourceCache
	}
	findInitializedVolumeForResourceCacheReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	FindContainerForIdentifierStub        func(lager.Logger, worker.Identifier) (worker.Container, bool, error)
	findContainerForIdentifierMutex       sync.RWMutex
	findContainerForIdentifierArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Identifier
	}
	findContainerForIdentifierReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	LookupContainerStub        func(lager.Logger, string) (worker.Container, bool, error)
	lookupContainerMutex       sync.RWMutex
	lookupContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupContainerReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	ValidateResourceCheckVersionStub        func(container db.SavedContainer) (bool, error)
	validateResourceCheckVersionMutex       sync.RWMutex
	validateResourceCheckVersionArgsForCall []struct {
		container db.SavedContainer
	}
	validateResourceCheckVersionReturns struct {
		result1 bool
		result2 error
	}
	FindResourceTypeByPathStub        func(path string) (atc.WorkerResourceType, bool)
	findResourceTypeByPathMutex       sync.RWMutex
	findResourceTypeByPathArgsForCall []struct {
		path string
	}
	findResourceTypeByPathReturns struct {
		result1 atc.WorkerResourceType
		result2 bool
	}
	LookupVolumeStub        func(lager.Logger, string) (worker.Volume, bool, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	SatisfyingStub        func(worker.WorkerSpec, atc.ResourceTypes) (worker.Worker, error)
	satisfyingMutex       sync.RWMutex
	satisfyingArgsForCall []struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}
	satisfyingReturns struct {
		result1 worker.Worker
		result2 error
	}
	AllSatisfyingStub        func(worker.WorkerSpec, atc.ResourceTypes) ([]worker.Worker, error)
	allSatisfyingMutex       sync.RWMutex
	allSatisfyingArgsForCall []struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}
	allSatisfyingReturns struct {
		result1 []worker.Worker
		result2 error
	}
	RunningWorkersStub        func() ([]worker.Worker, error)
	runningWorkersMutex       sync.RWMutex
	runningWorkersArgsForCall []struct{}
	runningWorkersReturns     struct {
		result1 []worker.Worker
		result2 error
	}
	GetWorkerStub        func(workerName string) (worker.Worker, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		workerName string
	}
	getWorkerReturns struct {
		result1 worker.Worker
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) FindOrCreateBuildContainer(arg1 lager.Logger, arg2 <-chan os.Signal, arg3 worker.ImageFetchingDelegate, arg4 worker.Identifier, arg5 worker.Metadata, arg6 worker.ContainerSpec, arg7 atc.ResourceTypes, arg8 map[string]string) (worker.Container, error) {
	fake.findOrCreateBuildContainerMutex.Lock()
	fake.findOrCreateBuildContainerArgsForCall = append(fake.findOrCreateBuildContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 worker.Identifier
		arg5 worker.Metadata
		arg6 worker.ContainerSpec
		arg7 atc.ResourceTypes
		arg8 map[string]string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.recordInvocation("FindOrCreateBuildContainer", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.findOrCreateBuildContainerMutex.Unlock()
	if fake.FindOrCreateBuildContainerStub != nil {
		return fake.FindOrCreateBuildContainerStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	} else {
		return fake.findOrCreateBuildContainerReturns.result1, fake.findOrCreateBuildContainerReturns.result2
	}
}

func (fake *FakeClient) FindOrCreateBuildContainerCallCount() int {
	fake.findOrCreateBuildContainerMutex.RLock()
	defer fake.findOrCreateBuildContainerMutex.RUnlock()
	return len(fake.findOrCreateBuildContainerArgsForCall)
}

func (fake *FakeClient) FindOrCreateBuildContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, map[string]string) {
	fake.findOrCreateBuildContainerMutex.RLock()
	defer fake.findOrCreateBuildContainerMutex.RUnlock()
	return fake.findOrCreateBuildContainerArgsForCall[i].arg1, fake.findOrCreateBuildContainerArgsForCall[i].arg2, fake.findOrCreateBuildContainerArgsForCall[i].arg3, fake.findOrCreateBuildContainerArgsForCall[i].arg4, fake.findOrCreateBuildContainerArgsForCall[i].arg5, fake.findOrCreateBuildContainerArgsForCall[i].arg6, fake.findOrCreateBuildContainerArgsForCall[i].arg7, fake.findOrCreateBuildContainerArgsForCall[i].arg8
}

func (fake *FakeClient) FindOrCreateBuildContainerReturns(result1 worker.Container, result2 error) {
	fake.FindOrCreateBuildContainerStub = nil
	fake.findOrCreateBuildContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindOrCreateResourceGetContainer(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, outputPaths map[string]string, resourceType string, version atc.Version, source atc.Source, params atc.Params) (worker.Container, error) {
	fake.findOrCreateResourceGetContainerMutex.Lock()
	fake.findOrCreateResourceGetContainerArgsForCall = append(fake.findOrCreateResourceGetContainerArgsForCall, struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		outputPaths   map[string]string
		resourceType  string
		version       atc.Version
		source        atc.Source
		params        atc.Params
	}{logger, cancel, delegate, id, metadata, spec, resourceTypes, outputPaths, resourceType, version, source, params})
	fake.recordInvocation("FindOrCreateResourceGetContainer", []interface{}{logger, cancel, delegate, id, metadata, spec, resourceTypes, outputPaths, resourceType, version, source, params})
	fake.findOrCreateResourceGetContainerMutex.Unlock()
	if fake.FindOrCreateResourceGetContainerStub != nil {
		return fake.FindOrCreateResourceGetContainerStub(logger, cancel, delegate, id, metadata, spec, resourceTypes, outputPaths, resourceType, version, source, params)
	} else {
		return fake.findOrCreateResourceGetContainerReturns.result1, fake.findOrCreateResourceGetContainerReturns.result2
	}
}

func (fake *FakeClient) FindOrCreateResourceGetContainerCallCount() int {
	fake.findOrCreateResourceGetContainerMutex.RLock()
	defer fake.findOrCreateResourceGetContainerMutex.RUnlock()
	return len(fake.findOrCreateResourceGetContainerArgsForCall)
}

func (fake *FakeClient) FindOrCreateResourceGetContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, map[string]string, string, atc.Version, atc.Source, atc.Params) {
	fake.findOrCreateResourceGetContainerMutex.RLock()
	defer fake.findOrCreateResourceGetContainerMutex.RUnlock()
	return fake.findOrCreateResourceGetContainerArgsForCall[i].logger, fake.findOrCreateResourceGetContainerArgsForCall[i].cancel, fake.findOrCreateResourceGetContainerArgsForCall[i].delegate, fake.findOrCreateResourceGetContainerArgsForCall[i].id, fake.findOrCreateResourceGetContainerArgsForCall[i].metadata, fake.findOrCreateResourceGetContainerArgsForCall[i].spec, fake.findOrCreateResourceGetContainerArgsForCall[i].resourceTypes, fake.findOrCreateResourceGetContainerArgsForCall[i].outputPaths, fake.findOrCreateResourceGetContainerArgsForCall[i].resourceType, fake.findOrCreateResourceGetContainerArgsForCall[i].version, fake.findOrCreateResourceGetContainerArgsForCall[i].source, fake.findOrCreateResourceGetContainerArgsForCall[i].params
}

func (fake *FakeClient) FindOrCreateResourceGetContainerReturns(result1 worker.Container, result2 error) {
	fake.FindOrCreateResourceGetContainerStub = nil
	fake.findOrCreateResourceGetContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindOrCreateResourceCheckContainer(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, resourceType string, source atc.Source) (worker.Container, error) {
	fake.findOrCreateResourceCheckContainerMutex.Lock()
	fake.findOrCreateResourceCheckContainerArgsForCall = append(fake.findOrCreateResourceCheckContainerArgsForCall, struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		resourceType  string
		source        atc.Source
	}{logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source})
	fake.recordInvocation("FindOrCreateResourceCheckContainer", []interface{}{logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source})
	fake.findOrCreateResourceCheckContainerMutex.Unlock()
	if fake.FindOrCreateResourceCheckContainerStub != nil {
		return fake.FindOrCreateResourceCheckContainerStub(logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source)
	} else {
		return fake.findOrCreateResourceCheckContainerReturns.result1, fake.findOrCreateResourceCheckContainerReturns.result2
	}
}

func (fake *FakeClient) FindOrCreateResourceCheckContainerCallCount() int {
	fake.findOrCreateResourceCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceCheckContainerMutex.RUnlock()
	return len(fake.findOrCreateResourceCheckContainerArgsForCall)
}

func (fake *FakeClient) FindOrCreateResourceCheckContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, string, atc.Source) {
	fake.findOrCreateResourceCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceCheckContainerMutex.RUnlock()
	return fake.findOrCreateResourceCheckContainerArgsForCall[i].logger, fake.findOrCreateResourceCheckContainerArgsForCall[i].cancel, fake.findOrCreateResourceCheckContainerArgsForCall[i].delegate, fake.findOrCreateResourceCheckContainerArgsForCall[i].id, fake.findOrCreateResourceCheckContainerArgsForCall[i].metadata, fake.findOrCreateResourceCheckContainerArgsForCall[i].spec, fake.findOrCreateResourceCheckContainerArgsForCall[i].resourceTypes, fake.findOrCreateResourceCheckContainerArgsForCall[i].resourceType, fake.findOrCreateResourceCheckContainerArgsForCall[i].source
}

func (fake *FakeClient) FindOrCreateResourceCheckContainerReturns(result1 worker.Container, result2 error) {
	fake.FindOrCreateResourceCheckContainerStub = nil
	fake.findOrCreateResourceCheckContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindOrCreateResourceTypeCheckContainer(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, resourceType string, source atc.Source) (worker.Container, error) {
	fake.findOrCreateResourceTypeCheckContainerMutex.Lock()
	fake.findOrCreateResourceTypeCheckContainerArgsForCall = append(fake.findOrCreateResourceTypeCheckContainerArgsForCall, struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		resourceType  string
		source        atc.Source
	}{logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source})
	fake.recordInvocation("FindOrCreateResourceTypeCheckContainer", []interface{}{logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source})
	fake.findOrCreateResourceTypeCheckContainerMutex.Unlock()
	if fake.FindOrCreateResourceTypeCheckContainerStub != nil {
		return fake.FindOrCreateResourceTypeCheckContainerStub(logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source)
	} else {
		return fake.findOrCreateResourceTypeCheckContainerReturns.result1, fake.findOrCreateResourceTypeCheckContainerReturns.result2
	}
}

func (fake *FakeClient) FindOrCreateResourceTypeCheckContainerCallCount() int {
	fake.findOrCreateResourceTypeCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceTypeCheckContainerMutex.RUnlock()
	return len(fake.findOrCreateResourceTypeCheckContainerArgsForCall)
}

func (fake *FakeClient) FindOrCreateResourceTypeCheckContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, string, atc.Source) {
	fake.findOrCreateResourceTypeCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceTypeCheckContainerMutex.RUnlock()
	return fake.findOrCreateResourceTypeCheckContainerArgsForCall[i].logger, fake.findOrCreateResourceTypeCheckContainerArgsForCall[i].cancel, fake.findOrCreateResourceTypeCheckContainerArgsForCall[i].delegate, fake.findOrCreateResourceTypeCheckContainerArgsForCall[i].id, fake.findOrCreateResourceTypeCheckContainerArgsForCall[i].metadata, fake.findOrCreateResourceTypeCheckContainerArgsForCall[i].spec, fake.findOrCreateResourceTypeCheckContainerArgsForCall[i].resourceTypes, fake.findOrCreateResourceTypeCheckContainerArgsForCall[i].resourceType, fake.findOrCreateResourceTypeCheckContainerArgsForCall[i].source
}

func (fake *FakeClient) FindOrCreateResourceTypeCheckContainerReturns(result1 worker.Container, result2 error) {
	fake.FindOrCreateResourceTypeCheckContainerStub = nil
	fake.findOrCreateResourceTypeCheckContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindOrCreateContainerForIdentifier(logger lager.Logger, id worker.Identifier, metadata worker.Metadata, containerSpec worker.ContainerSpec, resourceTypes atc.ResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate, resourceSources map[string]worker.ArtifactSource) (worker.Container, []string, error) {
	fake.findOrCreateContainerForIdentifierMutex.Lock()
	fake.findOrCreateContainerForIdentifierArgsForCall = append(fake.findOrCreateContainerForIdentifierArgsForCall, struct {
		logger                lager.Logger
		id                    worker.Identifier
		metadata              worker.Metadata
		containerSpec         worker.ContainerSpec
		resourceTypes         atc.ResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
		resourceSources       map[string]worker.ArtifactSource
	}{logger, id, metadata, containerSpec, resourceTypes, imageFetchingDelegate, resourceSources})
	fake.recordInvocation("FindOrCreateContainerForIdentifier", []interface{}{logger, id, metadata, containerSpec, resourceTypes, imageFetchingDelegate, resourceSources})
	fake.findOrCreateContainerForIdentifierMutex.Unlock()
	if fake.FindOrCreateContainerForIdentifierStub != nil {
		return fake.FindOrCreateContainerForIdentifierStub(logger, id, metadata, containerSpec, resourceTypes, imageFetchingDelegate, resourceSources)
	} else {
		return fake.findOrCreateContainerForIdentifierReturns.result1, fake.findOrCreateContainerForIdentifierReturns.result2, fake.findOrCreateContainerForIdentifierReturns.result3
	}
}

func (fake *FakeClient) FindOrCreateContainerForIdentifierCallCount() int {
	fake.findOrCreateContainerForIdentifierMutex.RLock()
	defer fake.findOrCreateContainerForIdentifierMutex.RUnlock()
	return len(fake.findOrCreateContainerForIdentifierArgsForCall)
}

func (fake *FakeClient) FindOrCreateContainerForIdentifierArgsForCall(i int) (lager.Logger, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, worker.ImageFetchingDelegate, map[string]worker.ArtifactSource) {
	fake.findOrCreateContainerForIdentifierMutex.RLock()
	defer fake.findOrCreateContainerForIdentifierMutex.RUnlock()
	return fake.findOrCreateContainerForIdentifierArgsForCall[i].logger, fake.findOrCreateContainerForIdentifierArgsForCall[i].id, fake.findOrCreateContainerForIdentifierArgsForCall[i].metadata, fake.findOrCreateContainerForIdentifierArgsForCall[i].containerSpec, fake.findOrCreateContainerForIdentifierArgsForCall[i].resourceTypes, fake.findOrCreateContainerForIdentifierArgsForCall[i].imageFetchingDelegate, fake.findOrCreateContainerForIdentifierArgsForCall[i].resourceSources
}

func (fake *FakeClient) FindOrCreateContainerForIdentifierReturns(result1 worker.Container, result2 []string, result3 error) {
	fake.FindOrCreateContainerForIdentifierStub = nil
	fake.findOrCreateContainerForIdentifierReturns = struct {
		result1 worker.Container
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindOrCreateVolumeForResourceCache(logger lager.Logger, vs worker.VolumeSpec, resourceCache *dbng.UsedResourceCache) (worker.Volume, error) {
	fake.findOrCreateVolumeForResourceCacheMutex.Lock()
	fake.findOrCreateVolumeForResourceCacheArgsForCall = append(fake.findOrCreateVolumeForResourceCacheArgsForCall, struct {
		logger        lager.Logger
		vs            worker.VolumeSpec
		resourceCache *dbng.UsedResourceCache
	}{logger, vs, resourceCache})
	fake.recordInvocation("FindOrCreateVolumeForResourceCache", []interface{}{logger, vs, resourceCache})
	fake.findOrCreateVolumeForResourceCacheMutex.Unlock()
	if fake.FindOrCreateVolumeForResourceCacheStub != nil {
		return fake.FindOrCreateVolumeForResourceCacheStub(logger, vs, resourceCache)
	} else {
		return fake.findOrCreateVolumeForResourceCacheReturns.result1, fake.findOrCreateVolumeForResourceCacheReturns.result2
	}
}

func (fake *FakeClient) FindOrCreateVolumeForResourceCacheCallCount() int {
	fake.findOrCreateVolumeForResourceCacheMutex.RLock()
	defer fake.findOrCreateVolumeForResourceCacheMutex.RUnlock()
	return len(fake.findOrCreateVolumeForResourceCacheArgsForCall)
}

func (fake *FakeClient) FindOrCreateVolumeForResourceCacheArgsForCall(i int) (lager.Logger, worker.VolumeSpec, *dbng.UsedResourceCache) {
	fake.findOrCreateVolumeForResourceCacheMutex.RLock()
	defer fake.findOrCreateVolumeForResourceCacheMutex.RUnlock()
	return fake.findOrCreateVolumeForResourceCacheArgsForCall[i].logger, fake.findOrCreateVolumeForResourceCacheArgsForCall[i].vs, fake.findOrCreateVolumeForResourceCacheArgsForCall[i].resourceCache
}

func (fake *FakeClient) FindOrCreateVolumeForResourceCacheReturns(result1 worker.Volume, result2 error) {
	fake.FindOrCreateVolumeForResourceCacheStub = nil
	fake.findOrCreateVolumeForResourceCacheReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindInitializedVolumeForResourceCache(logger lager.Logger, resourceCache *dbng.UsedResourceCache) (worker.Volume, bool, error) {
	fake.findInitializedVolumeForResourceCacheMutex.Lock()
	fake.findInitializedVolumeForResourceCacheArgsForCall = append(fake.findInitializedVolumeForResourceCacheArgsForCall, struct {
		logger        lager.Logger
		resourceCache *dbng.UsedResourceCache
	}{logger, resourceCache})
	fake.recordInvocation("FindInitializedVolumeForResourceCache", []interface{}{logger, resourceCache})
	fake.findInitializedVolumeForResourceCacheMutex.Unlock()
	if fake.FindInitializedVolumeForResourceCacheStub != nil {
		return fake.FindInitializedVolumeForResourceCacheStub(logger, resourceCache)
	} else {
		return fake.findInitializedVolumeForResourceCacheReturns.result1, fake.findInitializedVolumeForResourceCacheReturns.result2, fake.findInitializedVolumeForResourceCacheReturns.result3
	}
}

func (fake *FakeClient) FindInitializedVolumeForResourceCacheCallCount() int {
	fake.findInitializedVolumeForResourceCacheMutex.RLock()
	defer fake.findInitializedVolumeForResourceCacheMutex.RUnlock()
	return len(fake.findInitializedVolumeForResourceCacheArgsForCall)
}

func (fake *FakeClient) FindInitializedVolumeForResourceCacheArgsForCall(i int) (lager.Logger, *dbng.UsedResourceCache) {
	fake.findInitializedVolumeForResourceCacheMutex.RLock()
	defer fake.findInitializedVolumeForResourceCacheMutex.RUnlock()
	return fake.findInitializedVolumeForResourceCacheArgsForCall[i].logger, fake.findInitializedVolumeForResourceCacheArgsForCall[i].resourceCache
}

func (fake *FakeClient) FindInitializedVolumeForResourceCacheReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.FindInitializedVolumeForResourceCacheStub = nil
	fake.findInitializedVolumeForResourceCacheReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindContainerForIdentifier(arg1 lager.Logger, arg2 worker.Identifier) (worker.Container, bool, error) {
	fake.findContainerForIdentifierMutex.Lock()
	fake.findContainerForIdentifierArgsForCall = append(fake.findContainerForIdentifierArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Identifier
	}{arg1, arg2})
	fake.recordInvocation("FindContainerForIdentifier", []interface{}{arg1, arg2})
	fake.findContainerForIdentifierMutex.Unlock()
	if fake.FindContainerForIdentifierStub != nil {
		return fake.FindContainerForIdentifierStub(arg1, arg2)
	} else {
		return fake.findContainerForIdentifierReturns.result1, fake.findContainerForIdentifierReturns.result2, fake.findContainerForIdentifierReturns.result3
	}
}

func (fake *FakeClient) FindContainerForIdentifierCallCount() int {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return len(fake.findContainerForIdentifierArgsForCall)
}

func (fake *FakeClient) FindContainerForIdentifierArgsForCall(i int) (lager.Logger, worker.Identifier) {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return fake.findContainerForIdentifierArgsForCall[i].arg1, fake.findContainerForIdentifierArgsForCall[i].arg2
}

func (fake *FakeClient) FindContainerForIdentifierReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.FindContainerForIdentifierStub = nil
	fake.findContainerForIdentifierReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) LookupContainer(arg1 lager.Logger, arg2 string) (worker.Container, bool, error) {
	fake.lookupContainerMutex.Lock()
	fake.lookupContainerArgsForCall = append(fake.lookupContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupContainer", []interface{}{arg1, arg2})
	fake.lookupContainerMutex.Unlock()
	if fake.LookupContainerStub != nil {
		return fake.LookupContainerStub(arg1, arg2)
	} else {
		return fake.lookupContainerReturns.result1, fake.lookupContainerReturns.result2, fake.lookupContainerReturns.result3
	}
}

func (fake *FakeClient) LookupContainerCallCount() int {
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	return len(fake.lookupContainerArgsForCall)
}

func (fake *FakeClient) LookupContainerArgsForCall(i int) (lager.Logger, string) {
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	return fake.lookupContainerArgsForCall[i].arg1, fake.lookupContainerArgsForCall[i].arg2
}

func (fake *FakeClient) LookupContainerReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.LookupContainerStub = nil
	fake.lookupContainerReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ValidateResourceCheckVersion(container db.SavedContainer) (bool, error) {
	fake.validateResourceCheckVersionMutex.Lock()
	fake.validateResourceCheckVersionArgsForCall = append(fake.validateResourceCheckVersionArgsForCall, struct {
		container db.SavedContainer
	}{container})
	fake.recordInvocation("ValidateResourceCheckVersion", []interface{}{container})
	fake.validateResourceCheckVersionMutex.Unlock()
	if fake.ValidateResourceCheckVersionStub != nil {
		return fake.ValidateResourceCheckVersionStub(container)
	} else {
		return fake.validateResourceCheckVersionReturns.result1, fake.validateResourceCheckVersionReturns.result2
	}
}

func (fake *FakeClient) ValidateResourceCheckVersionCallCount() int {
	fake.validateResourceCheckVersionMutex.RLock()
	defer fake.validateResourceCheckVersionMutex.RUnlock()
	return len(fake.validateResourceCheckVersionArgsForCall)
}

func (fake *FakeClient) ValidateResourceCheckVersionArgsForCall(i int) db.SavedContainer {
	fake.validateResourceCheckVersionMutex.RLock()
	defer fake.validateResourceCheckVersionMutex.RUnlock()
	return fake.validateResourceCheckVersionArgsForCall[i].container
}

func (fake *FakeClient) ValidateResourceCheckVersionReturns(result1 bool, result2 error) {
	fake.ValidateResourceCheckVersionStub = nil
	fake.validateResourceCheckVersionReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindResourceTypeByPath(path string) (atc.WorkerResourceType, bool) {
	fake.findResourceTypeByPathMutex.Lock()
	fake.findResourceTypeByPathArgsForCall = append(fake.findResourceTypeByPathArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("FindResourceTypeByPath", []interface{}{path})
	fake.findResourceTypeByPathMutex.Unlock()
	if fake.FindResourceTypeByPathStub != nil {
		return fake.FindResourceTypeByPathStub(path)
	} else {
		return fake.findResourceTypeByPathReturns.result1, fake.findResourceTypeByPathReturns.result2
	}
}

func (fake *FakeClient) FindResourceTypeByPathCallCount() int {
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	return len(fake.findResourceTypeByPathArgsForCall)
}

func (fake *FakeClient) FindResourceTypeByPathArgsForCall(i int) string {
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	return fake.findResourceTypeByPathArgsForCall[i].path
}

func (fake *FakeClient) FindResourceTypeByPathReturns(result1 atc.WorkerResourceType, result2 bool) {
	fake.FindResourceTypeByPathStub = nil
	fake.findResourceTypeByPathReturns = struct {
		result1 atc.WorkerResourceType
		result2 bool
	}{result1, result2}
}

func (fake *FakeClient) LookupVolume(arg1 lager.Logger, arg2 string) (worker.Volume, bool, error) {
	fake.lookupVolumeMutex.Lock()
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupVolume", []interface{}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(arg1, arg2)
	} else {
		return fake.lookupVolumeReturns.result1, fake.lookupVolumeReturns.result2, fake.lookupVolumeReturns.result3
	}
}

func (fake *FakeClient) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeClient) LookupVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.lookupVolumeArgsForCall[i].arg1, fake.lookupVolumeArgsForCall[i].arg2
}

func (fake *FakeClient) LookupVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Satisfying(arg1 worker.WorkerSpec, arg2 atc.ResourceTypes) (worker.Worker, error) {
	fake.satisfyingMutex.Lock()
	fake.satisfyingArgsForCall = append(fake.satisfyingArgsForCall, struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}{arg1, arg2})
	fake.recordInvocation("Satisfying", []interface{}{arg1, arg2})
	fake.satisfyingMutex.Unlock()
	if fake.SatisfyingStub != nil {
		return fake.SatisfyingStub(arg1, arg2)
	} else {
		return fake.satisfyingReturns.result1, fake.satisfyingReturns.result2
	}
}

func (fake *FakeClient) SatisfyingCallCount() int {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return len(fake.satisfyingArgsForCall)
}

func (fake *FakeClient) SatisfyingArgsForCall(i int) (worker.WorkerSpec, atc.ResourceTypes) {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return fake.satisfyingArgsForCall[i].arg1, fake.satisfyingArgsForCall[i].arg2
}

func (fake *FakeClient) SatisfyingReturns(result1 worker.Worker, result2 error) {
	fake.SatisfyingStub = nil
	fake.satisfyingReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) AllSatisfying(arg1 worker.WorkerSpec, arg2 atc.ResourceTypes) ([]worker.Worker, error) {
	fake.allSatisfyingMutex.Lock()
	fake.allSatisfyingArgsForCall = append(fake.allSatisfyingArgsForCall, struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}{arg1, arg2})
	fake.recordInvocation("AllSatisfying", []interface{}{arg1, arg2})
	fake.allSatisfyingMutex.Unlock()
	if fake.AllSatisfyingStub != nil {
		return fake.AllSatisfyingStub(arg1, arg2)
	} else {
		return fake.allSatisfyingReturns.result1, fake.allSatisfyingReturns.result2
	}
}

func (fake *FakeClient) AllSatisfyingCallCount() int {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return len(fake.allSatisfyingArgsForCall)
}

func (fake *FakeClient) AllSatisfyingArgsForCall(i int) (worker.WorkerSpec, atc.ResourceTypes) {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return fake.allSatisfyingArgsForCall[i].arg1, fake.allSatisfyingArgsForCall[i].arg2
}

func (fake *FakeClient) AllSatisfyingReturns(result1 []worker.Worker, result2 error) {
	fake.AllSatisfyingStub = nil
	fake.allSatisfyingReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RunningWorkers() ([]worker.Worker, error) {
	fake.runningWorkersMutex.Lock()
	fake.runningWorkersArgsForCall = append(fake.runningWorkersArgsForCall, struct{}{})
	fake.recordInvocation("RunningWorkers", []interface{}{})
	fake.runningWorkersMutex.Unlock()
	if fake.RunningWorkersStub != nil {
		return fake.RunningWorkersStub()
	} else {
		return fake.runningWorkersReturns.result1, fake.runningWorkersReturns.result2
	}
}

func (fake *FakeClient) RunningWorkersCallCount() int {
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	return len(fake.runningWorkersArgsForCall)
}

func (fake *FakeClient) RunningWorkersReturns(result1 []worker.Worker, result2 error) {
	fake.RunningWorkersStub = nil
	fake.runningWorkersReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetWorker(workerName string) (worker.Worker, error) {
	fake.getWorkerMutex.Lock()
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		workerName string
	}{workerName})
	fake.recordInvocation("GetWorker", []interface{}{workerName})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(workerName)
	} else {
		return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2
	}
}

func (fake *FakeClient) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeClient) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].workerName
}

func (fake *FakeClient) GetWorkerReturns(result1 worker.Worker, result2 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateBuildContainerMutex.RLock()
	defer fake.findOrCreateBuildContainerMutex.RUnlock()
	fake.findOrCreateResourceGetContainerMutex.RLock()
	defer fake.findOrCreateResourceGetContainerMutex.RUnlock()
	fake.findOrCreateResourceCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceCheckContainerMutex.RUnlock()
	fake.findOrCreateResourceTypeCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceTypeCheckContainerMutex.RUnlock()
	fake.findOrCreateContainerForIdentifierMutex.RLock()
	defer fake.findOrCreateContainerForIdentifierMutex.RUnlock()
	fake.findOrCreateVolumeForResourceCacheMutex.RLock()
	defer fake.findOrCreateVolumeForResourceCacheMutex.RUnlock()
	fake.findInitializedVolumeForResourceCacheMutex.RLock()
	defer fake.findInitializedVolumeForResourceCacheMutex.RUnlock()
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	fake.validateResourceCheckVersionMutex.RLock()
	defer fake.validateResourceCheckVersionMutex.RUnlock()
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ worker.Client = new(FakeClient)
