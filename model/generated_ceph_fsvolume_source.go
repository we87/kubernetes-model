package model

const (
	CEPH_FSVOLUME_SOURCE_TYPE = "v1.CephFSVolumeSource"
)

type CephFSVolumeSource struct {
	Monitors []string `json:"monitors,omitempty" yaml:"monitors,omitempty"`

	ReadOnly bool `json:"readOnly,omitempty" yaml:"read_only,omitempty"`

	SecretFile string `json:"secretFile,omitempty" yaml:"secret_file,omitempty"`

	SecretRef *LocalObjectReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`

	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
