package mobilesecurityservicebackup

import (
	"context"
	"fmt"
	"github.com/aerogear/mobile-security-service-operator/pkg/apis/mobilesecurityservice/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// updateStatusWithInvalidNamespace returns error when status regards the all required resources could not be updated
func (r *ReconcileMobileSecurityServiceBackup) updateStatusWithInvalidNamespace(reqLogger logr.Logger, request reconcile.Request) error {
	reqLogger.Info("Updating Backup Status with invalid namespace for the MobileSecurityServiceBackup")

	// Get the latest version of CR
	instance, err := r.fetchBkpInstance(reqLogger, request)
	if err != nil {
		return err
	}

	status := "Invalid Namespace"

	//Update Bind CR Status with OK
	if !reflect.DeepEqual(status, instance.Status.BackupStatus) {
		// Get the latest version of the CR in order to try to avoid errors when try to update the CR
		instance, err := r.fetchBkpInstance(reqLogger, request)
		if err != nil {
			return err
		}

		// Set the data
		instance.Status.BackupStatus = status

		// Update the CR
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update Status for the MobileSecurityServiceBackup")
			return err
		}
	}
	return nil
}

//updateAppStatus returns error when status regards the all required resources could not be updated
func (r *ReconcileMobileSecurityServiceBackup) updateBackupStatus(reqLogger logr.Logger, cronJobStatus *v1beta1.CronJob, dbSecretStatus, awsSecretStatus *corev1.Secret, dbPod *corev1.Pod, dbService *corev1.Service, request reconcile.Request) error {
	reqLogger.Info("Updating Backup Status for the MobileSecurityServiceBackup")

	//Get the latest version of the CR
	instance, err := r.fetchBkpInstance(reqLogger, request)
	if err != nil {
		return err
	}

	// Check just encryptionKeySecret which is Optional
	hasErrorWithEncKey, err := r.hasErrorWithEncKey(instance, reqLogger)
	if err != nil {
		return err
	}

	// Check if ALL required objects are created
	if len(cronJobStatus.Name) < 1 || len(dbSecretStatus.Name) < 1 || len(awsSecretStatus.Name) < 1 || dbPod == nil || len(dbPod.Name) < 1 || dbService == nil || len(dbService.Name) < 1 || hasErrorWithEncKey {
		err := fmt.Errorf("Failed to get OK Status for MobileSecurityServiceBackup")
		reqLogger.Error(err, "One of the resources are not created", "MobileSecurityServiceBackup.Namespace", instance.Namespace, "MobileSecurityServiceBackup.Name", instance.Name)
		return err
	}
	status := "OK"

	// Update Backup Status == OK
	if !reflect.DeepEqual(status, instance.Status.BackupStatus) {
		// Get the latest version of the CR in order to try to avoid errors when try to update the CR
		instance, err := r.fetchBkpInstance(reqLogger, request)
		if err != nil {
			return err
		}

		// Set the data
		instance.Status.BackupStatus = status

		// Update the CR
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update BackupStatus for the MobileSecurityServiceBackup")
			return err
		}
	}
	return nil
}

//hasErrorWithEncKey return true when the name or details or informed and was not possible check it
func (r *ReconcileMobileSecurityServiceBackup) hasErrorWithEncKey(instance *v1alpha1.MobileSecurityServiceBackup, reqLogger logr.Logger) (bool, error) {
	if hasEncryptionKeySecret(instance) {
		//Get the latest version of the CR
		encKey, err := r.fetchSecret(reqLogger, getEncSecretNamespace(instance), getEncSecretName(instance))
		if err != nil {
			return false, err
		}
		if len(encKey.Name) < 1 {
			return true, nil
		}
	}
	return false, nil
}

