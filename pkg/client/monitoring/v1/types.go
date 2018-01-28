// Copyright 2016 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Prometheus defines a Prometheus deployment.
type Prometheus struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the Prometheus cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec PrometheusSpec `json:"spec"`
	// Most recent observed status of the Prometheus cluster. Read-only. Not
	// included when requesting from the apiserver, only from the Prometheus
	// Operator API itself. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status *PrometheusStatus `json:"status,omitempty"`
}

// PrometheusList is a list of Prometheuses.
type PrometheusList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of Prometheuses
	Items []*Prometheus `json:"items"`
}

// Specification of the desired behavior of the Prometheus cluster. More info:
// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
type PrometheusSpec struct {
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	// Metadata Labels and Annotations gets propagated to the prometheus pods.
	PodMetadata *metav1.ObjectMeta `json:"podMetadata,omitempty"`
	// ServiceMonitors to be selected for target discovery.
	ServiceMonitorSelector *metav1.LabelSelector `json:"serviceMonitorSelector,omitempty"`
	// Version of Prometheus to be deployed.
	Version string `json:"version,omitempty"`
	// When a Prometheus deployment is paused, no actions except for deletion
	// will be performed on the underlying objects.
	Paused bool `json:"paused,omitempty"`
	// Base image to use for a Prometheus deployment.
	BaseImage string `json:"baseImage,omitempty"`
	// An optional list of references to secrets in the same namespace
	// to use for pulling prometheus and alertmanager images from registries
	// see http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets []v1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	// Number of instances to deploy for a Prometheus deployment.
	Replicas *int32 `json:"replicas,omitempty"`
	// Time duration Prometheus shall retain data for.
	Retention string `json:"retention,omitempty"`
	// Log level for Prometheus be configured in.
	LogLevel string `json:"logLevel,omitempty"`
	// Interval between consecutive scrapes.
	ScrapeInterval string `json:"scrapeInterval,omitempty"`
	// Interval between consecutive evaluations.
	EvaluationInterval string `json:"evaluationInterval,omitempty"`
	// The labels to add to any time series or alerts when communicating with
	// external systems (federation, remote storage, Alertmanager).
	ExternalLabels map[string]string `json:"externalLabels,omitempty"`
	// The external URL the Prometheus instances will be available under. This is
	// necessary to generate correct URLs. This is necessary if Prometheus is not
	// served from root of a DNS name.
	ExternalURL string `json:"externalUrl,omitempty"`
	// The route prefix Prometheus registers HTTP handlers for. This is useful,
	// if using ExternalURL and a proxy is rewriting HTTP routes of a request,
	// and the actual ExternalURL is still true, but the server serves requests
	// under a different route prefix. For example for use with `kubectl proxy`.
	RoutePrefix string `json:"routePrefix,omitempty"`
	// Storage spec to specify how storage shall be used.
	Storage *StorageSpec `json:"storage,omitempty"`
	// A selector to select which ConfigMaps to mount for loading rule files from.
	RuleSelector *metav1.LabelSelector `json:"ruleSelector,omitempty"`
	// Define details regarding alerting.
	Alerting AlertingSpec `json:"alerting,omitempty"`
	// Define resources requests and limits for single Pods.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// Define which Nodes the Pods are scheduled on.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// ServiceAccountName is the name of the ServiceAccount to use to run the
	// Prometheus Pods.
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
	// Secrets is a list of Secrets in the same namespace as the Prometheus
	// object, which shall be mounted into the Prometheus Pods.
	// The Secrets are mounted into /etc/prometheus/secrets/<secret-name>.
	// Secrets changes after initial creation of a Prometheus object are not
	// reflected in the running Pods. To change the secrets mounted into the
	// Prometheus Pods, the object must be deleted and recreated with the new list
	// of secrets.
	Secrets []string `json:"secrets,omitempty"`
	// If specified, the pod's scheduling constraints.
	Affinity *v1.Affinity `json:"affinity,omitempty"`
	// If specified, the pod's tolerations.
	Tolerations []v1.Toleration `json:"tolerations,omitempty"`
	// If specified, the remote_write spec. This is an experimental feature, it may change in any upcoming release in a breaking way.
	RemoteWrite []RemoteWriteSpec `json:"remoteWrite,omitempty"`
	// If specified, the remote_read spec. This is an experimental feature, it may change in any upcoming release in a breaking way.
	RemoteRead []RemoteReadSpec `json:"remoteRead,omitempty"`
	// SecurityContext holds pod-level security attributes and common container settings.
	// This defaults to non root user with uid 1000 and gid 2000 for Prometheus >v2.0 and
	// default PodSecurityContext for other versions.
	SecurityContext *v1.PodSecurityContext
	// If specified, the pod hostNetwork is enabled
	HostNetwork bool `json:"hostNetwork,omitempty"`
 	// If specified, the prometheus -storage.local.chunk-encoding-version is set as this value
	ChunkEncodingVersion int32 `json:"chuckEncodingVersion,omitempty"`
	// If specified, set target-heap-size = requests.memory * targetHeapSizeRate, else set the rate to 1/2
	// The valid value is in range (0, 1)
	TargetHeapSizeRate string `json:"targetHeapSizeRate,omitempty"`
}

