package webhook

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Mutate mutates the pod spec and add the environment variables to the container
func Mutate(w http.ResponseWriter, r *http.Request) {
	// Parse the AdissionReview from the incoming request
	logrus.Info("Mutating request received")
	admissionReview := v1.AdmissionReview{}
	if err := json.NewDecoder(r.Body).Decode(&admissionReview); err != nil {
		logrus.WithError(err).Error("Failed to decode admission review")

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the AdmissionReview response
	response := v1.AdmissionResponse{
		UID:     admissionReview.Request.UID,
		Allowed: true,
		PatchType: func() *v1.PatchType {
			pt := v1.PatchTypeJSONPatch
			return &pt
		}(),
		Result: &metav1.Status{},
	}

	// Process the mutation
	logrus.Info("Mutating pod spec")
	mutatedPod, err := mutatePod(admissionReview.Request.Object.Raw)
	if err != nil {
		logrus.WithError(err).Error("Failed to mutate pod spec")

		response.Allowed = false
		response.Result.Status = "Failed"
		response.Result.Message = strconv.Quote(err.Error())
		response.Result.Reason = metav1.StatusReasonBadRequest
		response.Result.Details = &metav1.StatusDetails{
			Name:  "pod-webhook-mutator",
			Group: "admission.k8s.io",
			Kind:  "MutatingWebhookConfiguration",
			Causes: []metav1.StatusCause{
				{
					Type:    metav1.CauseTypeFieldValueInvalid,
					Message: err.Error(),
					Field:   "spec.containers",
				},
			},
		}
	} else {
		logrus.Info("Successfully mutated pod spec")

		response.Allowed = true
		response.Patch = mutatedPod
		response.Result.Status = "Success"
		response.Result.Message = "Successfully mutated pod"
		logrus.Infof("Successfully mutated pod: %v", string(mutatedPod))
	}

	// Encode the response
	admissionReview.Response = &response
	if err := json.NewEncoder(w).Encode(admissionReview); err != nil {
		logrus.WithError(err).Error("Failed to encode admission review response")

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.Info("Successfully sent admission review response")
}

// mutatePod takes a raw pod object, applies the mutation logic, and returns the mutated pod.
func mutatePod(raw []byte) ([]byte, error) {
	var pod corev1.Pod
	logrus.Info("Unmarshalling pod")
	if err := json.Unmarshal(raw, &pod); err != nil {
		logrus.WithError(err).Error("Failed to unmarshal pod")
		return nil, err
	}

	// Create a patch to add an environment variable to all containers
	patch := []map[string]interface{}{}

	// Loop through all containers and generate a patch to add the environment variable
	for i, c := range pod.Spec.Containers {
		logrus.Infof("Adding environment variable to container: %s", c.Name)
		patch = append(patch, map[string]interface{}{
			"op":   "add",
			"path": "/spec/containers/" + strconv.Itoa(i) + "/env",
			"value": []corev1.EnvVar{
				{
					Name:  "MUTATED_BY_WEBHOOK",
					Value: "true",
				},
			},
		})
	}

	// Marshal the patch to a JSON byte slice
	logrus.Info("Marshalling patch")
	return json.Marshal(patch)
}