func (r *ReconcileMobileSecurityServiceBackup) updateCronJobStatus(reqLogger logr.Logger, request reconcile.Request) (*v1beta1.CronJob, error) {
	reqLogger.Info("Updating CronJob Status and Name for the MobileSecurityServiceBackup")
	// Get the latest version of the CR
	instance, err := r.fetchBkpInstance(reqLogger, request)
	if err != nil {
		return nil, err
	}

	// Get Object
	cronJobStatus, err := r.fetchCronJob(reqLogger, instance)
	if err != nil {
		reqLogger.Error(err, "Failed to get cronJob for Status", "MobileSecurityServiceBackup.Namespace", instance.Namespace, "MobileSecurityServiceBackup.Name", instance.Name)
		return cronJobStatus, err
	}

	//Update the CR
	if cronJobStatus.Name != instance.Status.CronJobName || !reflect.DeepEqual(cronJobStatus.Status, instance.Status.CronJobStatus) {
		// Get the latest version of the CR in order to try to avoid errors when try to update the CR
		instance, err := r.fetchBkpInstance(reqLogger, request)
		if err != nil {
			return nil, err
		}

		// Set the data
		instance.Status.CronJobStatus = cronJobStatus.Status
		instance.Status.CronJobName = cronJobStatus.Name

		// Update the CR
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update cronJob Name and Status for the MobileSecurityServiceBackup")
			return cronJobStatus, err
		}
	}

	return cronJobStatus, nil
}

func (r *ReconcileMobileSecurityServiceBackup) updateAWSSecretStatus(reqLogger logr.Logger, request reconcile.Request) (*corev1.Secret, error) {
	reqLogger.Info("Updating AwsSecret Name and Data Status for the MobileSecurityServiceBackup")
	// Get the latest version of the CR
	instance, err := r.fetchBkpInstance(reqLogger, request)
	if err != nil {
		return nil, err
	}

	// Get Object
	awsSecretStatus, err := r.fetchSecret(reqLogger, getAwsSecretNamespace(instance), getAWSSecretName(instance))
	if err != nil {
		reqLogger.Error(err, "Failed to get AwsSecret for Status", "MobileSecurityServiceBackup.Namespace", instance.Namespace, "MobileSecurityServiceBackup.Name", instance.Name)
		return awsSecretStatus, err
	}

	data := make(map[string]string)
	if awsSecretStatus.Data != nil {
		for k, v := range awsSecretStatus.Data {
			value := ""
			if v != nil {
				value = string(v)
			}
			data[k] = value
		}
	}

	//Update the CR
	if awsSecretStatus.Name != instance.Status.AWSSecretName || !reflect.DeepEqual(data, instance.Status.AWSSecretData) || awsSecretStatus.Namespace != instance.Status.AwsCredentialsSecretNamespace {
		// Get the latest version of the CR in order to try to avoid errors when try to update the CR
		instance, err := r.fetchBkpInstance(reqLogger, request)
		if err != nil {
			return nil, err
		}

		// Set the data
		instance.Status.AWSSecretName = awsSecretStatus.Name
		instance.Status.AWSSecretData = data
		instance.Status.AwsCredentialsSecretNamespace = awsSecretStatus.Namespace

		// Update the CR
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update AwsSecret Name and Data Status for the MobileSecurityServiceBackup")
			return awsSecretStatus, err
		}
	}

	return awsSecretStatus, nil
}

func (r *ReconcileMobileSecurityServiceBackup) updateEncSecretStatus(reqLogger logr.Logger, request reconcile.Request) error {
	reqLogger.Info("Updating EncryptionKey Name and Data Status for the MobileSecurityServiceBackup")
	// Get the latest version of the CR
	instance, err := r.fetchBkpInstance(reqLogger, request)
	if err != nil {
		return err
	}

	// check if the optinal secret was or not used
	hasEncKeySecret := hasEncryptionKeySecret(instance)

	// If has, update name and data status
	if hasEncKeySecret {
		encSecretStatus, err := r.fetchSecret(reqLogger, getEncSecretNamespace(instance), getEncSecretName(instance))
		if err != nil {
			reqLogger.Error(err, "Failed to get EncryptionKey for Status", "MobileSecurityServiceBackup.Namespace", instance.Namespace, "MobileSecurityServiceBackup.Name", instance.Name)
			return err
		}

		data := make(map[string]string)
		if encSecretStatus.Data != nil {
			for k, v := range encSecretStatus.Data {
				value := ""
				if v != nil {
					value = string(v)
				}
				data[k] = value
			}
			for k, v := range encSecretStatus.StringData {
				data[k] = v
			}
		}

		//Update the CR
		if encSecretStatus.Name != instance.Status.AWSSecretName || !reflect.DeepEqual(data, instance.Status.AWSSecretData) {
			// Get the latest version of the CR in order to try to avoid errors when try to update the CR
			instance, err := r.fetchBkpInstance(reqLogger, request)
			if err != nil {
				return err
			}

			// Set the data
			instance.Status.EncryptionKeySecretName = encSecretStatus.Name
			instance.Status.EncryptionKeySecretData = data
			instance.Status.EncryptionKeySecretNamespace = encSecretStatus.Namespace

			// Update the CR
			err = r.client.Status().Update(context.TODO(), instance)
			if err != nil {
				reqLogger.Error(err, "Failed to update EncryptionKey Name and Data Status for the MobileSecurityServiceBackup")
				return err
			}
		}
	}

	// Update boolean status
	if hasEncKeySecret != instance.Status.HasEncryptionKey {
		// Set the data
		instance.Status.HasEncryptionKey = hasEncKeySecret

		// Update the CR
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update EncryptionKey boolean Status for the MobileSecurityServiceBackup")
			return err
		}
	}

	return nil
}