// Most recent observed status of the Prometheus cluster. Read-only. Not
// included when requesting from the apiserver, only from the Prometheus
// Operator API itself. More info:
// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
type PrometheusStatus struct {
	// Represents whether any actions on the underlaying managed objects are
	// being performed. Only delete actions will be performed.
	Paused bool `json:"paused"`
	// Total number of non-terminated pods targeted by this Prometheus deployment
	// (their labels match the selector).
	Replicas int32 `json:"replicas"`
	// Total number of non-terminated pods targeted by this Prometheus deployment
	// that have the desired version spec.
	UpdatedReplicas int32 `json:"updatedReplicas"`
	// Total number of available pods (ready for at least minReadySeconds)
	// targeted by this Prometheus deployment.
	AvailableReplicas int32 `json:"availableReplicas"`
	// Total number of unavailable pods targeted by this Prometheus deployment.
	UnavailableReplicas int32 `json:"unavailableReplicas"`
}

// AlertingSpec defines parameters for alerting configuration of Prometheus servers.
type AlertingSpec struct {
	// AlertmanagerEndpoints Prometheus should fire alerts against.
	Alertmanagers []AlertmanagerEndpoints `json:"alertmanagers"`
}

// StorageSpec defines the configured storage for a group Prometheus servers.
type StorageSpec struct {
	// Name of the StorageClass to use when requesting storage provisioning. More
	// info: https://kubernetes.io/docs/user-guide/persistent-volumes/#storageclasses
	// DEPRECATED
	Class string `json:"class"`
	// EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. More
	// info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir
	EmptyDir *v1.EmptyDirVolumeSource `json:"emptyDir,omitempty"`
	// A label query over volumes to consider for binding.
	// DEPRECATED
	Selector *metav1.LabelSelector `json:"selector"`
	// Resources represents the minimum resources the volume should have. More
	// info: http://kubernetes.io/docs/user-guide/persistent-volumes#resources
	// DEPRECATED
	Resources v1.ResourceRequirements `json:"resources"`
	// A PVC spec to be used by the Prometheus StatefulSets.
	VolumeClaimTemplate v1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`
}

// RemoteWriteSpec defines the remote_write configuration for prometheus.
type RemoteWriteSpec struct {
	//The URL of the endpoint to send samples to.
	URL string `json:"url"`
	//Timeout for requests to the remote write endpoint.
	RemoteTimeout string `json:"remoteTimeout,omitempty"`
	//The list of remote write relabel configurations.
	WriteRelabelConfigs []RelabelConfig `json:"writeRelabelConfigs,omitempty"`
	//BasicAuth for the URL.
	BasicAuth *BasicAuth `json:"basicAuth,omitempty"`
	// File to read bearer token for remote write.
	BearerToken string `json:"bearerToken,omitempty"`
	// File to read bearer token for remote write.
	BearerTokenFile string `json:"bearerTokenFile,omitempty"`
	// TLS Config to use for remote write.
	TLSConfig *TLSConfig `json:"tlsConfig,omitempty"`
	//Optional ProxyURL
	ProxyURL string `json:"proxy_url,omitempty"`
}

// RemoteReadSpec defines the remote_read configuration for prometheus.
type RemoteReadSpec struct {
	//The URL of the endpoint to send samples to.
	URL string `json:"url"`
	//Timeout for requests to the remote write endpoint.
	RemoteTimeout string `json:"remoteTimeout,omitempty"`
	//BasicAuth for the URL.
	BasicAuth *BasicAuth `json:"basicAuth,omitempty"`
	// bearer token for remote write.
	BearerToken string `json:"bearerToken,omitempty"`
	// File to read bearer token for remote write.
	BearerTokenFile string `json:"bearerTokenFile,omitempty"`
	// TLS Config to use for remote write.
	TLSConfig *TLSConfig `json:"tlsConfig,omitempty"`
	//Optional ProxyURL
	ProxyURL string `json:"proxy_url,omitempty"`
}

// RelabelConfig allows dynamic rewriting of the label set.
type RelabelConfig struct {
	//The source labels select values from existing labels. Their content is concatenated
	//using the configured separator and matched against the configured regular expression
	//for the replace, keep, and drop actions.
	SourceLabels []string `json:"sourceLabels"`
	//Separator placed between concatenated source label values. default is ';'.
	Separator string `json:"separator,omitempty"`
	//Label to which the resulting value is written in a replace action.
	//It is mandatory for replace actions. Regex capture groups are available.
	TargetLabel string `json:"targetLabel,omitempty"`
	//Regular expression against which the extracted value is matched. defailt is '(.*)'
	Regex string `json:"regex,omitempty"`
	// Modulus to take of the hash of the source label values.
	Modulus uint64 `json:"modulus,omitempty"`
	//Replacement value against which a regex replace is performed if the
	//regular expression matches. Regex capture groups are available. Default is '$1'
	Replacement string `json:"replacement"`
	// Action to perform based on regex matching. Default is 'replace'
	Action string `json:"action,omitempty"`
}

// AlertmanagerEndpoints defines a selection of a single Endpoints object
// containing alertmanager IPs to fire alerts against.
type AlertmanagerEndpoints struct {
	// Namespace of Endpoints object.
	Namespace string `json:"namespace"`
	// Name of Endpoints object in Namespace.
	Name string `json:"name"`
	// Port the Alertmanager API is exposed on.
	Port intstr.IntOrString `json:"port"`
	// Scheme to use when firing alerts.
	Scheme string `json:"scheme"`
	// Prefix for the HTTP path alerts are pushed to.
	PathPrefix string `json:"pathPrefix"`
}

// ServiceMonitor defines monitoring for a set of services.
type ServiceMonitor struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of desired Service selection for target discrovery by
	// Prometheus.
	Spec ServiceMonitorSpec `json:"spec"`
}

// ServiceMonitorSpec contains specification parameters for a ServiceMonitor.
type ServiceMonitorSpec struct {
	// The label to use to retrieve the job name from.
	JobLabel string `json:"jobLabel,omitempty"`
	// A list of endpoints allowed as part of this ServiceMonitor.
	Endpoints []Endpoint `json:"endpoints"`
	// Selector to select Endpoints objects.
	Selector metav1.LabelSelector `json:"selector"`
	// Selector to select which namespaces the Endpoints objects are discovered from.
	NamespaceSelector NamespaceSelector `json:"namespaceSelector,omitempty"`
}

// Endpoint defines a scrapeable endpoint serving Prometheus metrics.
type Endpoint struct {
	// Name of the service port this endpoint refers to. Mutually exclusive with targetPort.
	Port string `json:"port,omitempty"`
	// Name or number of the target port of the endpoint. Mutually exclusive with port.
	TargetPort intstr.IntOrString `json:"targetPort,omitempty"`
	// HTTP path to scrape for metrics.
	Path string `json:"path,omitempty"`
	// HTTP scheme to use for scraping.
	Scheme string `json:"scheme,omitempty"`
	// Optional HTTP URL parameters
	Params map[string][]string `json:"params,omitempty"`
	// Interval at which metrics should be scraped
	Interval string `json:"interval,omitempty"`
	// Timeout after which the scrape is ended
	ScrapeTimeout string `json:"scrapeTimeout,omitempty"`
	// TLS configuration to use when scraping the endpoint
	TLSConfig *TLSConfig `json:"tlsConfig,omitempty"`
	// File to read bearer token for scraping targets.
	BearerTokenFile string `json:"bearerTokenFile,omitempty"`
	// HonorLabels chooses the metric's labels on collisions with target labels.
	HonorLabels bool `json:"honorLabels,omitempty"`
	// BasicAuth allow an endpoint to authenticate over basic authentication
	// More info: https://prometheus.io/docs/operating/configuration/#endpoints
	BasicAuth *BasicAuth `json:"basicAuth,omitempty"`
	// MetricRelabelConfigs to apply to samples before ingestion.
	MetricRelabelConfigs []*RelabelConfig `json:"metricRelabelings,omitempty"`
}

// BasicAuth allow an endpoint to authenticate over basic authentication
// More info: https://prometheus.io/docs/operating/configuration/#endpoints
type BasicAuth struct {
	// The secret that contains the username for authenticate
	Username v1.SecretKeySelector `json:"username,omitempty"`
	// The secret that contains the password for authenticate
	Password v1.SecretKeySelector `json:"password,omitempty"`
}

// TLSConfig specifies TLS configuration parameters.
type TLSConfig struct {
	// The CA cert to use for the targets.
	CAFile string `json:"caFile,omitempty"`
	// The client cert file for the targets.
	CertFile string `json:"certFile,omitempty"`
	// The client key file for the targets.
	KeyFile string `json:"keyFile,omitempty"`
	// Used to verify the hostname for the targets.
	ServerName string `json:"serverName,omitempty"`
	// Disable target certificate validation.
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}

// A list of ServiceMonitors.
type ServiceMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of ServiceMonitors
	Items []*ServiceMonitor `json:"items"`
}

// Describes an Alertmanager cluster.
type Alertmanager struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the Alertmanager cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec AlertmanagerSpec `json:"spec"`
	// Most recent observed status of the Alertmanager cluster. Read-only. Not
	// included when requesting from the apiserver, only from the Prometheus
	// Operator API itself. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status *AlertmanagerStatus `json:"status,omitempty"`
}

// Specification of the desired behavior of the Alertmanager cluster. More info:
// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
type AlertmanagerSpec struct {
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	// Metadata Labels and Annotations gets propagated to the prometheus pods.
	PodMetadata *metav1.ObjectMeta `json:"podMetadata,omitempty"`
	// Version the cluster should be on.
	Version string `json:"version,omitempty"`
	// Base image that is used to deploy pods.
	BaseImage string `json:"baseImage,omitempty"`
	// An optional list of references to secrets in the same namespace
	// to use for pulling prometheus and alertmanager images from registries
	// see http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets []v1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	// Size is the expected size of the alertmanager cluster. The controller will
	// eventually make the size of the running cluster equal to the expected
	// size.
	Replicas *int32 `json:"replicas,omitempty"`
	// Storage is the definition of how storage will be used by the Alertmanager
	// instances.
	Storage *StorageSpec `json:"storage,omitempty"`
	// The external URL the Alertmanager instances will be available under. This is
	// necessary to generate correct URLs. This is necessary if Alertmanager is not
	// served from root of a DNS name.
	ExternalURL string `json:"externalUrl,omitempty"`
	// The route prefix Alertmanager registers HTTP handlers for. This is useful,
	// if using ExternalURL and a proxy is rewriting HTTP routes of a request,
	// and the actual ExternalURL is still true, but the server serves requests
	// under a different route prefix. For example for use with `kubectl proxy`.
	RoutePrefix string `json:"routePrefix,omitempty"`
	// If set to true all actions on the underlaying managed objects are not
	// goint to be performed, except for delete actions.
	Paused bool `json:"paused,omitempty"`
	// Define which Nodes the Pods are scheduled on.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// Define resources requests and limits for single Pods.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// If specified, the pod's scheduling constraints.
	Affinity *v1.Affinity `json:"affinity,omitempty"`
	// If specified, the pod's tolerations.
	Tolerations []v1.Toleration `json:"tolerations,omitempty"`
	// If specified, the pod's hostNetwork will be true
	HostNetwork bool `json:"hostNetwork,omitempty"`
}

// A list of Alertmanagers.
type AlertmanagerList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of Alertmanagers
	Items []Alertmanager `json:"items"`
}

// Most recent observed status of the Alertmanager cluster. Read-only. Not
// included when requesting from the apiserver, only from the Prometheus
// Operator API itself. More info:
// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
type AlertmanagerStatus struct {
	// Represents whether any actions on the underlaying managed objects are
	// being performed. Only delete actions will be performed.
	Paused bool `json:"paused"`
	// Total number of non-terminated pods targeted by this Alertmanager
	// cluster (their labels match the selector).
	Replicas int32 `json:"replicas"`
	// Total number of non-terminated pods targeted by this Alertmanager
	// cluster that have the desired version spec.
	UpdatedReplicas int32 `json:"updatedReplicas"`
	// Total number of available pods (ready for at least minReadySeconds)
	// targeted by this Alertmanager cluster.
	AvailableReplicas int32 `json:"availableReplicas"`
	// Total number of unavailable pods targeted by this Alertmanager cluster.
	UnavailableReplicas int32 `json:"unavailableReplicas"`
}

// A selector for selecting namespaces either selecting all namespaces or a
// list of namespaces.
type NamespaceSelector struct {
	// Boolean describing whether all namespaces are selected in contrast to a
	// list restricting them.
	Any bool `json:"any,omitempty"`
	// List of namespace names.
	MatchNames []string `json:"matchNames,omitempty"`

	// TODO(fabxc): this should embed metav1.LabelSelector eventually.
	// Currently the selector is only used for namespaces which require more complex
	// implementation to support label selections.
}

func (l *Alertmanager) DeepCopyObject() runtime.Object {
	return l.DeepCopy()
}

func (l *AlertmanagerList) DeepCopyObject() runtime.Object {
	return l.DeepCopy()
}

func (l *Prometheus) DeepCopyObject() runtime.Object {
	return l.DeepCopy()
}

func (l *PrometheusList) DeepCopyObject() runtime.Object {
	return l.DeepCopy()
}

func (l *ServiceMonitor) DeepCopyObject() runtime.Object {
	return l.DeepCopy()
}

func (l *ServiceMonitorList) DeepCopyObject() runtime.Object {
	return l.DeepCopy()
}
