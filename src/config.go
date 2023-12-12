package main

type Config struct {
	Version string `yaml:"version"`
	Debug   bool   `yaml:"debug"`
	Persist bool   `yaml:"persist"`
	Machine struct {
		Type  string `yaml:"type"`
		Token string `yaml:"token"`
		Ca    struct {
			Crt string `yaml:"crt"`
			Key string `yaml:"key"`
		} `yaml:"ca"`
		CertSANs []interface{} `yaml:"certSANs"`
		Kubelet  struct {
			Image                               string `yaml:"image"`
			DefaultRuntimeSeccompProfileEnabled bool   `yaml:"defaultRuntimeSeccompProfileEnabled"`
			DisableManifestsDirectory           bool   `yaml:"disableManifestsDirectory"`
		} `yaml:"kubelet"`
		Network struct {
			Interfaces []struct {
				Interface string `yaml:"interface"`
				Vip       struct {
					IP string `yaml:"ip"`
				} `yaml:"vip"`
				Addresses []string `yaml:"addresses"`
			} `yaml:"interfaces"`
		} `yaml:"network"`
		Install struct {
			Disk  string `yaml:"disk"`
			Image string `yaml:"image"`
			Wipe  bool   `yaml:"wipe"`
		} `yaml:"install"`
		Registries struct {
		} `yaml:"registries"`
		Features struct {
			Rbac                 bool `yaml:"rbac"`
			StableHostname       bool `yaml:"stableHostname"`
			ApidCheckExtKeyUsage bool `yaml:"apidCheckExtKeyUsage"`
			DiskQuotaSupport     bool `yaml:"diskQuotaSupport"`
		} `yaml:"features"`
	} `yaml:"machine"`
	Cluster struct {
		ID           string `yaml:"id"`
		Secret       string `yaml:"secret"`
		ControlPlane struct {
			Endpoint string `yaml:"endpoint"`
		} `yaml:"controlPlane"`
		ClusterName string `yaml:"clusterName"`
		Network     struct {
			DNSDomain      string   `yaml:"dnsDomain"`
			PodSubnets     []string `yaml:"podSubnets"`
			ServiceSubnets []string `yaml:"serviceSubnets"`
		} `yaml:"network"`
		Token                     string `yaml:"token"`
		SecretboxEncryptionSecret string `yaml:"secretboxEncryptionSecret"`
		Ca                        struct {
			Crt string `yaml:"crt"`
			Key string `yaml:"key"`
		} `yaml:"ca"`
		AggregatorCA struct {
			Crt string `yaml:"crt"`
			Key string `yaml:"key"`
		} `yaml:"aggregatorCA"`
		ServiceAccount struct {
			Key string `yaml:"key"`
		} `yaml:"serviceAccount"`
		APIServer struct {
			Image                    string   `yaml:"image"`
			CertSANs                 []string `yaml:"certSANs"`
			DisablePodSecurityPolicy bool     `yaml:"disablePodSecurityPolicy"`
			AdmissionControl         []struct {
				Name          string `yaml:"name"`
				Configuration struct {
					APIVersion string `yaml:"apiVersion"`
					Defaults   struct {
						Audit          string `yaml:"audit"`
						AuditVersion   string `yaml:"audit-version"`
						Enforce        string `yaml:"enforce"`
						EnforceVersion string `yaml:"enforce-version"`
						Warn           string `yaml:"warn"`
						WarnVersion    string `yaml:"warn-version"`
					} `yaml:"defaults"`
					Exemptions struct {
						Namespaces     []string      `yaml:"namespaces"`
						RuntimeClasses []interface{} `yaml:"runtimeClasses"`
						Usernames      []interface{} `yaml:"usernames"`
					} `yaml:"exemptions"`
					Kind string `yaml:"kind"`
				} `yaml:"configuration"`
			} `yaml:"admissionControl"`
			AuditPolicy struct {
				APIVersion string `yaml:"apiVersion"`
				Kind       string `yaml:"kind"`
				Rules      []struct {
					Level string `yaml:"level"`
				} `yaml:"rules"`
			} `yaml:"auditPolicy"`
		} `yaml:"apiServer"`
		ControllerManager struct {
			Image string `yaml:"image"`
		} `yaml:"controllerManager"`
		Proxy struct {
			Image string `yaml:"image"`
		} `yaml:"proxy"`
		Scheduler struct {
			Image string `yaml:"image"`
		} `yaml:"scheduler"`
		Discovery struct {
			Enabled    bool `yaml:"enabled"`
			Registries struct {
				Kubernetes struct {
					Disabled bool `yaml:"disabled"`
				} `yaml:"kubernetes"`
				Service struct {
				} `yaml:"service"`
			} `yaml:"registries"`
		} `yaml:"discovery"`
		Etcd struct {
			Ca struct {
				Crt string `yaml:"crt"`
				Key string `yaml:"key"`
			} `yaml:"ca"`
		} `yaml:"etcd"`
		ExtraManifests  []interface{} `yaml:"extraManifests"`
		InlineManifests []interface{} `yaml:"inlineManifests"`
	} `yaml:"cluster"`
}