func (r *ReconcileMobileSecurityServiceBackup) updateDBSecretStatus(reqLogger logr.Logger, request reconcile.Request) (*corev1.Secret, error) {
	reqLogger.Info("Updating DBSecret Name Status for the MobileSecurityServiceBackup")
	// Get the latest version of the CR
	instance, err := r.fetchBkpInstance(reqLogger, request)
	if err != nil {
		return nil, err
	}

	// Get Object
	dbSecretStatus, err := r.fetchSecret(reqLogger, instance.Namespace, dbSecretPrefix+instance.Name)
	if err != nil {
		reqLogger.Error(err, "Failed to get DBSecret for Status", "MobileSecurityServiceBackup.Namespace", instance.Namespace, "MobileSecurityServiceBackup.Name", instance.Name)
		return dbSecretStatus, err
	}

	data := make(map[string]string)
	if dbSecretStatus.Data != nil {
		for k, v := range dbSecretStatus.Data {
			value := ""
			if v != nil {
				value = string(v)
			}
			data[k] = value
		}
	}
	//Update the CR
	if dbSecretStatus.Name != instance.Status.DBSecretName || !reflect.DeepEqual(data, instance.Status.DBSecretData) {
		// Get the latest version of the CR in order to try to avoid errors when try to update the CR
		instance, err := r.fetchBkpInstance(reqLogger, request)
		if err != nil {
			return nil, err
		}

		// Set the data
		instance.Status.DBSecretName = dbSecretStatus.Name
		instance.Status.DBSecretData = data

		// Update the CR
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update DBSecret Name Status for the MobileSecurityServiceBackup")
			return dbSecretStatus, err
		}
	}

	return dbSecretStatus, nil
}

func (r *ReconcileMobileSecurityServiceBackup) updatePodDatabaseFoundStatus(reqLogger logr.Logger, request reconcile.Request, dbPod *corev1.Pod) error {
	reqLogger.Info("Updating PodDatabaseFoundStatus and Name for the MobileSecurityServiceBackup")
	// Get the latest version of the CR
	instance, err := r.fetchBkpInstance(reqLogger, request)
	if err != nil {
		return err
	}

	found := false
	if dbPod != nil && len(dbPod.Name) > 0 {
		found = true
	}

	//Update the CR
	if found != instance.Status.DatabasePodFound {
		// Get the latest version of the CR in order to try to avoid errors when try to update the CR
		instance, err := r.fetchBkpInstance(reqLogger, request)
		if err != nil {
			return err
		}

		// Set the data
		instance.Status.DatabasePodFound = found

		// Update the CR
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update DatabasePodFound Status for the MobileSecurityServiceBackup")
			return err
		}
	}

	return nil
}

func (r *ReconcileMobileSecurityServiceBackup) updateServiceDatabaseFoundStatus(reqLogger logr.Logger, request reconcile.Request, dbService *corev1.Service) error {
	reqLogger.Info("Updating ServiceDatabaseFoundStatus and Name for the MobileSecurityServiceBackup")
	// Get the latest version of the CR
	instance, err := r.fetchBkpInstance(reqLogger, request)
	if err != nil {
		return err
	}

	found := false
	if dbService != nil && len(dbService.Name) > 0 {
		found = true
	}

	//Update the CR
	if found != instance.Status.DatabaseServiceFound {
		// Get the latest version of the CR in order to try to avoid errors when try to update the CR
		instance, err := r.fetchBkpInstance(reqLogger, request)
		if err != nil {
			return err
		}

		// Set the data
		instance.Status.DatabaseServiceFound = found

		// Update the CR
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update DatabasePodFound Status for the MobileSecurityServiceBackup")
			return err
		}
	}

	return nil
}
