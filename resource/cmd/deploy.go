package cmd

import (
	"io"
	"os"
	"strings"

	"github.com/juju/errors"
	"github.com/juju/loggo"
	charmresource "gopkg.in/juju/charm.v6-unstable/resource"

	"github.com/juju/juju/api"
	"github.com/juju/juju/api/base"
	"github.com/juju/juju/resource"
	"github.com/juju/juju/resource/api/client"
	"github.com/juju/juju/resource/api/server"
)

var logger = loggo.GetLogger("resource/cmd")

type uploadClient interface {
	// AddPendingResources adds pending metadata for store-based resources.
	AddPendingResources(serviceID string, resources []charmresource.Resource) (ids []string, err error)
	// AddPendingResource uploads data and metadata for a pending resource for the given service.
	AddPendingResource(serviceID string, resource charmresource.Resource, r io.ReadSeeker) (id string, err error)
}

type deployUploader struct {
	serviceID string
	resources map[string]charmresource.Meta
	client    uploadClient
	osOpen    func(path string) (ReadSeekCloser, error)
}

// DeployResources uploads the bytes and metadata for the given resourcename -
// filename pairs, as well as uploading metadata for any resources not mentioned
// in the files. It returns a map of resource name to pending resource IDs.
func DeployResources(serviceID string, files map[string]string, resources map[string]charmresource.Meta, conn api.Connection) (ids map[string]string, err error) {
	client, err := newClient(conn)
	if err != nil {
		return nil, errors.Trace(err)
	}
	d := deployUploader{
		serviceID: serviceID,
		client:    client,
		resources: resources,
		osOpen:    func(s string) (ReadSeekCloser, error) { return os.Open(s) },
	}

	return d.upload(files)
}

func (d deployUploader) upload(files map[string]string) (map[string]string, error) {
	if err := d.validateResources(); err != nil {
		return nil, errors.Trace(err)
	}

	if err := d.checkExpectedResources(files); err != nil {
		return nil, errors.Trace(err)
	}

	storeResources := d.storeResources(files)
	pending := map[string]string{}
	if len(storeResources) > 0 {
		ids, err := d.client.AddPendingResources(d.serviceID, storeResources)
		if err != nil {
			return nil, errors.Trace(err)
		}
		// guaranteed 1:1 correlation between ids and resources.
		for i, res := range storeResources {
			pending[res.Name] = ids[i]
		}
	}

	for name, filename := range files {
		id, err := d.uploadFile(name, filename)
		if err != nil {
			return nil, errors.Trace(err)
		}
		pending[name] = id
	}

	return pending, nil
}

func (d deployUploader) validateResources() error {
	var errs []error
	for _, meta := range d.resources {
		if err := meta.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 1 {
		return errors.Trace(errs[0])
	}
	if len(errs) > 1 {
		msgs := make([]string, len(errs))
		for i, err := range errs {
			msgs[i] = err.Error()
		}
		return errors.NewNotValid(nil, strings.Join(msgs, ", "))
	}
	return nil
}

func (d deployUploader) storeResources(uploads map[string]string) []charmresource.Resource {
	var resources []charmresource.Resource
	for name, meta := range d.resources {
		if _, ok := uploads[name]; !ok {
			resources = append(resources, charmresource.Resource{
				Meta:   meta,
				Origin: charmresource.OriginStore,
				// Revision, Fingerprint, and Size will be added server-side,
				// when we download the bytes from the store.
			})
		}
	}
	return resources
}

func (d deployUploader) uploadFile(resourcename, filename string) (id string, err error) {
	f, err := d.osOpen(filename)
	if err != nil {
		return "", errors.Trace(err)
	}
	defer f.Close()
	res := charmresource.Resource{
		Meta:   d.resources[resourcename],
		Origin: charmresource.OriginUpload,
	}

	id, err = d.client.AddPendingResource(d.serviceID, res, f)
	if err != nil {
		return "", errors.Trace(err)
	}
	return id, err
}

func (d deployUploader) checkExpectedResources(provided map[string]string) error {
	var unknown []string

	for name := range provided {
		if _, ok := d.resources[name]; !ok {
			unknown = append(unknown, name)
		}
	}
	if len(unknown) == 1 {
		return errors.Errorf("unrecognized resource %q", unknown[0])
	}
	if len(unknown) > 1 {
		return errors.Errorf("unrecognized resources: %s", strings.Join(unknown, ", "))
	}
	return nil
}

// newClient is mostly a copy of the newClient code in
// component/all/resources.go.  It lives here because it simplifies this code
// immensely.
func newClient(conn api.Connection) (*client.Client, error) {
	caller := base.NewFacadeCallerForVersion(conn, resource.ComponentName, server.Version)

	cl, err := conn.HTTPClient()
	if err != nil {
		return nil, errors.Trace(err)
	}
	// The apiCaller takes care of prepending /environment/<envUUID>.
	return client.NewClient(caller, cl, conn), nil
}
