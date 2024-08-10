package global

import (
	"k8s.io/client-go/kubernetes"
	"kubeimooc.com/config"
)

// 全局变量，存储配置信息
var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)
