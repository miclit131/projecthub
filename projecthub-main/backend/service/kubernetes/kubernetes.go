package kubernetes

import (
	"context"
	"fmt"
	"os/exec"
	"projecthub/config"
	"projecthub/util/filemanager"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"

	//core "k8s.io/api/core"
	corev1 "k8s.io/api/core/v1"
	//meta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ApplySimple(configPath string) error {
	cmd := exec.Command("kubectl", "apply", "-f", configPath)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(stdout))
	return nil
}

//https://stackoverflow.com/questions/58783939/using-client-go-to-kubectl-apply-against-the-kubernetes-api-directly-with-mult
//https://github.com/billiford/go-clouddriver/blob/master/pkg/kubernetes/client.go#L63
func Apply(objConfigPath string) error {
	// get config
	k8sconfig, err := config.GetK8sRestConfig()
	if err != nil {
		panic(err)
	}

	// create kubernetes clientset. this clientset can be used to create,delete,patch,list etc for the kubernetes resources
	clientset, err := kubernetes.NewForConfig(k8sconfig)
	if err != nil {
		panic(err)
	}

	// read config
	objBytes, err := filemanager.ReadFile(objConfigPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(objBytes))
	// decode config
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, gvk, _ := decode(objBytes, nil, nil)

	// TODO fix this part
	if err := createObject(clientset, *k8sconfig, obj, *gvk); err != nil {
		panic(err)
	}

	fmt.Println("Pod(s) created successfully...")
	return nil
}

func createObject(kubeClientset kubernetes.Interface, restConfig rest.Config, obj runtime.Object, gvk schema.GroupVersionKind) error {
	// Create a REST mapper that tracks information about the available resources in the cluster.
	groupResources, err := restmapper.GetAPIGroupResources(kubeClientset.Discovery())
	if err != nil {
		panic(err)
	}
	rm := restmapper.NewDiscoveryRESTMapper(groupResources)

	// get some metadata needed to make the REST request
	//gvk := obj.GetObjectKind().GroupVersionKind() //take this or the gvk from decofing the obj?
	gk := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
	mapping, err := rm.RESTMapping(gk, gvk.Version)
	if err != nil {
		panic(err)
	}

	/* 	name, err := meta.NewAccessor().Name(obj)
	   	if err != nil {
	   		panic(err)
	   	} */
	//fmt.Println(name)

	// Create a client specifically for creating the object.
	restClient, err := newRestClient(restConfig, mapping.GroupVersionKind.GroupVersion())
	if err != nil {
		panic(err)
	}

	//TODO why new restCLient? kubeClientset.CoreV1().RESTClient()

	// Use the REST helper to create the object in the "default" namespace.
	restHelper := resource.NewHelper(restClient, mapping)

	//depObj, err := restHelper.CreateWithOptions("default", false, obj, &metav1.CreateOptions{})
	cObj, err := restHelper.Create("default", false, obj) //TODO Error
	if err != nil {
		panic(err)
	}
	fmt.Println(cObj)

	return nil
}

func newRestClient(restConfig rest.Config, gv schema.GroupVersion) (rest.Interface, error) {
	restConfig.ContentConfig = resource.UnstructuredPlusDefaultContentConfig()
	restConfig.GroupVersion = &gv
	if len(gv.Group) == 0 {
		restConfig.APIPath = "/api"
	} else {
		restConfig.APIPath = "/apis"
	}

	return rest.RESTClientFor(&restConfig)
}

/* 	// create pod in kubernetes cluster using the clientset
   	pod, err = clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
   	if err != nil {
   		panic(err)
   	} */

//https://stackoverflow.com/questions/53341727/how-to-submit-generic-runtime-object-to-kubernetes-api-using-client-go
func Apply2(objConfigPath string) error {
	// get config
	k8sconfig, err := config.GetK8sRestConfig()
	if err != nil {
		panic(err)
	}

	// dynamic client deals with unstructured.Unstructured objects and all runtime.Objects can be converted to it
	dynamicClient, err := dynamic.NewForConfig(k8sconfig)
	if err != nil {
		panic(err)
	}

	// create kubernetes clientset. this clientset can be used to create,delete,patch,list etc for the kubernetes resources
	clientset, err := kubernetes.NewForConfig(k8sconfig)
	if err != nil {
		panic(err)
	}

	// read obj config
	objBytes, err := filemanager.ReadFile(objConfigPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(objBytes))
	// decode config
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, gvk, _ := decode(objBytes, nil, nil)

	// convert the runtime.Object to unstructured map[string]interface{}
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		panic(err)
	}
	unstrObj := unstructured.Unstructured{/*Object: unstructuredObj*/} //Why does ToUnstructured() not return unstructured?!
	unstrObj.SetUnstructuredContent(unstructuredObj)
	// Create a REST mapper that tracks information about the available resources in the cluster.
	groupResources, err := restmapper.GetAPIGroupResources(clientset.Discovery())
	if err != nil {
		panic(err)
	}
	rm := restmapper.NewDiscoveryRESTMapper(groupResources)

	// get some metadata needed to make the REST request
	//gvk := obj.GetObjectKind().GroupVersionKind() //take this or the gvk from obj?
	gk := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
	mapping, err := rm.RESTMapping(gk, gvk.Version)
	if err != nil {
		panic(err)
	}

	// create the object using the dynamic client
	createdUnstructuredObj, err := dynamicClient.Resource(mapping.Resource).Namespace("default").Create(context.TODO(), &unstrObj, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	// convert unstructured.Unstructured to a Node
	var node *corev1.Node
	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(createdUnstructuredObj.Object, node); err != nil {
		panic(err)
	}

	return nil
}
