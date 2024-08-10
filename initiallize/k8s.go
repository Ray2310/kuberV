package initiallize

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kubeimooc.com/global"
)

// K8S 初始化k8s客户端
func K8S() {
	// kubeconfig 文件的路径
	kubeconfig := ".kube/config"
	// duq
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	global.KubeConfigSet = clientset
}
