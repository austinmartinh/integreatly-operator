// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package threescale

import (
	"net/http"
	"sync"
)

var (
	lockThreeScaleInterfaceMockAddAuthenticationProvider       sync.RWMutex
	lockThreeScaleInterfaceMockAddUser                         sync.RWMutex
	lockThreeScaleInterfaceMockDeleteUser                      sync.RWMutex
	lockThreeScaleInterfaceMockGetAuthenticationProviderByName sync.RWMutex
	lockThreeScaleInterfaceMockGetAuthenticationProviders      sync.RWMutex
	lockThreeScaleInterfaceMockGetUser                         sync.RWMutex
	lockThreeScaleInterfaceMockGetUsers                        sync.RWMutex
	lockThreeScaleInterfaceMockSetNamespace                    sync.RWMutex
	lockThreeScaleInterfaceMockSetUserAsAdmin                  sync.RWMutex
	lockThreeScaleInterfaceMockSetUserAsMember                 sync.RWMutex
	lockThreeScaleInterfaceMockUpdateUser                      sync.RWMutex
)

// Ensure, that ThreeScaleInterfaceMock does implement ThreeScaleInterface.
// If this is not the case, regenerate this file with moq.
var _ ThreeScaleInterface = &ThreeScaleInterfaceMock{}

