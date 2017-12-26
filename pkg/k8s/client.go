package k8s

import (
	"io"

	"github.com/nii236/k/pkg/k"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RealClientSet contains an embedded Kubernetes client set
type RealClientSet struct {
	clientset *kubernetes.Clientset
}

// New returns a new clientset
func New(flags *k.ParsedFlags) (*RealClientSet, error) {
	// Use the current context in kubeconfig
	cc, err := clientcmd.BuildConfigFromFlags("", flags.KubeConfigPath)
	if err != nil {
		return nil, err
	}

	// Create the client set
	clientSet, err := kubernetes.NewForConfig(cc)
	if err != nil {
		return nil, err
	}

	return &RealClientSet{
		clientSet,
	}, nil
}

// Get pods (use namespace)
func (cs *RealClientSet) GetPods(namespace string) (*v1.PodList, error) {
	return cs.clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
}

// Get namespaces
func (cs *RealClientSet) GetNamespaces() (*v1.NamespaceList, error) {
	return cs.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
}

// Get the pod containers
func (cs *RealClientSet) GetPodContainers(podName string, namespace string) []string {
	var pc []string

	pod, _ := cs.clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	for _, c := range pod.Spec.Containers {
		pc = append(pc, c.Name)
	}

	return pc
}

// Delete pod
func (cs *RealClientSet) DeletePod(podName string, namespace string) error {
	return cs.clientset.CoreV1().Pods(namespace).Delete(podName, &metav1.DeleteOptions{})
}

// Get pod container logs
func (cs *RealClientSet) GetPodContainerLogs(podName string, containerName string, namespace string, o io.Writer) error {
	tl := int64(50)

	opts := &v1.PodLogOptions{
		Container: containerName,
		TailLines: &tl,
	}

	req := cs.clientset.CoreV1().Pods(namespace).GetLogs(podName, opts)

	readCloser, err := req.Stream()
	if err != nil {
		return err
	}

	_, err = io.Copy(o, readCloser)

	readCloser.Close()

	return err
}