package delegate

import (
	"strings"

	"github.com/google/uuid"
	spoolerpkg "github.com/tliron/kubernetes-registry-spooler/client"
	"github.com/tliron/puccini/common/format"
	urlpkg "github.com/tliron/puccini/url"
	"github.com/tliron/turandot/common"
	resources "github.com/tliron/turandot/resources/turandot.puccini.cloud/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (self *Client) DeployServiceFromTemplate(serviceName string, serviceTemplateName string, inputs map[string]interface{}, mode string, urlContext *urlpkg.Context) error {
	if url, err := self.GetInventoryServiceTemplateURL(serviceTemplateName, urlContext); err == nil {
		_, err := self.CreateService(serviceName, url, inputs, mode)
		return err
	} else {
		return err
	}
}

func (self *Client) DeployServiceFromURL(serviceName string, url string, inputs map[string]interface{}, mode string, urlContext *urlpkg.Context) error {
	if url_, err := urlpkg.NewURL(url, urlContext); err == nil {
		_, err = self.CreateService(serviceName, url_, inputs, mode)
		return err
	} else {
		return err
	}
}

func (self *Client) DeployServiceFromContent(serviceName string, spooler *spoolerpkg.Client, url urlpkg.URL, inputs map[string]interface{}, mode string, urlContext *urlpkg.Context) error {
	serviceTemplateName := uuid.New().String()
	imageName := GetInventoryImageName(serviceTemplateName)
	if err := common.PublishOnRegistry(imageName, url, spooler); err == nil {
		return self.DeployServiceFromTemplate(serviceName, serviceTemplateName, inputs, mode, urlContext)
	} else {
		return err
	}
}

func (self *Client) GetService(serviceName string) (*resources.Service, error) {
	if service, err := self.Turandot.TurandotV1alpha1().Services(self.Namespace).Get(self.Context, serviceName, meta.GetOptions{}); err == nil {
		// When retrieved from cache the GVK may be empty
		if service.Kind == "" {
			service = service.DeepCopy()
			service.APIVersion, service.Kind = resources.ServiceGVK.ToAPIVersionAndKind()
		}
		return service, nil
	} else {
		return nil, err
	}
}

func (self *Client) UpdateServiceSpec(service *resources.Service) (*resources.Service, error) {
	if service_, err := self.Turandot.TurandotV1alpha1().Services(service.Namespace).Update(self.Context, service, meta.UpdateOptions{}); err == nil {
		// When retrieved from cache the GVK may be empty
		if service_.Kind == "" {
			service_ = service_.DeepCopy()
			service_.APIVersion, service_.Kind = resources.ServiceGVK.ToAPIVersionAndKind()
		}
		return service_, nil
	} else {
		return service, err
	}
}

func (self *Client) UpdateServiceStatus(service *resources.Service) (*resources.Service, error) {
	if service_, err := self.Turandot.TurandotV1alpha1().Services(service.Namespace).UpdateStatus(self.Context, service, meta.UpdateOptions{}); err == nil {
		// When retrieved from cache the GVK may be empty
		if service_.Kind == "" {
			service_ = service_.DeepCopy()
			service_.APIVersion, service_.Kind = resources.ServiceGVK.ToAPIVersionAndKind()
		}
		return service_, nil
	} else {
		return service, err
	}
}

func (self *Client) DeleteService(serviceName string) error {
	return self.Turandot.TurandotV1alpha1().Services(self.Namespace).Delete(self.Context, serviceName, meta.DeleteOptions{})
}

func (self *Client) ListServices() (*resources.ServiceList, error) {
	return self.Turandot.TurandotV1alpha1().Services(self.Namespace).List(self.Context, meta.ListOptions{})
}

func (self *Client) CreateService(name string, url urlpkg.URL, inputs map[string]interface{}, mode string) (*resources.Service, error) {
	// Encode inputs
	var inputs_ map[string]string
	if (inputs != nil) && len(inputs) > 0 {
		inputs_ = make(map[string]string)
		for key, input := range inputs {
			var err error
			if inputs_[key], err = format.EncodeYAML(input, " ", false); err == nil {
				inputs_[key] = strings.TrimRight(inputs_[key], "\n")
			} else {
				return nil, err
			}
		}
	}

	service := &resources.Service{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: self.Namespace,
		},
		Spec: resources.ServiceSpec{
			ServiceTemplateURL: url.String(),
			Inputs:             inputs_,
			Mode:               mode,
		},
	}

	if service, err := self.Turandot.TurandotV1alpha1().Services(self.Namespace).Create(self.Context, service, meta.CreateOptions{}); err == nil {
		return service, nil
	} else if errors.IsAlreadyExists(err) {
		return self.Turandot.TurandotV1alpha1().Services(self.Namespace).Get(self.Context, name, meta.GetOptions{})
	} else {
		return nil, err
	}
}
