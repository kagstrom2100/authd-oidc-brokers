// Package authmodes lists the authentication modes that providers can support.
package authmodes

const (
	// Password is the ID of the password authentication method.
	Password = "password"

	// Device is the ID of the device authentication method.
	Device = "device_auth"

	// DeviceQr is the ID of the device authentication method when QrCode rendering is enabled.
	DeviceQr = "device_auth_qr"

	// NewPassword is the ID of the new password configuration method.
	NewPassword = "newpassword"
)
