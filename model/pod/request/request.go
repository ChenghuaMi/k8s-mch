/**
 * @author mch
 */

package request

import "k8s-mch/model/base"

type Base struct {
	Name string `json:"name"`
	Labels []base.ListMapItem `json:"labels"`
	Namespace string `json:"namespace"`
	RestartPolicy string `json:"restartPolicy"` // Always  Never On-Failure
}

type Volume struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
//hostNetwork: false
//dnsPolicy: "Default"
//dnsConfig:
//nameservers:
//- 8.8.8.8
//hostAliases:
//- ip: 192.168.2.115
//hostnames:
//- "fool.local"
//- "bar.local"

type DnsConfig struct {
	Nameservers []string `json:"nameservers"`
}
type NetWorking struct {
	HostNetwork bool `json:"hostNetwork"`
	HostName string `json:"hostName"`
	DnsPolicy string `json:"dnsPolicy"`
	DnsConfig DnsConfig `json:"dnsConfig"`
	HostAliases []base.ListMapItem `json:"hostAliases"`

}
type Resources struct {
	Enable bool `json:"enable"`  // 是否配置容器配额
	MemRequest int32  `json:"memRequest"`
	MemLimit int32 `json:"memLimit"`
	CpuRequest int32 `json:"cpuRequest"`
	CpuLimit int32 `json:"cpuLimit"`
}
type VolumeMount struct {
	MountName string `json:"mountName"` //挂在卷名称
	MountPath string `json:"mountPath"` // 挂在卷 - 容器内的路径
	ReadOnly bool  `json:"readOnly"` // 是否只读
}

type ProbeHttpGet struct {
	Scheme string `json:"scheme"`  // http | https
	Host string `json:"host"`  // host 为空，pod 内请求
	Path string `json:"path"` // qi
	Port int32 `json:"port"`
	HttpHeaders []base.ListMapItem `json:"httpHeaders"` //请求head
}
type ProbeCommand struct {
	Command []string `json:"command"`
}
type ProbeTcpSocket struct {
	Host string `json:"host"`
	Port int32  `json:"port"`

}

type ProbeTime struct {
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`  //初始化时间多少秒 开始探针
	PeriodSeconds  int32  `json:"periodSeconds"` // 每隔多少秒 探针
	TimeoutSeconds int32 `json:"timeoutSeconds"` // 多少秒没有返回，探测失败
	SuccessThreshold int32 `json:"successThreshold"` // 探测多少次 ，算成功
	FailureThreshold int32  `json:"failureThreshold"` // 探测多少次，算失败
}
type ContainerProbe struct {
	Enable bool `json:"enable"`  // 是否打开探针
	Type string `json:"type"` // 探针类型 TCP  | http | exec
	HttpGet ProbeHttpGet `json:"httpGet"`
	Exec ProbeCommand  `json:"exec"`
	TcpSocket  ProbeTcpSocket `json:"tcpSocket"`
	ProbeTime
}
type ContainerPort struct {
	Name string `json:"name,omitempty"`
	ContainerPort int32 `json:"containerPort"`
	HostPort int32 `json:"hostPort"`
}
type Container struct {
	Name string `json:"name"` //容器名称
	Image string `json:"image"` //镜像
	ImagePullPolicy string `json:"imagePullPolicy"` //拉取策略
	Tty bool `json:"tty"` //为终端开启
	Ports []ContainerPort `json:"ports"`
	WorkingDir string `json:"workingDir"` //工作目录
	Command []string `json:"command"` // 命令
	Args []string `json:"args"`//参数
	Envs []base.ListMapItem `json:"envs"` //环境变量
	Privileged bool `json:"privileged"`
	Resources Resources `json:"resources"`
	VolumeMounts  []VolumeMount `json:"volumeMounts"`
	StartupProbe ContainerProbe `json:"startupProbe"` //启动探针
	LivenessProbe ContainerProbe `json:"livenessProbe"` //存活探针
	ReadinessProbe ContainerProbe  `json:"readinessProbe"` //就绪探针

}
type Pod struct {
	Base Base `json:"base"`
	Volumes []Volume `json:"volumes"`
	NetWorking NetWorking `json:"netWorking"`
	InitContainers []Container `json:"initContainers"`
	Containers []Container `json:"containers"`

}