// ThreeScaleInterfaceMock is a mock implementation of ThreeScaleInterface.
//
//     func TestSomethingThatUsesThreeScaleInterface(t *testing.T) {
//
//         // make and configure a mocked ThreeScaleInterface
//         mockedThreeScaleInterface := &ThreeScaleInterfaceMock{
//             AddAuthenticationProviderFunc: func(data map[string]string, accessToken string) (*http.Response, error) {
// 	               panic("mock out the AddAuthenticationProvider method")
//             },
//             AddUserFunc: func(username string, email string, password string, accessToken string) (*http.Response, error) {
// 	               panic("mock out the AddUser method")
//             },
//             DeleteUserFunc: func(userId int, accessToken string) (*http.Response, error) {
// 	               panic("mock out the DeleteUser method")
//             },
//             GetAuthenticationProviderByNameFunc: func(name string, accessToken string) (*AuthProvider, error) {
// 	               panic("mock out the GetAuthenticationProviderByName method")
//             },
//             GetAuthenticationProvidersFunc: func(accessToken string) (*AuthProviders, error) {
// 	               panic("mock out the GetAuthenticationProviders method")
//             },
//             GetUserFunc: func(username string, accessToken string) (*User, error) {
// 	               panic("mock out the GetUser method")
//             },
//             GetUsersFunc: func(accessToken string) (*Users, error) {
// 	               panic("mock out the GetUsers method")
//             },
//             SetNamespaceFunc: func(ns string)  {
// 	               panic("mock out the SetNamespace method")
//             },
//             SetUserAsAdminFunc: func(userId int, accessToken string) (*http.Response, error) {
// 	               panic("mock out the SetUserAsAdmin method")
//             },
//             SetUserAsMemberFunc: func(userId int, accessToken string) (*http.Response, error) {
// 	               panic("mock out the SetUserAsMember method")
//             },
//             UpdateUserFunc: func(userId int, username string, email string, accessToken string) (*http.Response, error) {
// 	               panic("mock out the UpdateUser method")
//             },
//         }
//
//         // use mockedThreeScaleInterface in code that requires ThreeScaleInterface
//         // and then make assertions.
//
//     }
type ThreeScaleInterfaceMock struct {
	// AddAuthenticationProviderFunc mocks the AddAuthenticationProvider method.
	AddAuthenticationProviderFunc func(data map[string]string, accessToken string) (*http.Response, error)

	// AddUserFunc mocks the AddUser method.
	AddUserFunc func(username string, email string, password string, accessToken string) (*http.Response, error)

	// DeleteUserFunc mocks the DeleteUser method.
	DeleteUserFunc func(userId int, accessToken string) (*http.Response, error)

	// GetAuthenticationProviderByNameFunc mocks the GetAuthenticationProviderByName method.
	GetAuthenticationProviderByNameFunc func(name string, accessToken string) (*AuthProvider, error)

	// GetAuthenticationProvidersFunc mocks the GetAuthenticationProviders method.
	GetAuthenticationProvidersFunc func(accessToken string) (*AuthProviders, error)

	// GetUserFunc mocks the GetUser method.
	GetUserFunc func(username string, accessToken string) (*User, error)

	// GetUsersFunc mocks the GetUsers method.
	GetUsersFunc func(accessToken string) (*Users, error)

	// SetNamespaceFunc mocks the SetNamespace method.
	SetNamespaceFunc func(ns string)

	// SetUserAsAdminFunc mocks the SetUserAsAdmin method.
	SetUserAsAdminFunc func(userId int, accessToken string) (*http.Response, error)

	// SetUserAsMemberFunc mocks the SetUserAsMember method.
	SetUserAsMemberFunc func(userId int, accessToken string) (*http.Response, error)

	// UpdateUserFunc mocks the UpdateUser method.
	UpdateUserFunc func(userId int, username string, email string, accessToken string) (*http.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddAuthenticationProvider holds details about calls to the AddAuthenticationProvider method.
		AddAuthenticationProvider []struct {
			// Data is the data argument value.
			Data map[string]string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// AddUser holds details about calls to the AddUser method.
		AddUser []struct {
			// Username is the username argument value.
			Username string
			// Email is the email argument value.
			Email string
			// Password is the password argument value.
			Password string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// DeleteUser holds details about calls to the DeleteUser method.
		DeleteUser []struct {
			// UserId is the userId argument value.
			UserId int
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// GetAuthenticationProviderByName holds details about calls to the GetAuthenticationProviderByName method.
		GetAuthenticationProviderByName []struct {
			// Name is the name argument value.
			Name string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// GetAuthenticationProviders holds details about calls to the GetAuthenticationProviders method.
		GetAuthenticationProviders []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// GetUser holds details about calls to the GetUser method.
		GetUser []struct {
			// Username is the username argument value.
			Username string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// GetUsers holds details about calls to the GetUsers method.
		GetUsers []struct {
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// SetNamespace holds details about calls to the SetNamespace method.
		SetNamespace []struct {
			// Ns is the ns argument value.
			Ns string
		}
		// SetUserAsAdmin holds details about calls to the SetUserAsAdmin method.
		SetUserAsAdmin []struct {
			// UserId is the userId argument value.
			UserId int
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// SetUserAsMember holds details about calls to the SetUserAsMember method.
		SetUserAsMember []struct {
			// UserId is the userId argument value.
			UserId int
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
		// UpdateUser holds details about calls to the UpdateUser method.
		UpdateUser []struct {
			// UserId is the userId argument value.
			UserId int
			// Username is the username argument value.
			Username string
			// Email is the email argument value.
			Email string
			// AccessToken is the accessToken argument value.
			AccessToken string
		}
	}
}

// AddAuthenticationProvider calls AddAuthenticationProviderFunc.
func (mock *ThreeScaleInterfaceMock) AddAuthenticationProvider(data map[string]string, accessToken string) (*http.Response, error) {
	if mock.AddAuthenticationProviderFunc == nil {
		panic("ThreeScaleInterfaceMock.AddAuthenticationProviderFunc: method is nil but ThreeScaleInterface.AddAuthenticationProvider was just called")
	}
	callInfo := struct {
		Data        map[string]string
		AccessToken string
	}{
		Data:        data,
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockAddAuthenticationProvider.Lock()
	mock.calls.AddAuthenticationProvider = append(mock.calls.AddAuthenticationProvider, callInfo)
	lockThreeScaleInterfaceMockAddAuthenticationProvider.Unlock()
	return mock.AddAuthenticationProviderFunc(data, accessToken)
}

// AddAuthenticationProviderCalls gets all the calls that were made to AddAuthenticationProvider.
// Check the length with:
//     len(mockedThreeScaleInterface.AddAuthenticationProviderCalls())
func (mock *ThreeScaleInterfaceMock) AddAuthenticationProviderCalls() []struct {
	Data        map[string]string
	AccessToken string
} {
	var calls []struct {
		Data        map[string]string
		AccessToken string
	}
	lockThreeScaleInterfaceMockAddAuthenticationProvider.RLock()
	calls = mock.calls.AddAuthenticationProvider
	lockThreeScaleInterfaceMockAddAuthenticationProvider.RUnlock()
	return calls
}

// AddUser calls AddUserFunc.
func (mock *ThreeScaleInterfaceMock) AddUser(username string, email string, password string, accessToken string) (*http.Response, error) {
	if mock.AddUserFunc == nil {
		panic("ThreeScaleInterfaceMock.AddUserFunc: method is nil but ThreeScaleInterface.AddUser was just called")
	}
	callInfo := struct {
		Username    string
		Email       string
		Password    string
		AccessToken string
	}{
		Username:    username,
		Email:       email,
		Password:    password,
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockAddUser.Lock()
	mock.calls.AddUser = append(mock.calls.AddUser, callInfo)
	lockThreeScaleInterfaceMockAddUser.Unlock()
	return mock.AddUserFunc(username, email, password, accessToken)
}

// AddUserCalls gets all the calls that were made to AddUser.
// Check the length with:
//     len(mockedThreeScaleInterface.AddUserCalls())
func (mock *ThreeScaleInterfaceMock) AddUserCalls() []struct {
	Username    string
	Email       string
	Password    string
	AccessToken string
} {
	var calls []struct {
		Username    string
		Email       string
		Password    string
		AccessToken string
	}
	lockThreeScaleInterfaceMockAddUser.RLock()
	calls = mock.calls.AddUser
	lockThreeScaleInterfaceMockAddUser.RUnlock()
	return calls
}

// DeleteUser calls DeleteUserFunc.
func (mock *ThreeScaleInterfaceMock) DeleteUser(userId int, accessToken string) (*http.Response, error) {
	if mock.DeleteUserFunc == nil {
		panic("ThreeScaleInterfaceMock.DeleteUserFunc: method is nil but ThreeScaleInterface.DeleteUser was just called")
	}
	callInfo := struct {
		UserId      int
		AccessToken string
	}{
		UserId:      userId,
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockDeleteUser.Lock()
	mock.calls.DeleteUser = append(mock.calls.DeleteUser, callInfo)
	lockThreeScaleInterfaceMockDeleteUser.Unlock()
	return mock.DeleteUserFunc(userId, accessToken)
}

// DeleteUserCalls gets all the calls that were made to DeleteUser.
// Check the length with:
//     len(mockedThreeScaleInterface.DeleteUserCalls())
func (mock *ThreeScaleInterfaceMock) DeleteUserCalls() []struct {
	UserId      int
	AccessToken string
} {
	var calls []struct {
		UserId      int
		AccessToken string
	}
	lockThreeScaleInterfaceMockDeleteUser.RLock()
	calls = mock.calls.DeleteUser
	lockThreeScaleInterfaceMockDeleteUser.RUnlock()
	return calls
}

// GetAuthenticationProviderByName calls GetAuthenticationProviderByNameFunc.
func (mock *ThreeScaleInterfaceMock) GetAuthenticationProviderByName(name string, accessToken string) (*AuthProvider, error) {
	if mock.GetAuthenticationProviderByNameFunc == nil {
		panic("ThreeScaleInterfaceMock.GetAuthenticationProviderByNameFunc: method is nil but ThreeScaleInterface.GetAuthenticationProviderByName was just called")
	}
	callInfo := struct {
		Name        string
		AccessToken string
	}{
		Name:        name,
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockGetAuthenticationProviderByName.Lock()
	mock.calls.GetAuthenticationProviderByName = append(mock.calls.GetAuthenticationProviderByName, callInfo)
	lockThreeScaleInterfaceMockGetAuthenticationProviderByName.Unlock()
	return mock.GetAuthenticationProviderByNameFunc(name, accessToken)
}

// GetAuthenticationProviderByNameCalls gets all the calls that were made to GetAuthenticationProviderByName.
// Check the length with:
//     len(mockedThreeScaleInterface.GetAuthenticationProviderByNameCalls())
func (mock *ThreeScaleInterfaceMock) GetAuthenticationProviderByNameCalls() []struct {
	Name        string
	AccessToken string
} {
	var calls []struct {
		Name        string
		AccessToken string
	}
	lockThreeScaleInterfaceMockGetAuthenticationProviderByName.RLock()
	calls = mock.calls.GetAuthenticationProviderByName
	lockThreeScaleInterfaceMockGetAuthenticationProviderByName.RUnlock()
	return calls
}

// GetAuthenticationProviders calls GetAuthenticationProvidersFunc.
func (mock *ThreeScaleInterfaceMock) GetAuthenticationProviders(accessToken string) (*AuthProviders, error) {
	if mock.GetAuthenticationProvidersFunc == nil {
		panic("ThreeScaleInterfaceMock.GetAuthenticationProvidersFunc: method is nil but ThreeScaleInterface.GetAuthenticationProviders was just called")
	}
	callInfo := struct {
		AccessToken string
	}{
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockGetAuthenticationProviders.Lock()
	mock.calls.GetAuthenticationProviders = append(mock.calls.GetAuthenticationProviders, callInfo)
	lockThreeScaleInterfaceMockGetAuthenticationProviders.Unlock()
	return mock.GetAuthenticationProvidersFunc(accessToken)
}

// GetAuthenticationProvidersCalls gets all the calls that were made to GetAuthenticationProviders.
// Check the length with:
//     len(mockedThreeScaleInterface.GetAuthenticationProvidersCalls())
func (mock *ThreeScaleInterfaceMock) GetAuthenticationProvidersCalls() []struct {
	AccessToken string
} {
	var calls []struct {
		AccessToken string
	}
	lockThreeScaleInterfaceMockGetAuthenticationProviders.RLock()
	calls = mock.calls.GetAuthenticationProviders
	lockThreeScaleInterfaceMockGetAuthenticationProviders.RUnlock()
	return calls
}

// GetUser calls GetUserFunc.
func (mock *ThreeScaleInterfaceMock) GetUser(username string, accessToken string) (*User, error) {
	if mock.GetUserFunc == nil {
		panic("ThreeScaleInterfaceMock.GetUserFunc: method is nil but ThreeScaleInterface.GetUser was just called")
	}
	callInfo := struct {
		Username    string
		AccessToken string
	}{
		Username:    username,
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockGetUser.Lock()
	mock.calls.GetUser = append(mock.calls.GetUser, callInfo)
	lockThreeScaleInterfaceMockGetUser.Unlock()
	return mock.GetUserFunc(username, accessToken)
}

// GetUserCalls gets all the calls that were made to GetUser.
// Check the length with:
//     len(mockedThreeScaleInterface.GetUserCalls())
func (mock *ThreeScaleInterfaceMock) GetUserCalls() []struct {
	Username    string
	AccessToken string
} {
	var calls []struct {
		Username    string
		AccessToken string
	}
	lockThreeScaleInterfaceMockGetUser.RLock()
	calls = mock.calls.GetUser
	lockThreeScaleInterfaceMockGetUser.RUnlock()
	return calls
}

// GetUsers calls GetUsersFunc.
func (mock *ThreeScaleInterfaceMock) GetUsers(accessToken string) (*Users, error) {
	if mock.GetUsersFunc == nil {
		panic("ThreeScaleInterfaceMock.GetUsersFunc: method is nil but ThreeScaleInterface.GetUsers was just called")
	}
	callInfo := struct {
		AccessToken string
	}{
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockGetUsers.Lock()
	mock.calls.GetUsers = append(mock.calls.GetUsers, callInfo)
	lockThreeScaleInterfaceMockGetUsers.Unlock()
	return mock.GetUsersFunc(accessToken)
}

// GetUsersCalls gets all the calls that were made to GetUsers.
// Check the length with:
//     len(mockedThreeScaleInterface.GetUsersCalls())
func (mock *ThreeScaleInterfaceMock) GetUsersCalls() []struct {
	AccessToken string
} {
	var calls []struct {
		AccessToken string
	}
	lockThreeScaleInterfaceMockGetUsers.RLock()
	calls = mock.calls.GetUsers
	lockThreeScaleInterfaceMockGetUsers.RUnlock()
	return calls
}

// SetNamespace calls SetNamespaceFunc.
func (mock *ThreeScaleInterfaceMock) SetNamespace(ns string) {
	if mock.SetNamespaceFunc == nil {
		panic("ThreeScaleInterfaceMock.SetNamespaceFunc: method is nil but ThreeScaleInterface.SetNamespace was just called")
	}
	callInfo := struct {
		Ns string
	}{
		Ns: ns,
	}
	lockThreeScaleInterfaceMockSetNamespace.Lock()
	mock.calls.SetNamespace = append(mock.calls.SetNamespace, callInfo)
	lockThreeScaleInterfaceMockSetNamespace.Unlock()
	mock.SetNamespaceFunc(ns)
}

// SetNamespaceCalls gets all the calls that were made to SetNamespace.
// Check the length with:
//     len(mockedThreeScaleInterface.SetNamespaceCalls())
func (mock *ThreeScaleInterfaceMock) SetNamespaceCalls() []struct {
	Ns string
} {
	var calls []struct {
		Ns string
	}
	lockThreeScaleInterfaceMockSetNamespace.RLock()
	calls = mock.calls.SetNamespace
	lockThreeScaleInterfaceMockSetNamespace.RUnlock()
	return calls
}

// SetUserAsAdmin calls SetUserAsAdminFunc.
func (mock *ThreeScaleInterfaceMock) SetUserAsAdmin(userId int, accessToken string) (*http.Response, error) {
	if mock.SetUserAsAdminFunc == nil {
		panic("ThreeScaleInterfaceMock.SetUserAsAdminFunc: method is nil but ThreeScaleInterface.SetUserAsAdmin was just called")
	}
	callInfo := struct {
		UserId      int
		AccessToken string
	}{
		UserId:      userId,
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockSetUserAsAdmin.Lock()
	mock.calls.SetUserAsAdmin = append(mock.calls.SetUserAsAdmin, callInfo)
	lockThreeScaleInterfaceMockSetUserAsAdmin.Unlock()
	return mock.SetUserAsAdminFunc(userId, accessToken)
}

// SetUserAsAdminCalls gets all the calls that were made to SetUserAsAdmin.
// Check the length with:
//     len(mockedThreeScaleInterface.SetUserAsAdminCalls())
func (mock *ThreeScaleInterfaceMock) SetUserAsAdminCalls() []struct {
	UserId      int
	AccessToken string
} {
	var calls []struct {
		UserId      int
		AccessToken string
	}
	lockThreeScaleInterfaceMockSetUserAsAdmin.RLock()
	calls = mock.calls.SetUserAsAdmin
	lockThreeScaleInterfaceMockSetUserAsAdmin.RUnlock()
	return calls
}

// SetUserAsMember calls SetUserAsMemberFunc.
func (mock *ThreeScaleInterfaceMock) SetUserAsMember(userId int, accessToken string) (*http.Response, error) {
	if mock.SetUserAsMemberFunc == nil {
		panic("ThreeScaleInterfaceMock.SetUserAsMemberFunc: method is nil but ThreeScaleInterface.SetUserAsMember was just called")
	}
	callInfo := struct {
		UserId      int
		AccessToken string
	}{
		UserId:      userId,
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockSetUserAsMember.Lock()
	mock.calls.SetUserAsMember = append(mock.calls.SetUserAsMember, callInfo)
	lockThreeScaleInterfaceMockSetUserAsMember.Unlock()
	return mock.SetUserAsMemberFunc(userId, accessToken)
}

// SetUserAsMemberCalls gets all the calls that were made to SetUserAsMember.
// Check the length with:
//     len(mockedThreeScaleInterface.SetUserAsMemberCalls())
func (mock *ThreeScaleInterfaceMock) SetUserAsMemberCalls() []struct {
	UserId      int
	AccessToken string
} {
	var calls []struct {
		UserId      int
		AccessToken string
	}
	lockThreeScaleInterfaceMockSetUserAsMember.RLock()
	calls = mock.calls.SetUserAsMember
	lockThreeScaleInterfaceMockSetUserAsMember.RUnlock()
	return calls
}

// UpdateUser calls UpdateUserFunc.
func (mock *ThreeScaleInterfaceMock) UpdateUser(userId int, username string, email string, accessToken string) (*http.Response, error) {
	if mock.UpdateUserFunc == nil {
		panic("ThreeScaleInterfaceMock.UpdateUserFunc: method is nil but ThreeScaleInterface.UpdateUser was just called")
	}
	callInfo := struct {
		UserId      int
		Username    string
		Email       string
		AccessToken string
	}{
		UserId:      userId,
		Username:    username,
		Email:       email,
		AccessToken: accessToken,
	}
	lockThreeScaleInterfaceMockUpdateUser.Lock()
	mock.calls.UpdateUser = append(mock.calls.UpdateUser, callInfo)
	lockThreeScaleInterfaceMockUpdateUser.Unlock()
	return mock.UpdateUserFunc(userId, username, email, accessToken)
}

// UpdateUserCalls gets all the calls that were made to UpdateUser.
// Check the length with:
//     len(mockedThreeScaleInterface.UpdateUserCalls())
func (mock *ThreeScaleInterfaceMock) UpdateUserCalls() []struct {
	UserId      int
	Username    string
	Email       string
	AccessToken string
} {
	var calls []struct {
		UserId      int
		Username    string
		Email       string
		AccessToken string
	}
	lockThreeScaleInterfaceMockUpdateUser.RLock()
	calls = mock.calls.UpdateUser
	lockThreeScaleInterfaceMockUpdateUser.RUnlock()
	return calls
}