package ch9329

const (
	// Standard commands
	CmdGetChipInfo     = 0x01 // Get chip info
	CmdSendKeyboard    = 0x02 // Send keyboard general data
	CmdSendMediaKey    = 0x03 // Send media key data
	CmdSendMouseAbs    = 0x04 // Send absolute mouse data
	CmdSendMouseRel    = 0x05 // Send relative mouse data
	CmdSendHID         = 0x06 // Send custom HID data
	CmdGetParams       = 0x08 // Get parameter config
	CmdSetParams       = 0x09 // Set parameter config
	CmdGetUSBString    = 0x0A // Get USB string descriptor
	CmdSetUSBString    = 0x0B // Set USB string descriptor
	CmdResetToDefaults = 0x0C // Reset to factory defaults
	CmdReboot          = 0x0F // Reboot chip

	// Extended commands (CH9329F only)
	CmdSendHID2         = 0x10 // Send extended HID data
	CmdReadHID2         = 0x91 // Read extended HID data
	CmdGPIOControl      = 0x12 // GPIO control
	CmdGetSpecialParams = 0x13 // Get special parameters
	CmdSetSpecialParams = 0x14 // Set special parameters
	CmdEnterBootloader  = 0x15 // Enter bootloader (IAP mode)
	CmdSendTouchData    = 0x16 // Send touch panel data

	// HID read command
	CmdReadHID = 0x87 // Read HID data
)
